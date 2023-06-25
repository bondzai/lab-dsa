import pika
import message_pb2

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()

channel.queue_declare(queue='hello')

# Create a new MyMessage
my_message = message_pb2.MyMessage()
my_message.name = "Test"
my_message.id = 1234
my_message.email.append("test@example.com")

# Serialize the message to a byte string
message_bytes = my_message.SerializeToString()

# Publish the message to the queue
channel.basic_publish(exchange='',
                      routing_key='hello',
                      body=message_bytes)

print(" [x] Sent message")
connection.close()
