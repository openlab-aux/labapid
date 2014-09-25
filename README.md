labapid
=======

Generates and serves SpaceAPI (http://spaceapi.net/documentation) data and reacts to events

Installation
------------
This assumes a working golang environment.

```
# get depencies
go get github.com/waaaaargh/gospaceapi

# get labapid
git clone https://github.com/openlab-aux/labapid.git

# fill in your space's information
cp spaceapi.json.data.example spaceapi.json.data
vim spaceapi.json.data

# edit example config and define access tokens for writing clients
vim spaceapi.json.data

# build it!
go build

# run it!
./labapid
```

Usage
-----
`http` is a very useful tool to debug RESTful APIs. You can install it from your distro's repositories or with `pip install httpie`.

Query SpaceAPI information
```
http :5000 /
```

Update door status:
```
~ » cat json_input.json 
{
    "token": "foobar",
    "status": true
}
~ » http -j POST :5000/edit/door/ < json_input.json 
HTTP/1.1 200 OK
Content-Length: 16
Content-Type: text/plain; charset=utf-8
Date: Thu, 25 Sep 2014 23:33:41 GMT

{"success":true}
```

Update Sensor Status
```
~ » cat sensor_input.json                      
{
    "token": "foobar",
    "sensor": {
        "temperature": [
            {
                "value": 12,
                "unit": "°C"
            }
        ]
    }
}
~ » http -j POST :5000/edit/sensor/ < sensor_input.json    
HTTP/1.1 200 OK
Content-Length: 16
Content-Type: text/plain; charset=utf-8
Date: Thu, 25 Sep 2014 23:35:33 GMT

{"success":true}
```
