# Go simple redis service

Simple web service for connecting to Redis and testing basic CRUD actions for [Microlib](https://github.com/microlib).
Uses a generic script `script.sh` to simulate start & stop for Linux & MacOS based systems. 

## Usage 
Replace the `EXEC` variable in `script.sh` with the name of your executable, if it's different.
Update the `config.json` with your redis host and port settings, then:

```bash
# cd to project directory and build executable
$ go build .
$ chmod u+x script.sh

# start the service
$ ./script.sh start

# stop the service
$ ./script.sh stop
```

To check, execute the following commands (change your host and port settings accordingly)
```bash
# Get values from redis
$ curl http://localhost:9000 

# Create value 
$ curl http://localhost:9000/create/2

# Update value and check
$ curl http://localhost:9000/update/6
$ curl http://localhost:9000 

# Delete value and check
$ curl http://localhost:9000/delete/6
$ curl http://localhost:9000 
```

## Note
The http server uses signals to allow for graceful shutdown. Use this as a standard pattern when creating all web services. 
