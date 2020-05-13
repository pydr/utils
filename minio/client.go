package minio

import "github.com/minio/minio-go/v6"

type Client struct {
	client *minio.Client
}

func NewClient(endpoint, username, password string, useSSL bool) (*Client, error) {
	cli, err := minio.New(endpoint, username, password, useSSL)
	if err != nil {
		return nil, err
	}

	return &Client{client: cli}, nil
}

func (c *Client) Upload(bucket, location, filename, filepath string) error {
	exists, err := c.client.BucketExists(bucket)
	if err != nil {
		return err
	}

	if !exists {
		if err = c.client.MakeBucket(bucket, location); err != nil {
			return err
		}
	}

	_, err = c.client.FPutObject(bucket, filename, filepath, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	return err
}
