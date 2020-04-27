import mysql.connector
from mysql.connector import Error
import uuid

from db_conf.setup_mysql import mysql_conf

from queries.drop_tables import drop_all_tables
from queries.migrations import run_migrations
from queries.sql import Query

from Utils.Utils import Utils

from kafka_utils.setup_kafka import KafkaMsg

from Scrape.Lever import lever
from Scrape.GreenHouse import greenhouse

def main():
    try:
        connection =  mysql.connector.connect(**mysql_conf()) 
        cursor = None
        if connection.is_connected():
            cursor = connection.cursor()
        # drop_all_tables(cursor)
        # run_migrations(cursor)
        # return

        query = Query(cursor)
        utils = Utils()
        kafka = KafkaMsg()
        kafka.setup_protobuf_producer()

        companies = query.get_all_companies()

        for company_uuid, company_name, company_scrape_website, gh, lvr, oth in companies:
            if gh:
                greenhouse(company_uuid, company_name, company_scrape_website, query, utils, kafka)
            elif lvr:
                lever(company_uuid, company_name, company_scrape_website, query, utils, kafka)
            else:
                print("oops?")
    except Error as e:
        print("Error while connecting to MySQL", e)
    finally:
        if (connection.is_connected()):
            cursor.close()
            connection.close()

if __name__ == "__main__":
    main()