# hashit
hash (sha512) and base64 encode strings

## running locally
* run with ```go run main.go -port=<port>``` (defaults to 8080)

## usage
1. the /hash POST endpoint accepts a 'password' form field and returns the id of the hash
    * example: ```curl -X POST http://localhost:{port}/hash --data password=angryMonkey```
1. the /hash/{id} GET endpoint returns the specified hash
1. the /shutdown endpoint will trigger a graceful shutdown, stopping the server listening immediately but resolving any remaining requests
1. the /stats endpoint returns the server stats. Total is the total number of hash requests and Average is the average time it takes to process a request in milliseconds
    * ```{"total":3,"average":0.066073}```


    
