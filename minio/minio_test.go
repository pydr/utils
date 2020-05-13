package minio

import "testing"

func TestMinio(t *testing.T) {
	client, err := NewClient("192.168.1.160:9000", "admin", "12345678", false)
	if err != nil {
		t.Fatalf("init minio client failed: %v", err)
	}

	err = client.Upload("test", "local", "1.txt", "./test.txt")
	if err != nil {
		t.Fatalf("upload file to minio failed: %v", err)
	}

	err = client.Download("test", "1.txt", "./2.txt")
	if err != nil {
		t.Fatalf("download file from minio failed: %v", err)
	}
}
