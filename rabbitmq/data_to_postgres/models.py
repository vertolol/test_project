import peewee
from .db import postgres_database
from playhouse.postgres_ext import ArrayField, TextField
import settings


class BaseModel(peewee.Model):
    class Meta:
        database = postgres_database


class Product(BaseModel):
    name = peewee.CharField()
    price = peewee.IntegerField()
    description = ArrayField(field_class=TextField, dimensions=10)

    class Meta:
        db_table = settings.POSTGRES_TABLE_NAME
