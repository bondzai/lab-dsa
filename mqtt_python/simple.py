import paho.mqtt.client as mqtt

def on_connect(client, userdata, flags, rc):
    print("Connected with result code "+str(rc))
    client.subscribe("JB")
    client.publish("JB", "Hello from python")

def on_message(client, userdata, msg):
    print("topic: " + msg.topic)
    print("string(bytes): " + str(msg.payload))
    print("string: " + msg.payload.decode())

client = mqtt.Client()
client.on_connect = on_connect
client.on_message = on_message

client.connect("localhost", 1883, 60)

client.loop_forever()

def on_message(client, userdata, message):
    print("Received message '" + str(message.payload) + "' on topic '" + message.topic + "'")

client.on_message = on_message

client.subscribe("test")