module github.com/kpenfound/go-desk

go 1.16

require (
	github.com/Khan/genqlient v0.5.0
	github.com/aws/aws-sdk-go v1.38.15
	github.com/dagger/cloak v0.0.0-20220906161451-1c7c9f8f035e
	github.com/hashicorp/vault/api v1.7.2
	github.com/mitchellh/cli v1.1.2
)

replace github.com/docker/docker => github.com/docker/docker v20.10.3-0.20220414164044-61404de7df1a+incompatible
