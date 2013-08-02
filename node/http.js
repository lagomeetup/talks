var http = require('http'),
    socketio = require('socket.io'),
    fs = require('fs'),
    zmq = require('zmq');

var pull = zmq.socket('pull'),
    app = http.createServer(handler),
    io = socketio.listen(app),
    messages = ['Welcome!'];

function handler(req, res) {
  fs.readFile(__dirname + '/index.html', function (err, data) {
    if (err) {
      res.writeHead(500);
      return res.end('Error loading index.html');
    }

    res.setHeader('Content-Type', 'text/html');
    res.writeHead(200);
    res.end(data);
  });
}

pull.bind('tcp://127.0.0.1:4000');

app.listen(8888);

pull.on('message', function(msg) {
  messages.push(msg.toString());
});

io.sockets.on('connection', function (socket) {
  function emitMessage (msg) {
    socket.emit('news', { message: msg });
  }
  messages.forEach(emitMessage);
  pull.on('message', function (msg) {
    console.log(msg.toString());
    emitMessage(msg.toString());
  });
  socket.on('my other event', function (data) {
    console.log(data);
  });
});

