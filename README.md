This project is a simple script to upload a set of files to a digital ocean.

go build
export $(cat .env | xargs)
doctl put -a put -b bucket SOURCE [SOURCE...]

 
