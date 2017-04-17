# Go simple redis service

Simple web service for connecting to Redis and testing basic CRUD actions for [Microlib](https://github.com/microlib).
Uses a generic script `script.sh` to simulate start & stop for Linux & MacOS based systems. 

## Usage 

```
# cd to project directory and build executable
$ go build .
$ chmod u+x script.sh

# start the service
$ ./script.sh start

# stop the service
$ ./script.sh stop
```

Replace the `EXEC` variable in `script.sh` with the name of your executable, if it's different.

## Note
The http server uses signals to allow for graceful shutdown. Use this as a standard pattern when creating all web services. 