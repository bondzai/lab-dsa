const mqtt = require('mqtt');
const client = mqtt.connect('mqtt://127.0.0.1');

client.on('connect', function() {
    client.subscribe('JB', function(error) {
        if (!error) {
            client.publish('JB', 'Hello from node.js');
        };
    });
});

client.on('message', function(topic, message) {
    console.log(message.toString());
    client.end();
});