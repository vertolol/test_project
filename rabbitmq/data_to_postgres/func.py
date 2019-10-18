from .models import Product
from .settings import psql_db


def pg_create_instance(name, price, description):
    check_table_exist()
    query = Product.select().where(Product.name == name)

    if query.exists():
        instance = query.get()
        update(instance, price, description)

    else:
        instance = Product.create(
            name=name,
            price=price,
            description=description
        )

    return instance


def update(pr, price, description):
    pr.price = price
    pr.description = description

    pr.save()


def pg_get_id(name):
    query = Product.select().where(Product.name == name)

    if query.exists():
        return query.get()
    else:
        return None


def check_table_exist():
    if 'products' in psql_db.get_tables():
        return None
    else:
        Product.create_table()
