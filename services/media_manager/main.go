package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	client, err := connect()
	if err != nil {
		log.Fatalf("Failed to connect to MinIO: %s\n", err)
	}

	s := &storage{c: client, bucketName: os.Getenv("MINIO_BUCKET_NAME")}

	http.HandleFunc("POST /", uploadImage(s))
	http.HandleFunc("GET /{name}", receiveImage(s))

	log.Fatal(http.ListenAndServe(":80", nil))
}

func connect() (*minio.Client, error) {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyId := os.Getenv("MINIO_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACCESS_KEY")
	useSSL := strings.ToLower(os.Getenv("MINIO_USE_SSL")) == "true"

	return minio.New(
		endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
			Secure: useSSL,
		},
	)
}
