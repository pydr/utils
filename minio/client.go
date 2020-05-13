package minio

import (
	"github.com/minio/minio-go/v6"
	"io"
)

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

func (c *Client) UploadSteam(bucket, location, filename string, file io.Reader, fileSize int64) error {
	exists, err := c.client.BucketExists(bucket)
	if err != nil {
		return err
	}

	if !exists {
		if err = c.client.MakeBucket(bucket, location); err != nil {
			return err
		}
	}

	_, err = c.client.PutObject(bucket, filename, file, fileSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	return err
}

func (c *Client) Download(bucket, filename, saveDir string) error {
	return c.client.FGetObject(bucket, filename, saveDir, minio.GetObjectOptions{})
}
