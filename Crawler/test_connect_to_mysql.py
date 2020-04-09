import mysql.connector
from mysql.connector import Error
import uuid
from Scrape.Honey import Honey
from Queries.SQL import Query
from Utils.Utils import Utils
from Queries.Migrations import run_migrations
from Queries.DropTables import drop_all_tables

# temp
def addHoney(cursor):
    UUID = str(uuid.uuid4())
    Name = "Honey"
    Website = "https://api.greenhouse.io/v1/boards/honey/departments"
    SQL = "INSERT INTO Companies (UUID, Name, Website) VALUES (%s, %s, %s)"
    cursor.execute(SQL, (UUID, Name, Website))


try:
    connection_config_dict = {
        'user': 'root',
        'password': 'fasd1423f',
        'host': '127.0.0.1',
        'database': 'jobs',
        'raise_on_warnings': False,
        'use_pure': False,
        'autocommit': True,
        'pool_size': 5
    }
    connection = mysql.connector.connect(**connection_config_dict)
    if connection.is_connected():
        cursor = connection.cursor()

        drop_all_tables(cursor)
        run_migrations(cursor)
        addHoney(cursor)

        query = Query(cursor)
        utils = Utils()

        companies = query.get_all_companies()

        for UUID, Name, Website in companies:
            Honey(UUID, Name, Website, query, utils)

except Error as e:
    print("Error while connecting to MySQL", e)
finally:
    if (connection.is_connected()):
        cursor.close()
        connection.close()
        print("MySQL connection is closed")
