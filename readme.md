# hashit
expose and http endpoint to hash (sha512) and base64 encode strings

## running locally
* build with ```go build```
* run with ```./hashit <port>``` (defaults to 8080)

## usage
1. the /hash POST endpoint accepts a 'password' form field
    * example: ```curl -X POST http://localhost:{port}/hash --data password=angryMonkey```
1. the /shutdown endpoint will trigger a graceful shutdown, stopping the server listening immediately but resolving any remaining requests


    
