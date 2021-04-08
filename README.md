# Mimic
Mimic is standalone HTTP-based mock server.

Mimic can be used for rapid development of an existing or non existing API, supporting different sort of failure modes to give you chance to develop against it.

## Installation 
You can install `mimic` via `brew` 
```
brew tap  enginyoyen/homebrew-tap 
brew install enginyoyen/homebrew-tap/mimic
```

or download the executable binary from the release page


## Getting started with `mimic`

Create a new directory and add a simple HTTP request/response mapping as a JSON file.
````
mkdir mappings
cd mappings

cat << EOF > ping-endpoint.json
{
    "request":{
        "url": "/ping",
        "method" : "GET"
    },
    "response":{
        "status": 200, 
        "body": "pong"
    }
}
EOF

````


Start the mimic with the directory that you created above, which should contain the request/response mapping. 

````
mimic -m [PATH]/mappings
````

Test it via `curl`

````
curl -i http://localhost:8080/ping
````

You should get the response

````
HTTP/1.1 200 OK
Date: Sat, 03 Apr 2021 18:39:13 GMT
Content-Length: 4
Content-Type: text/plain; charset=utf-8

pong
````


## Mappings
Each JSON file represents a single HTTP request/response mapping. 
You can point `mimic` to load mappings from a directory. A single directory can contain any number of mappings or sub-directories that contains mappings.  

A following JSON object has a two fields represents a HTTP request that is processed by the mock server and the response that is generated based on that request.

````
{
    "request":{
        "url": "/tell-me",
        "method" : "POST",
        "body": "Some message"
    },
    "response":{
        "status": 200, 
        "body": "Here it is!",
        "headers":{
            "Content-Type":"text/plain"
        }
    }
}

````

### Response Body as JSON
Response body can also be JSON object as in the following example:
{
    "request":{
        "url": "/user",
        "method" : "GET"

    },
    "response":{
        "status": 200, 
        "body": {
            "login": "octocat",
            "id": 1,
            "node_id": "MDQ6VXNlcjE=",
            "avatar_url": "https://github.com/images/error/octocat_happy.gif",
            "url": "https://api.github.com/users/octocat",
            "location": "San Francisco",
            "email": "octocat@github.com",
            "plan": {
              "name": "Medium",
              "space": 400,
              "private_repos": 20,
              "collaborators": 0
            }
        },
        "headers":{
            "Content-Type":"application/json"
        }
    }
}

### Fault Simulation


### Examples
`examples` directory contains diverse request/response mapping to get you started easily.


# Licence
Use of this source code is governed by an MIT license that can be found in the `LICENSE` file.