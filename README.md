
TLDR
export $(cat .env | xargs) && go run main.go


This project is a simple script to upload a set of files to a digital ocean.


Code Pseudocode

doctl put -b bucketName [directory | file | list of files | 
