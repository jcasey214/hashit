# hashit
expose and http endpoint to hash (sha512) and base64 encode strings

## running locally
* build with ```go build```
* run with ```./hashit <port>``` (defaults to 8080)

## usage
the /hash POST endpoint accepts a 'password' form field

example curl: \
curl -X POST http://localhost:{port}/hash --data password=angryMonkey
    
