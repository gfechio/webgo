# GOLang projects

Web app in Golang

## Getting Started

 - Serving port 8080 to serve with with NGINX 
 -  Parse OSM.PBF file
 -  Preprocessing

### Prerequisites

 - Would like to serve that in a container 
 - Golang 
 - osm.pbf file from the alps -> http://download.geofabrik.de/europe.html
 - Docker ( Using a redis instance on a docker container)

### Installing

```
$ go get github.com/go-redis/redis
$ go get github.com/qedus/osmpbf
$ docker pull redis ; docker run -p 6379:6379 redis &

```

## Running the tests

WEB:
```
http://127.0.0.1:8080/test

http://127.0.0.1:8080/maps/maps

```

Parser
```
To check Nodes
$ go run redisClient.go Node:120

To check Ways
$ go run redisClient.go Way:12345
```

## Built With

Container and Golang 
Need to drop the libraries here


## Authors

* **Giordano Fechio** - *Initial work* 

## License

Copyright [2017] [Giordano Dezute Fechio]

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

## Acknowledgments


