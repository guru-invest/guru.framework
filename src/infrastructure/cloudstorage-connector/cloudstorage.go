package cloudstorage_connector

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type CloudStorageConnector struct {
}

type FileData struct {
	Obj []byte
}

func (h CloudStorageConnector) GetObject(cloudCredential interface{}, bucketName, objectName string) FileData {
	ctx := context.Background()
	credentialsData, err := json.Marshal(cloudCredential)
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

	return FileData{
		Obj: result,
	}
}
