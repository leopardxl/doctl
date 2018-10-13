package doctl

import (
	"log"

	"github.com/minio/minio-go"
)

func ListContents(spaceName string) string {

	return " "
}

func Upload(spaceName string, file string) {

}

func Get(client *minio.Client, bucketName string, files []string) {

}

func Put(client *minio.Client, bucketName string, files []string) {
	for i := range files {
		filePath := files[i]
		contentType, _ := MimeType(filePath)
		objectName := ObjectName(filePath)
		n, err := client.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
	}

}
