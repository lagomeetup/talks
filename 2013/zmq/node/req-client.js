var zmq = require('zmq');

var sock = zmq.socket('req');

sock.connect('tcp://127.0.0.1:5000');

sock.on('message', function (data) {
  console.log('Got', data + '')
});

var count = 0;
setInterval(function(){
  console.log('Sending', count);
  sock.send(count);
  count += 1;
}, 100);

