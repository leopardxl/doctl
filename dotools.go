// CMD Line utility for dealing with common Digital Ocean Workflows

package dotools

import (
	"fmt"
	"log"
	"os"

	minio "github.com/minio/minio-go"
)

func main() {
	accessKey := os.Getenv("DO_SPACES_KEY")
	secKey := os.Getenv("DO_SPACES_SECRET")
	endpoint := "nyc3.digitaloceanspaces.com"
	// spaceName := "avelli" // Space names must be globally unique
	ssl := true

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

	for _, space := range spaces {
		fmt.Println(space.Name)
	}
}
