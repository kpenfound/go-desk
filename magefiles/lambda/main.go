package main

import (
	"context"

	"github.com/dagger/cloak/sdk/go/dagger"
)

//
func (r *lambda) upload(ctx context.Context, awsAccessKeyID string, awsSecretAccessKey dagger.SecretID, s3Bucket string, s3Key string) (bool, error) {

	panic("implement me")

}

//
func (r *lambda) compress(ctx context.Context, directory dagger.FSID) (*dagger.Filesystem, error) {

	panic("implement me")

}
