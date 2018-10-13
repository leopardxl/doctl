// CMD Line utility for copying files to DigitalOcean
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
	action := flag.String("a", "", `Type of action, "get" or "put"`)
	bucketName := flag.String("b", "", "Name of the destination bucket") //Pointer to bucketName

	flag.Parse()
	files := flag.Args()
	fmt.Printf("flags: -b %s\nfiles:%v\n", *bucketName, files)

	validPaths := validatePaths(flag.Args())

	fmt.Printf("Minio client settings:%v, %v, %v, %v\n",
		endpoint, accessKey, secKey, ssl)
	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKey, secKey, ssl)
	if err != nil {
		log.Fatal(err)
	}

	//List all spaces. Digtal Ocean spaces are S3 buckets
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
	// List all buckets
	buckets, err := client.ListBuckets()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listing Buckets")
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}

	prefixTest(client, spaceName, "dl")

	//Do the desired action
	switch *action {
	case "put":
		doctl.Put(client, *bucketName, validPaths)
	case "get":
		doctl.Get(client, *bucketName, validPaths)
	}

}

func validatePaths(paths []string) []string {
	var validPaths []string
	for _, path := range paths {
		isDir, err := doctl.IsDirectory(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading %s, skipping", path)
			continue
		}
		if isDir {
			fmt.Fprintf(os.Stderr, "Directories are currently unsupported, skipping %s", path)
			continue
		}
		validPaths = append(validPaths, path)

	}
	return validPaths
}

func prefixTest(client *minio.Client, bucket, prefix string) {
	// Create a done channel to control 'ListObjectsV2' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true
	objectCh := client.ListObjectsV2(bucket, prefix, isRecursive, doneCh)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println(object)
	}
}
