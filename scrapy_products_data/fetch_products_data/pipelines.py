import pika
import pickle


class PutDataInMqPipeline(object):
    QUEUE_NAME = "product_data"

    def open_spider(self, spider):
        credentials = pika.PlainCredentials('rabbitmq', 'rabbitmq')
        parameters = pika.ConnectionParameters('localhost',
                                               6666,
                                               '/',
                                               credentials)
        # self.connection = pika.BlockingConnection(pika.URLParameters('amqp://guest:guest@localhost:5672/'))
        self.connection = pika.BlockingConnection(parameters)
        self.channel = self.connection.channel()
        self.channel.queue_declare(queue=self.QUEUE_NAME, durable=True)

    def close_spider(self, spider):
        self.connection.close()

    def process_item(self, item, spider):
        item['name'] = item['name'][0]
        item['price'] = int(item['price'][0].strip().replace(' ', ''))
        item['description'] = item['description'][1:]

        body = pickle.dumps(item)

        self.channel.basic_publish(exchange='',
                                   routing_key=self.QUEUE_NAME,
                                   body=body,
                                   properties=pika.BasicProperties(
                                       delivery_mode=2,
                                   ))
