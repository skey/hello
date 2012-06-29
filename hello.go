/*
  tools for check server:port available
  author: smallfish
*/

package main

import (
    "net"
    "fmt"
    "time"
    "flag"
    "strings"
    "sync"
    "os"
    "io/ioutil"
)

var (
    file        string
    limit       int
    timeout     int
    wg          sync.WaitGroup
    ch          chan int
)

func init() {
    flag.StringVar(&file, "file", "", "host:port file")
    flag.IntVar(&limit, "limit", 10, "limit concurrency, default: 10")
    flag.IntVar(&timeout, "timeout", 5, "connect timeout, default: 5")
}

func getAddrs(file string) ([]string) {
    byte, err := ioutil.ReadFile(file)
    if err != nil {
        return nil
    }
    var ret []string
    for _, val := range strings.Split(string(byte), "\n") {
        if len(val) > 2 {
            ret = append(ret, val)
        }
    }
    return ret
}

func work(addr string) {
    ch <- 1
    conn, err := net.DialTimeout("tcp", addr, time.Duration(timeout) * time.Second)
    if err != nil {
        fmt.Printf("%s\t\033[31mERROR\033[0m\n", addr) // red ERROR
    } else {
        conn.Close()
        fmt.Printf("%s\t\033[32mOK\033[0m\n", addr) // green OK
    }
    wg.Done()
    <-ch
}

func main() {
    flag.Usage = func() { flag.PrintDefaults() }
    flag.Parse()
    if file == "" {
        flag.PrintDefaults()
        os.Exit(0)
    }
    ch = make(chan int, limit)
    for _, addr := range getAddrs(file) {
        wg.Add(1)
        go work(addr)
    }
    wg.Wait()
}
