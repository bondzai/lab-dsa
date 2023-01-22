const express = require('express');
const http = require('http');
const socketio = require('socket.io');

const app = express();
const server = http.createServer(app);
const io = socketio(server);

io.on('connection', (socket) => {
  console.log('Client connected');

  socket.on('message', (data) => {
    try {
      // Parse the incoming data as JSON
      const message = JSON.parse(data);
      console.log(`Received message => ${message}`);
      socket.emit('message', { response: `Echo: ${message.message}` });
    } catch (err) {
      console.log(`Error parsing JSON: ${err}`);
    }
  });

  socket.on('disconnect', () => {
    console.log('Client disconnected');
  });
});

server.listen(8000, () => {
  console.log('Server started on port 8000');
});
