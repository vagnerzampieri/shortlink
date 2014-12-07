###Sample shortlink in Go

###Use:

    go run server.go

    or

    ./shortlink

    curl -v "http://localhost:8888/api/shorten" -d "http://frelleto.com"

    2014/12/06 21:48:47 Initializer server in port 8888...
    2014/12/06 21:57:58 URL http://frelleto.com successfully shortened to http://localhost:8888/r/Xz+ss.
    
    curl http://localhost:8888/r/Xz+ss
    
    2014/12/07 14:26:24 Click successfully regitered for Xz+ss.
    
    curl http://localhost:8888/api/stats/Xz+ss
    
    {"url":{"id":"Xz+ss","criation":"2014-12-07T16:55:51.131866986Z","destination":"http://frelleto.com"},"clicks":0}
###Help:

    ./shortlink -h

    Usage of ./shortlink:
      -d="localhost": domain
      -l=true: log on/off
      -p=8888: port

###With my [Dockerfile][Dockerfile]

    $ sudo docker ps
    CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
    ef4c64eb792b        go:latest           "/bin/bash"         19 hours ago        Up 19 hours                             naughty_kirch

    $ sudo docker inspect go | grep IPAddress

    $ sudo docker inspect naughty_kirch | grep IPAddress
            "IPAddress": "172.17.0.4",

Access browser `http://172.17.0.4:8888/`

[Dockerfile]: https://github.com/vagnerzampieri/docker-files/tree/master/go
