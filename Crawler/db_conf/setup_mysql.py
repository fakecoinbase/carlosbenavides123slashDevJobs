def mysql_conf():
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
    return connection_config_dict

        # if connection.is_connected():
    #     #     cursor = connection.cursor()
    #     #     print(cursor)
    #     #     return cursor

    #         # drop_all_tables(cursor)
    #         # run_migrations(cursor)
    #         # addHoney(cursor)

    #         # query = Query(cursor)
    #         # utils = Utils()

    #         # companies = query.get_all_companies()

    #         # for UUID, Name, Website in companies:
    #         #     Honey(UUID, Name, Website, query, utils)
    # except Error as e:
    #     print("Error while connecting to MySQL", e)
    # finally:
    #     if connection.is_connected():
    #         cursor = connection.cursor()
    #         print("###############333")
    #         return cursor
    #     return None
