import mysql.connector
from mysql.connector import Error
import uuid

from db_conf.setup_mysql import mysql_conf

from queries.drop_tables import drop_all_tables
from queries.migrations import run_migrations
from queries.sql import Query

from Utils.Utils import Utils

from kafka_utils.setup_kafka import KafkaMsg
from kafka_utils.setup_kafka_consumer import KafkaConsumerScheduler

from Scrape.Lever import lever
from Scrape.GreenHouse import greenhouse
import time

def main():
    try:
        connection =  mysql.connector.connect(**mysql_conf()) 
        cursor = None
        if connection.is_connected():
            cursor = connection.cursor()
            print("cursor is connected!")
        else:
            print("cursor is not connected!")
        # drop_all_tables(cursor)
        # run_migrations(cursor)
        # return

        query = Query(cursor)
        utils = Utils()
        kafka = KafkaMsg(query)
        kafka.setup_protobuf_producer()
        kafka_consumer = KafkaConsumerScheduler(kafka, query)
        kafka_consumer.start()

        while True:
            companies = query.get_all_companies()
            time.sleep(4)
            for company_uuid, company_name, company_scrape_website, gh, lvr, oth in companies:
                time.sleep(0.4)
                if gh:
                    greenhouse(company_uuid, company_name, company_scrape_website, query, utils, kafka)
                elif lvr:
                    lever(company_uuid, company_name, company_scrape_website, query, utils, kafka)
                else:
                    print("oops?")
            connection.commit()
    except Error as e:
        print("Error while connecting to MySQL", e)
    finally:
        if (connection.is_connected()):
            cursor.close()
            connection.close()

if __name__ == "__main__":
    main()