from elasticsearch_dsl import Document, Text, Integer, Keyword
import settings


class Product(Document):
    name= Text(fields={'raw': Keyword()})
    price = Integer()
    description = Text()

    class Index:
        name = settings.ELASTIC_INDEX_NAME
