// CMD Line utility for dealing with common Digital Ocean Workflows
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	doctl "github.com/leopardxl/doctl/dotools"
	"github.com/minio/minio-go"
)

func main() {

	//Setup configuration parameters
	accessKey := os.Getenv("DO_SPACES_KEY")
	secKey := os.Getenv("DO_SPACES_SECRET")
	endpoint := os.Getenv("DO_SPACES_ENDPOINT") // "nyc3.digitaloceanspaces.com"
	spaceName := os.Getenv("DO_SPACES_NAME")    // Space names must be globally unique
	ssl := true

	// Read user input
	action := os.Args[1]
	bucketName := *flag.String("b", "", "Name of the destination bucket")
	files := os.Args[4:]

	switch action {
	case "put":
		doctl.Put(bucketName, files)
	case "get":
		doctl.Get(bucketName, files)
	}

	fmt.Println(endpoint, accessKey, secKey, ssl)
	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKey, secKey, ssl)
	if err != nil {
		log.Fatal(err)
	}

	//List all spaces.
	spaces, err := client.ListBuckets()
	if err != nil {
		log.Fatal(err)
	}
	// objectName := "fp.py"
	// filePath := "/home/kemmanuel/projects/tmp/fp.py"
	// contentType := doctl.MimeType(filePath)
	for _, space := range spaces {

		if space.Name == spaceName {
			fmt.Printf("Space %s exists\n", spaceName)
			//fmt.Printf("checking if files already exist")
			//fmt.Printf("uploading unique files")

		}
		fmt.Println(space.Name)
	}
}
