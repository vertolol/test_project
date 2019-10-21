import peewee
from playhouse.postgres_ext import PostgresqlDatabase
from playhouse.postgres_ext import ArrayField, TextField
import settings


postgres_database = PostgresqlDatabase(database=settings.POSTGRES_DB_NAME,
                             user=settings.POSTGRES_USER_NAME,
                             password=settings.POSTGRES_PASSWORD,
                             host=settings.POSTGRES_HOST_NAME,
                             port=settings.POSTGRES_PORT
                             )


class BaseModel(peewee.Model):
    class Meta:
        database = postgres_database


class Product(BaseModel):
    name = peewee.CharField()
    price = peewee.IntegerField()
    description = ArrayField(field_class=TextField, dimensions=10)

    class Meta:
        db_table = settings.POSTGRES_TABLE_NAME
