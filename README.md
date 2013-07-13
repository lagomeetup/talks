# Go, ØMQ and node

Demo showing patterns of using ØMQ with Go talking to other processes.

# Install

The installation instructions are written for Debian 7.0, though they should
work on other Debian based distros.

## Install ØMQ from src

```bash
$ # Download ØMQ source
$ wget http://download.zeromq.org/zeromq-3.2.3.tar.gz
$ # Untar the download
$ tar xf zeromq-3.2.3.tar.gz
$ cd zeromq-3.2.3
$ # Install build dependencies
$ sudo apt-get install -y libtool autoconf automake uuid-dev build-essential
$ # Do the install dance
$ ./configure
$ make
$ sudo make install
$ # This set may no be needed
$ echo "/usr/local/lib" | sudo tee /etc/ld.so.conf.d/zeromq.conf
```

## Install zeromq for go

There are a few for bindings for zeromq with Go. I'll be using
https://github.com/alecthomas/gozmq

You can Install gozmq by running:
`$ go get -tags zmq_3_x github.com/alecthomas/gozmq`

### Note on -tags

By sending the arguments `-tags zmq_3_x `, you are telling go to build zmq
for version 3.x instead of 2.x.

## Install zeromq for node.js

I'm using the zmq package from npm. To install the dependencies for the node
project, cd to the `node` directory and run `npm install`.

# Running Demo

## Build and run Go Server

Use `Ctrl+C` to exit.

```bash
$ # Go to server directory
$ cd go/server
$ go build
$ ./server
```

## Build Go Client

Use `Ctrl+C` to exit.

```bash
$ # Go to client directory
$ cd go/client/
$ go build
$ ./client
```

## Build Node Client

Use `Ctrl+C` to exit.

```bash
$ # Go to node client directory
$ cd node
$ npm install
$ npm run client
```

