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
from Scrape.Twilio import Twilio
from Scrape.Asana import Asana

def addHoney(cursor):
    UUID = "50b3dae9-0bec-456f-af6d-61a8fabe0935"
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
        # drop_all_tables(cursor)
        # run_migrations(cursor)
        # addHoney(cursor)

        query = Query(cursor)
        utils = Utils()
        kafka = KafkaMsg()
        kafka.setup_protobuf_producer()

        companies = query.get_all_companies()

        for company_uuid, company_name, company_scrape_website in companies:
            if company_name == "Honey":
                Honey(company_uuid, company_name, company_scrape_website, query, utils, kafka)
            if company_name == "Twilio":
                Twilio(company_uuid, company_name, company_scrape_website, query, utils, kafka)
            if company_name == "Asana":
                Asana(company_uuid, company_name, company_scrape_website, query, utils, kafka)


    except Error as e:
        print("Error while connecting to MySQL", e)
    finally:
        if (connection.is_connected()):
            cursor.close()
            connection.close()

if __name__ == "__main__":
    main()