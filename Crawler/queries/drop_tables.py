def drop_all_tables(cursor):
    cursor.execute("SET FOREIGN_KEY_CHECKS = 0;")
    cursor.execute("DROP TABLE IF EXISTS Companies")
    cursor.execute("DROP TABLE IF EXISTS Remembered_Jobs;")
    cursor.execute("DROP TABLE IF EXISTS Jobs;")
    cursor.execute("SET FOREIGN_KEY_CHECKS = 1;")