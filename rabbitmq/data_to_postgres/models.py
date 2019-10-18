from .settings import psql_db
import peewee
from playhouse.postgres_ext import ArrayField, TextField


class Product(peewee.Model):
    name = peewee.CharField()
    price = peewee.IntegerField()
    description = ArrayField(field_class=TextField, dimensions=10)

    class Meta:
        database = psql_db
        db_table = 'products'
