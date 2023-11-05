require('dotenv').config();
const mqtt = require('mqtt');

const options = {
    host: process.env.MQTT_HOST,
    port: parseInt(process.env.MQTT_PORT),
    protocol: process.env.MQTT_PROTOCOL,
    username: process.env.MQTT_USERNAME,
    password: process.env.MQTT_PASSWORD
};

const client = mqtt.connect(options);

client.on('connect', () => {
    console.log('Connected');
});

client.on('error', (error) => {
    console.error(error);
});

client.on('message', (topic, message) => {
    console.log('Received message:', topic, message.toString());
});

client.subscribe('my/test/topic');

client.publish('my/test/topic', 'Hello from JS client');
