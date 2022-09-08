// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package godesk

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/dagger/cloak/sdk/go/dagger"
)

// GetLocalDirHost includes the requested fields of the GraphQL type Host.
// The GraphQL type's documentation follows.
//
// Interactions with the user's host filesystem
type GetLocalDirHost struct {
	// Fetch a client directory
	Dir GetLocalDirHostDirLocalDir `json:"dir"`
}

// GetDir returns GetLocalDirHost.Dir, and is useful for accessing the field via an interface.
func (v *GetLocalDirHost) GetDir() GetLocalDirHostDirLocalDir { return v.Dir }

// GetLocalDirHostDirLocalDir includes the requested fields of the GraphQL type LocalDir.
// The GraphQL type's documentation follows.
//
// A directory on the user's host filesystem
type GetLocalDirHostDirLocalDir struct {
	// Read the contents of the directory
	Read dagger.Filesystem `json:"read"`
}

// GetRead returns GetLocalDirHostDirLocalDir.Read, and is useful for accessing the field via an interface.
func (v *GetLocalDirHostDirLocalDir) GetRead() dagger.Filesystem { return v.Read }

// GetLocalDirResponse is returned by GetLocalDir on success.
type GetLocalDirResponse struct {
	// Host API
	Host GetLocalDirHost `json:"host"`
}

// GetHost returns GetLocalDirResponse.Host, and is useful for accessing the field via an interface.
func (v *GetLocalDirResponse) GetHost() GetLocalDirHost { return v.Host }

// GetQvault includes the requested fields of the GraphQL type Qvault.
type GetQvault struct {
	Secret dagger.SecretID `json:"secret"`
}

// GetSecret returns GetQvault.Secret, and is useful for accessing the field via an interface.
func (v *GetQvault) GetSecret() dagger.SecretID { return v.Secret }

// GetResponse is returned by Get on success.
type GetResponse struct {
	Qvault GetQvault `json:"qvault"`
}

// GetQvault returns GetResponse.Qvault, and is useful for accessing the field via an interface.
func (v *GetResponse) GetQvault() GetQvault { return v.Qvault }

// __GetInput is used internally by genqlient
type __GetInput struct {
	Address string `json:"address"`
	Token   string `json:"token"`
	Path    string `json:"path"`
	Key     string `json:"key"`
}

// GetAddress returns __GetInput.Address, and is useful for accessing the field via an interface.
func (v *__GetInput) GetAddress() string { return v.Address }

// GetToken returns __GetInput.Token, and is useful for accessing the field via an interface.
func (v *__GetInput) GetToken() string { return v.Token }

// GetPath returns __GetInput.Path, and is useful for accessing the field via an interface.
func (v *__GetInput) GetPath() string { return v.Path }

// GetKey returns __GetInput.Key, and is useful for accessing the field via an interface.
func (v *__GetInput) GetKey() string { return v.Key }

// __GetLocalDirInput is used internally by genqlient
type __GetLocalDirInput struct {
	Path string `json:"path"`
}

// GetPath returns __GetLocalDirInput.Path, and is useful for accessing the field via an interface.
func (v *__GetLocalDirInput) GetPath() string { return v.Path }

func Get(
	ctx context.Context,
	address string,
	token string,
	path string,
	key string,
) (*GetResponse, error) {
	req := &graphql.Request{
		OpName: "Get",
		Query: `
query Get ($address: String!, $token: String!, $path: String!, $key: String!) {
	qvault {
		secret(address: $address, token: $token, path: $path, key: $key)
	}
}
`,
		Variables: &__GetInput{
			Address: address,
			Token:   token,
			Path:    path,
			Key:     key,
		},
	}
	var err error
	var client graphql.Client

	client, err = dagger.Client(ctx)
	if err != nil {
		return nil, err
	}

	var data GetResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func GetLocalDir(
	ctx context.Context,
	path string,
) (*GetLocalDirResponse, error) {
	req := &graphql.Request{
		OpName: "GetLocalDir",
		Query: `
query GetLocalDir ($path: String!) {
	host {
		dir(id: $path) {
			read {
				id
			}
		}
	}
}
`,
		Variables: &__GetLocalDirInput{
			Path: path,
		},
	}
	var err error
	var client graphql.Client

	client, err = dagger.Client(ctx)
	if err != nil {
		return nil, err
	}

	var data GetLocalDirResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}