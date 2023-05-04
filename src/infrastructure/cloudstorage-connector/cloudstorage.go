package cloudstorage_connector

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"cloud.google.com/go/storage"
	"github.com/guru-invest/guru.framework/src/crossCutting/options"
	"github.com/guru-invest/guru.framework/src/models"
	"google.golang.org/api/option"
)

type Credential interface {
	options.GCPCredentialsOption | options.AWSCredentialsOption
}

type CloudStorageConnector[T Credential] struct {
	data T
}

func (c *CloudStorageConnector[T]) GetObject(cloudCredential T, bucketName, objectName string) models.StorageFile {
	ctx := context.Background()
	c.data = cloudCredential
	credentialsData, err := json.Marshal(c.data)
	if err != nil {
		log.Fatalf("Failed to parse credentials: %v", err)
	}
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON([]byte(credentialsData)))
	if err != nil {
		log.Fatalf("Failed to create storage client: %v", err)
	}

	bucket := client.Bucket(bucketName)

	obj, err := bucket.Object(objectName).NewReader(ctx)
	if err != nil {
		log.Fatalf("Failed to open object: %v", err)
	}
	defer obj.Close()

	result, err := ioutil.ReadAll(obj)
	if err != nil {
		log.Fatalf("Failed to read object: %v", err)
	}

	return models.StorageFile{
		Obj: result,
	}
}
