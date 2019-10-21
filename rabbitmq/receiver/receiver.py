from data_to_postgres.func import Postgres
from data_to_elastic.func import Elastic
import settings
import pika
import pickle
from time import sleep


class RabbitMqWorker():
    def __init__(self):
        self.connection = self._create_connection()
        self.channel = self.connection.channel()
        self.channel.queue_declare(queue=settings.RABBIT_QUEUE_NAME,
                                   durable=True)

        self.channel.basic_qos(prefetch_count=1)
        self.channel.basic_consume(queue=settings.RABBIT_QUEUE_NAME,
                                   on_message_callback=self.callback)

        self.elastic = Elastic()
        self.postgres = Postgres()

    def _create_connection(self):
        credentials = pika.PlainCredentials(username=settings.RABBIT_USERNAME,
                                            password=settings.RABBIT_PASSWORD)
        parameters = pika.ConnectionParameters(host=settings.RABBIT_HOST,
                                               port=settings.RABBIT_PORT,
                                               credentials=credentials)
        return pika.BlockingConnection(parameters)

    def callback(self, ch, method, properties, body):
        data = pickle.loads(body)

        instance_id = self.postgres.create_instance(data['name'],
                                                    data['price'],
                                                    data['description'])
        print(f" [x] instance with id {instance_id} created on postgres")

        self.elastic.create_document(instance_id,
                                     data['name'],
                                     data['price'],
                                     data['description'])
        print(f" [0] document with id {instance_id} created on elastic")

        ch.basic_ack(delivery_tag=method.delivery_tag)


    def start_receive(self):
        print("start_receive")
        self.channel.start_consuming()


if __name__ == '__main__':
    sleep(30)
    RabbitMqWorker().start_receive()
