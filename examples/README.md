# Mapping Examples
This directory contains different examples of HTTP request/response.

Once you start the `mimic` with these examples(simply downmload and point `mimic` to directory), content of each JSON file will turn be used to create appropiate HTTP endpoint.

### GitHub User API Examples
Some examples from GitHub user API which can be tested via following commands

````
# Will return 200 with a user response in get-authenticated-user.json file
curl -i http://localhost:8080/user -H "Authorization: token someValidToken"


# Will return 401 with a response in get-authenticated-user-fails.json file
curl -i http://localhost:8080/user -H "Authorization: token someInvalidToken"


# Will return 200 with a response in get-user.json file
curl -i http://localhost:8080/user/octocat 

# username is dynamic, which  means following example will return the same response
curl -i http://localhost:8080/user/someotheruser 

````

For more details check out the JSON files.


### Fault Simulation
Purpose of fault simulation is to create artificial failures in responses. At this stage only supported failure is response-delay.

````
{
    "request":{
        "url": "/health",
        "method" : "GET"

    },
    "response":{
        "status": 200, 
        "body": "All healthy",
        "headers":{
            "Content-Type":"text/plain"
        },
        "withDelay": 4000
    }
}

````
The mapping above, will delay the response for 4 seconds, try it out with:

````
curl -i http://localhost:8080/health
````