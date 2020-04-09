import mysql.connector
from mysql.connector import Error
import uuid
from Scrape.Honey import Honey
from Queries.SQL import Query
from Utils.Utils import Utils

def runMigrations(cursor):
    cursor.execute("CREATE TABLE IF NOT EXISTS Companies ( UUID VARCHAR(255) NOT NULL, Name VARCHAR(255), Website VARCHAR(255), PRIMARY KEY(UUID) )")
    cursor.execute("CREATE TABLE IF NOT EXISTS Remembered_Jobs ( JobID VARCHAR(255) NOT NULL, Company_UUID VARCHAR(255) NOT NULL, isActive TINYINT(1), PRIMARY KEY(JobID), FOREIGN KEY(Company_UUID) REFERENCES Companies(UUID) )")
    cursor.execute("CREATE TABLE IF NOT EXISTS Jobs ( JobID VARCHAR(255) NOT NULL, Company_UUID VARCHAR(255) NOT NULL, Joblink VARCHAR(255), DefaultLink VARCHAR(255), ProvidedID VARCHAR(255), Internship TINYINT(1), Entry TINYINT(1), Mid TINYINT(1), Senior TINYINT(1), Manager TINYINT(1), Active TINYINT(1), PRIMARY KEY(JobID), FOREIGN KEY(Company_UUID) REFERENCES Companies(UUID) )")


def dropAllTables(cursor):
    cursor.execute("SET FOREIGN_KEY_CHECKS = 0;")
    cursor.execute("DROP TABLE IF EXISTS Companies")
    cursor.execute("DROP TABLE IF EXISTS Remembered_Jobs;")
    cursor.execute("DROP TABLE IF EXISTS Jobs;")
    cursor.execute("SET FOREIGN_KEY_CHECKS = 1;")

# temp
def addHoney(cursor):
    UUID = str(uuid.uuid4())
    Name = "Honey"
    Website = "https://api.greenhouse.io/v1/boards/honey/departments"
    SQL = "INSERT INTO Companies (UUID, Name, Website) VALUES (%s, %s, %s)"
    cursor.execute(SQL, (UUID, Name, Website))
    # cursor.commit()


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

        dropAllTables(cursor)
        runMigrations(cursor)
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
