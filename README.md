
How to View the Presentation
============================
This presentation uses the go.talks presenter tools.
To view on your local machine you need to have installed go already. (1.0.3 at the moment of this writing november 20 2012)

 $ git co git@github.com:lagomeetup/talks.git talks
 $ cd talks
 $ export GOPATH=`pwd`
 $ go get code.google.com/p/go.talks/present/
 $ ./bin/present -play=true -base= ./src/code.google.com/p/go.talks/

open your browser (chrome) to http://127.0.0.1:3999/
click on 2012 -> go-stdlib-part-1-io.slide#1

http://127.0.0.1:3999/2012/go-stdlib-part-1-io.slide#1

Click on the links to go directly to the online google documentation (internet required)
Push run the test code locally (make sure the play=true flag is set when launching the present tool)
The full source code is located inside the subdirectories. The example should be fully functional

Enjoy! 

-- The Los Angeles Go Meetup Team