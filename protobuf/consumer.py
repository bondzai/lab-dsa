import pika
import message_pb2

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()

channel.queue_declare(queue='hello')

def callback(ch, method, properties, body):
    # Create a new MyMessage and parse the byte string into it
    my_message = message_pb2.MyMessage()
    my_message.ParseFromString(body)
    print(" [x] Received message: ", my_message)

channel.basic_consume(queue='hello',
                      auto_ack=True,
                      on_message_callback=callback)

print(' [*] Waiting for messages. To exit press CTRL+C')
channel.start_consuming()
