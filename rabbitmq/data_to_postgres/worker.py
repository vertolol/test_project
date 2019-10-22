from .db import postgres_database
from .models import Product
import settings


class PostgresWorker():
    def __init__(self):
        postgres_database.connect()

    def create_instance(self, name, price, description):
        self._create_table_if_not_exist()
        query = Product.select().where(Product.name == name)

        if query.exists():
            instance = query.get()
            self._update_instance(instance, price, description)
        else:
            instance = Product.create(
                name=name,
                price=price,
                description=description
            )
        return instance

    def _update_instance(self, pr, price, description):
        pr.price = price
        pr.description = description
        pr.save()

    def _create_table_if_not_exist(self):
        if settings.POSTGRES_TABLE_NAME not in postgres_database.get_tables():
            Product.create_table()
