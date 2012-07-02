"""
  tools for check server:port available, require gevent(use coroutine)
  author: smallfish
"""
from sys import exit
from optparse import OptionParser
from itertools import izip_longest
import gevent
from gevent import socket

def gen_groups(array, n):
    args = [iter(array)] * n
    return list(izip_longest(*args))

def get_connected(addr, port):
    try:
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.connect((addr, port))
        sock.close()
        return addr, port, True
    except:
        return addr, port, False

def main():
    parser = OptionParser(usage='%prog --file=urlfile --limit=1000 --timeout=5')
    parser.add_option('', '--file', help='host:port file')
    parser.add_option('', '--limit', default=10, type="int", help='limit concurrency, default: %default')
    parser.add_option('', '--timeout',  default=5, type="int", help='connect timeout, default: %default')
    (options, args) = parser.parse_args()
    if not options.file:
        parser.print_help()
        exit(0)
    socket.setdefaulttimeout(options.timeout)
    groups = gen_groups([line.strip() for line in open(options.file) if line], options.limit)
    for group in groups:
        jobs = []
        for item in group:
            if item:
                addr, port = item.split(":")
                jobs.append(gevent.spawn(get_connected, addr, int(port)))
        gevent.joinall(jobs)
        for job in jobs:
            addr, port, connected = job.value
            if connected:
                print "%s:%d\tOK" % (addr, port)
            else:
                print "%s:%d\tERROR" % (addr, port)

if __name__ == '__main__':
    main()
