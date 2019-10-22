from .models import Product
from elasticsearch import Elasticsearch
import settings


class ElasticWorker():
    def __init__(self):
        host = "http://{}:{}".format(settings.ELASTIC_HOST_NAME,
                                     settings.ELASTIC_PORT)
        self.client = Elasticsearch(host)

    def create_document(self, instance_id, name, price, description):
        new_doc = Product(
            name=name,
            price=price,
            description=description
        )

        new_doc.meta.id = instance_id
        new_doc.save(using=self.client,
                     index=settings.ELASTIC_INDEX_NAME)
