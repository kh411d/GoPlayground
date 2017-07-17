To build the application (let's call it queued), run the following command:

$ go build -o queued *.go

Now, to run our application, how about we start it with 2048 workers, just for kicks?

$ ./queued -n 2048
...
Starting worker 2047
Starting worker 2048
Registering the collector
HTTP server listening on 127.0.0.1:8000

Sweet! Now, in another terminal window, let's write a little Bash one-liner, to flood our collector with requests:

$ for i in {1..4096}; do curl localhost:8000/work -d name=$USER -d delay=$(expr $i % 11)s; done