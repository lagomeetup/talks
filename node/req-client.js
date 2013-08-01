var zmq = require('zmq');

sock = zmq.socket('req');

sock.connect('tcp://127.0.0.1:5000');

var count = 0;
setInterval(function(){
  console.log('sending work: ' + count);
  sock.send(count);
  count += 1;
}, 100);

