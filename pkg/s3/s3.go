package s3

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Options struct {
	Endpoint       string `mapstructure:"endpoint"`
	ForcePathStyle bool   `mapstructure:"force-path-style"`
	AccessKey      string `mapstructure:"access-key"`
	SecretKey      string `mapstructure:"secret-key"`
	SessionToken   string `mapstructure:"session-token"`
	DisableSSL     bool   `mapstructure:"disable-ssl"`
	Region         string `mapstructure:"region"`
}

func Create(ctx context.Context, options Options, bucket string) (*s3.S3, error) {
	var endpoint *string
	if options.Endpoint != "" {
		endpoint = &options.Endpoint
	}

	cfg := &aws.Config{
		Credentials: credentials.NewStaticCredentials(
			options.AccessKey, options.SecretKey, options.SessionToken,
		),
		Endpoint:         endpoint,
		Region:           &options.Region,
		DisableSSL:       &options.DisableSSL,
		S3ForcePathStyle: &options.ForcePathStyle,
	}

	sess, err := session.NewSession(cfg)
	if err != nil {
		return nil, err
	}

	s3Client := s3.New(sess, cfg)

	if bucket != "" {
		bucketExist := false

		buckets, err := s3Client.ListBucketsWithContext(ctx, &s3.ListBucketsInput{})
		if err != nil {
			return nil, err
		}

		for _, b := range buckets.Buckets {
			if *b.Name == bucket {
				bucketExist = true
				break
			}
		}

		if !bucketExist {
			input := &s3.CreateBucketInput{
				Bucket: &bucket,
				CreateBucketConfiguration: &s3.CreateBucketConfiguration{
					LocationConstraint: &options.Region,
				},
			}

			if _, err = s3Client.CreateBucketWithContext(ctx, input); err != nil {
				return nil, err
			}
		}
	}

	return s3Client, nil
}
