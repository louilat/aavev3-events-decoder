package datalab

import (
	"encoding/json"
	"io"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/minio/minio-go"
)

func ReadRawEvents(endpoint, bucket, key, accessKeyID, secretAccessKey string) ([]types.Log, error) {
	useSSL := false

	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		return make([]types.Log, 0), err
	}

	usr, err := minioClient.GetObject(bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return make([]types.Log, 0), err
	}

	usersBytes, err := io.ReadAll(usr)
	if err != nil {
		return make([]types.Log, 0), err
	}

	users := make([]types.Log, 0)
	err = json.Unmarshal(usersBytes, &users)
	if err != nil {
		return make([]types.Log, 0), err
	}
	return users, nil
}
