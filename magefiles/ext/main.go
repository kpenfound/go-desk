package main

import (
	"context"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/dagger/cloak/sdk/go/dagger"
	vault "github.com/hashicorp/vault/api"
)

// Based on https://github.com/hashicorp/vault-examples/blob/main/examples/_quick-start/go/example.go
func (r *qvault) secret(ctx context.Context, address, token, path, key string) (dagger.SecretID, error) {
	var sid dagger.SecretID

	// Configure Vault client
	config := vault.DefaultConfig()

	config.Address = address

	client, err := vault.NewClient(config)
	if err != nil {
		return sid, fmt.Errorf("unable to initialize Vault client: %v", err)
	}

	client.SetToken(token)

	// Read from secrets kv store
	secret, err := client.KVv2("secret").Get(ctx, path)
	if err != nil {
		return sid, fmt.Errorf("unable to read secret: %v", err)
	}

	value, ok := secret.Data[key].(string)
	if !ok {
		return sid, fmt.Errorf("value type assertion failed: %T %#v", secret.Data[key], secret.Data[key])
	}

	// Add secret to dagger secrets store
	resp, err := AddSecret(ctx, value)
	if err != nil {
		return sid, fmt.Errorf("unable to store dagger secret: %v", err)
	}

	// Return secret id
	sid = resp.Core.AddSecret

	return sid, nil
}

type AddSecretCore struct {
	// Add a secret
	AddSecret dagger.SecretID `json:"addSecret"`
}

// AddSecretResponse is returned by AddSecret on success.
type AddSecretResponse struct {
	// Core API
	Core AddSecretCore `json:"core"`
}

// __AddSecretInput is used internally by genqlient
type __AddSecretInput struct {
	Plaintext string `json:"plaintext"`
}

func AddSecret(
	ctx context.Context,
	plaintext string,
) (*AddSecretResponse, error) {
	req := &graphql.Request{
		OpName: "AddSecret",
		Query: `
query AddSecret ($plaintext: String!) {
	core {
		addSecret(plaintext: $plaintext)
	}
}
`,
		Variables: &__AddSecretInput{
			Plaintext: plaintext,
		},
	}
	var err error
	var client graphql.Client

	client, err = dagger.Client(ctx)
	if err != nil {
		return nil, err
	}

	var data AddSecretResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
