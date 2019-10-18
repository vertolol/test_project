from elasticsearch_dsl import Document, Text, Integer, Keyword


class Product(Document):
    name= Text(fields={'raw': Keyword()})
    price = Integer()
    description = Text()

    class Index:
        name = "products"
