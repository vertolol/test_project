from .models import Product
from .settings import client


def elastic_create_document(instance_id, name, price, description):
    new_doc = Product(
        name=name,
        price=price,
        description=description
    )

    new_doc.meta.id = instance_id
    new_doc.save(using=client, index='products')


# from elasticsearch_dsl import Search
# from elasticsearch_dsl.query import Match
# def search_by_name(name):
#     query = Match(name={"query": name, "operator": "and"})
#     result = Search(using=client, index='products').query(query)
#
#     return result.execute()
