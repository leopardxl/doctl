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
	debug := flag.Bool("d", false, "debug")

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

	if *debug {
		//List all spaces. Digtal Ocean spaces are S3 buckets
		spaces, err := client.ListBuckets()
		if err != nil {
			log.Fatal(err)
		}

		for _, space := range spaces {
			if space.Name == spaceName {
				fmt.Printf("Space %s exists\n", spaceName)
			}
			fmt.Println(space.Name)
		}
	}

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
