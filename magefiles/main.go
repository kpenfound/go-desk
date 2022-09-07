//go:build mage

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dagger/cloak/engine"
	"github.com/dagger/cloak/sdk/go/dagger"

	vault "github.com/kpenfound/go-desk/magefiles/gen/godesk"
	"github.com/kpenfound/go-desk/magefiles/gen/netlify"
	"github.com/kpenfound/go-desk/magefiles/gen/yarn"
)

// Build and deploy the api to AWS Lambda
func Api(ctx context.Context) {

}

// Deploy the AWS Infrastructure with Terraform
func Infra(ctx context.Context) {

}

// Where our secrets are stored in Vault
const GODESK_VAULT_PATH = "godesk/deploy"
const NETLIFY_TOKEN = "netlify_token"

// Build and deploy the website to Netlify
func Website(ctx context.Context) {
	websiteFSID := "website"
	websitePath, err := filepath.Abs("./website")
	if err != nil {
		panic("Failed to deterimine workdir")
	}
	buildDirs := make(map[string]string)
	buildDirs[websiteFSID] = websitePath
	cfg := &engine.Config{
		LocalDirs: buildDirs,
	}

	if err := engine.Start(ctx, cfg, func(ctx engine.Context) error {
		// Configure Vault
		address, ok := os.LookupEnv("VAULT_ADDR")
		if !ok {
			return fmt.Errorf("VAULT_ADDR not set")
		}
		vaultToken, ok := os.LookupEnv("VAULT_TOKEN")
		if !ok {
			return fmt.Errorf("VAULT_TOKEN not set")
		}
		path := GODESK_VAULT_PATH
		key := NETLIFY_TOKEN

		// Get secret from Vault
		token, err := vault.Get(ctx, address, vaultToken, path, key)
		if err != nil {
			return fmt.Errorf("Failed to get vault secret")
		}
		siteName, ok := os.LookupEnv("NETLIFY_SITE_NAME")
		if !ok {
			user, _ := os.LookupEnv("USER")
			siteName = fmt.Sprintf("%s-desk", user)
		}
		fmt.Printf("Using Netlify site name %q\n", siteName)

		// Load website subdirectory
		var website dagger.FSID
		ldResp, err := vault.GetLocalDir(ctx, websiteFSID)
		if err != nil {
			return err
		}
		website = ldResp.Host.Dir.Read.ID

		// Yarn build
		yarnArgs := []string{"build"}
		_, err = yarn.Script(ctx, website, yarnArgs)
		if err != nil {
			return err
		}

		// Deploy to Netlify
		resp, err := netlify.Deploy(ctx, website, "build", siteName, token.Qvault.Secret)
		if err != nil {
			return err
		}

		// Print deployment info to the user
		fmt.Println("URL:", resp.Netlify.Deploy)

		return nil
	}); err != nil {
		panic(err)
	}
}

// Build the listener cli
func Listener(ctx context.Context) {

}

// Run the listener cli
func Listen(ctx context.Context) {

}
