###Sample shortlink in Go

###Use:

    go run server.go

    or

    ./shortlink

    curl -v "http://localhost:8888/api/shorten" -d "http://frelleto.com"

    2014/12/06 21:48:47 Initializer server in port 8888...
    2014/12/06 21:57:58 URL http://frelleto.com successfully shortened to http://localhost:8888/r/Xz+ss.
###Help:

    ./shortlink -h

    Usage of ./shortlink:
      -d="localhost": domain
      -l=true: log on/off
      -p=8888: port
