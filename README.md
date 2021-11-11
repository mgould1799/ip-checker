# IP Checker

## Description

This is a RESTFUL API that provides two routes. The first route is health checker to see if it the API is working. The second is route that takes in a JSON body that is provided an IP as a string and a list of Country names that are a white list

## Getting Started

### Dependencies

* golang needs to be installed 
* all package dependencies can be found within the gomod file
* docker


### Executing program

```
docker build ./ -t ip-checker 
docker-compose up 
curl localhost:8080/healthz 
```

#### Example API Calls and their Responses

Request to Check Ip is in a white list 
```
curl --location --request GET '0.0.0.0:8080/check-ip' \
--header 'Content-Type: application/json' \
--data-raw '
{
    "ip": "81.2.69.142",
    "whiteListedCountries": ["United States"]
}'
```

Response to see if an api is in a white list 

```
{
    "inList": false
}
```

Checking the health 

```
curl --location --request GET '0.0.0.0:8080/healthz' \
--header 'Content-Type: application/json' \
--data-raw '
'
```

Checking the health response 

```
{
    "working": true
}
```


## Fun bits

* If you desire to deploy into a kubernetes envrioment in the default namespace, you can use the kubespecs in the kubespec folder. A deployment and service files have been provided for you.



## License
 
The MIT License (MIT)

Copyright (c) 2021 Meagan Gould

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.