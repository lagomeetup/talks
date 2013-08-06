var http = require('http'),
    socketio = require('socket.io'),
    fs = require('fs'),
    zmq = require('zmq');

var pull = zmq.socket('pull'),
    app = http.createServer(handler),
    io = socketio.listen(app),
    indexHtml = fs.readFileSync(__dirname + '/index.html'),
    messages = ['Welcome!'];

function handler(req, res) {
  res.setHeader('Content-Type', 'text/html');
  res.writeHead(200);
  res.end(indexHtml);
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

