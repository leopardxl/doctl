This project is a simple script to upload a set of files to a digital ocean.

go build
export $(cat .env | xargs)
doctl put -a put -b bucket SOURCE [SOURCE...]

Status: Abandoned

The desired functionality at this time is implemented by the MinIO client.
https://docs.minio.io/docs/minio-client-complete-guide.html
