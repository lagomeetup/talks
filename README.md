To View Online this Presentation go to:
=======================================

2013
----
* [Whispering Gophers](http://whispering-gophers.appspot.com/talk.slide#1)
  workshop
* [ØMQ](http://talks.godoc.org/github.com/lagomeetup/talks/2013/zmq.slide)
* [Best Practices](http://talks.golang.org/2013/bestpractices.slide)
* [DVID](http://talks.godoc.org/github.com/lagomeetup/talks/2013/dvid.slide)
* [nsrsc](http://talks.godoc.org/github.com/lagomeetup/talks/2013/nrsc.slide)
* [Writing a file system in Go](http://bazil.org/talks/2013-06-10-la-gophers/)
  / [bazil.org/fuse](http://bazil.org/fuse/): a pure-Go FUSE library
* [Uncomment Your Code](https://docs.google.com/presentation/d/1iiPWo1zJRkk8siX-Yj1qwDSu9WWjbJkOJEBn9XjP18Q/pub?start=false&loop=false&delayms=3000)
* [Testing, Debugging and Profiling](http://talks.godoc.org/github.com/lagomeetup/talks/2013/test-debug-prof.slide)
* [import "C"](http://talks.godoc.org/github.com/lagomeetup/talks/2013/import-c.slide)
* [stdlib part 3 - reflection](http://talks.godoc.org/github.com/lagomeetup/talks/2013/go-stdlib-part3-reflection.slide)
* [stdlib part 2 - io & concurrency](http://talks.godoc.org/github.com/lagomeetup/talks/2013/go-stdlib-part2-io-concurrency.slide)

2012
----
* [stdlib part 1 - io](http://talks.godoc.org/github.com/lagomeetup/talks/2012/go-stdlib-part-1-io.slide)

How to View the Presentation
============================

This presentation uses the go.talks presenter tools.
To view on your local machine you need to have installed go already.
 (1.0.3 at the moment of this writing november 20 2012):

Install
--------

    git co git@github.com:lagomeetup/talks.git talks
    cd talks
    export GOPATH=`pwd`
    go get code.google.com/p/go.talks/present/
    ./bin/present -play=true -base= ./src/code.google.com/p/go.talks/

* Open your browser (chrome) to http://127.0.0.1:3999/
* Go to one of the year (say 2013) and click on of the `.slide` files

Click on the links to go directly to the online google documentation (internet required)

Push run the test code locally (make sure the play=true flag is set when launching the present tool)

The full source code is located inside the subdirectories. 

The example should be fully functional.

Enjoy!!

[*The Los Angeles Go Meetup Team*](http://www.meetup.com/Los-Angeles-Gophers/)
