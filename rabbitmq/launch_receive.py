import pika
import pickle
from time import sleep

from data_to_postgres.func import pg_create_instance
from data_to_elastic.func import elastic_create_document


sleep(30)
credentials = pika.PlainCredentials('rabbitmq', 'rabbitmq')
parameters = pika.ConnectionParameters(host="rabbit",
                                       port=5672,
                                       # virtual_host='/',
                                       credentials=credentials)
connection = pika.BlockingConnection(parameters)


channel = connection.channel()
channel.queue_declare(queue="product_data",
                      durable=True)


def callback(ch, method, properties, body):
    data = pickle.loads(body)

    instance_id = pg_create_instance(data['name'],
                                     data['price'],
                                     data['description'])
    print(f" [x] instance with id {instance_id} created on postgres")

    elastic_create_document(instance_id,
                            data['name'],
                            data['price'],
                            data['description'])
    print(f" [0] document with id {instance_id} created on elastic")

    ch.basic_ack(delivery_tag=method.delivery_tag)


channel.basic_qos(prefetch_count=1)
channel.basic_consume(queue="product_data",
                      on_message_callback=callback)

if __name__ == '__main__':
    print("start_receive")
    channel.start_consuming()
