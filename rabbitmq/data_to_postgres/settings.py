from playhouse.postgres_ext import PostgresqlExtDatabase


psql_db = PostgresqlExtDatabase("gorm",
                                user='gorm',
                                password="gorm",
                                host="postgres_host",
                                port="5432"
                                )
