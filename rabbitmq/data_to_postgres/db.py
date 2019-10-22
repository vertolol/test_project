from playhouse.postgres_ext import PostgresqlDatabase
import settings


postgres_database = PostgresqlDatabase(database=settings.POSTGRES_DB_NAME,
                                       user=settings.POSTGRES_USER_NAME,
                                       password=settings.POSTGRES_PASSWORD,
                                       host=settings.POSTGRES_HOST_NAME,
                                       port=settings.POSTGRES_PORT
                                       )