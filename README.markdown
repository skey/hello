Tools for check server:port available, write by golang (test go1.0.2), python (test python2.6, require gevent)

Install:

    $ git clone https://github.com/smallfish/hello.git
    $ cd hello
    $ go build hello.go

    or use "go get":
    $ go get github.com/smallfish/hello
    $ ls $GOPATH/bin/hello

Usage:

    $ ./hello
      -file="": host:port file
      -limit=10: limit concurrency, default: 10
      -timeout=5: connect timeout, default: 5

    $ python hello.py 
      Usage: hello.py --file=urlfile --limit=1000 --timeout=5

      Options:
        -h, --help         show this help message and exit
        --file=FILE        host:port file
        --limit=LIMIT      limit concurrency, default: 10
        --timeout=TIMEOUT  connect timeout, default: 5

Example:

    $ cat urlfile
    twitter.com:80
    youtube.com:80
    qq.com:80
    dev.twitter.com:80
    163.com:80
    google.com:80
    sohu.com:80

    $ ./hello -file=urlfile 
    163.com:80      OK
    dev.twitter.com:80      OK
    google.com:80   OK
    qq.com:80       OK
    sohu.com:80     OK
    twitter.com:80  ERROR
    youtube.com:80  ERROR

    $ python hello.py --file=urlfile --limit=1000 --timeout=2
    twitter.com:80  ERROR
    youtube.com:80  ERROR
    qq.com:80       OK
    dev.twitter.com:80      ERROR
    163.com:80      OK
    google.com:80   OK
    sohu.com:80     OK
