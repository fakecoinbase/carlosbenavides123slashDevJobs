import mysql.connector
from mysql.connector import Error
import uuid

from db_conf.setup_mysql import mysql_conf

from queries.drop_tables import drop_all_tables
from queries.migrations import run_migrations
from queries.sql import Query

from Utils.Utils import Utils

from kafka_utils.setup_kafka import KafkaMsg

from Scrape.Honey import Honey

def addHoney(cursor):
    UUID = str(uuid.uuid4())
    Name = "Honey"
    Website = "https://api.greenhouse.io/v1/boards/honey/departments"
    SQL = "INSERT INTO Companies (UUID, Name, Website) VALUES (%s, %s, %s)"
    cursor.execute(SQL, (UUID, Name, Website))


def main():
    try:
        connection =  mysql.connector.connect(**mysql_conf()) 
        cursor = None
        if connection.is_connected():
            cursor = connection.cursor()

        if not cursor:
            print("LOG: ERROR CURSOR ISN'T ACTIVE")
        else:
            print("!!!!!!!!!!!!!!!!1")

        drop_all_tables(cursor)
        run_migrations(cursor)
        addHoney(cursor)

        query = Query(cursor)
        utils = Utils()
        kafka = KafkaMsg()
        kafka.setup_json_producer()

        companies = query.get_all_companies()

        for UUID, Name, Website in companies:
            Honey(UUID, Name, Website, query, utils, kafka)
    except Error as e:
        print("Error while connecting to MySQL", e)
    finally:
        if (connection.is_connected()):
            cursor.close()
            connection.close()

if __name__ == "__main__":
    main()