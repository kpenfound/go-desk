//go:build mage

package main

import (

)

// Build and deploy the api to AWS Lambda
func Api(ctx context.Context) {

}

// Deploy the AWS Infrastructure with Terraform
func Infra(ctx context.Context) {

}

// Build and deploy the website to Netlify
func Website(ctx context.Context) {
	netlifyToken, ok := os.LookupEnv("NETLIFY_AUTH_TOKEN")
		if !ok {
			return fmt.Errorf("NETLIFY_AUTH_TOKEN not set")
		}

		// Load API token into a secret
		var token dagger.SecretID
		if resp, err := core.AddSecret(ctx, netlifyToken); err != nil {
			return err
		} else {
			token = resp.Core.AddSecret
		}

		// Load source code from workdir
		var source dagger.FSID
		if resp, err := core.Workdir(ctx); err != nil {
			return err
		} else {
			source = resp.Host.Workdir.Read.ID
		}

		// Deploy using the todoapp deploy extension
		resp, err := todoapp.Deploy(ctx, source, siteName, token)
		if err != nil {
			return err
		}

		// Print deployment info to the user
		fmt.Println("URL:", resp.Todoapp.Deploy)

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

