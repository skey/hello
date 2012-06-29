Tools for check server:port available, write by Golang (test go1.0.2)

Install:

    $ git clone https://github.com/smallfish/hello.git
    $ cd hello
    $ go build hello.go

Usage:

    $ ./hello
      -file="": host:port file
      -limit=10: limit concurrency, default: 10
      -timeout=5: connect timeout, default: 5

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
