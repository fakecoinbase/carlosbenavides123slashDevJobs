def drop_all_tables(cursor):
    cursor.execute("SET FOREIGN_KEY_CHECKS = 0;")
    cursor.execute("DROP TABLE IF EXISTS companies")
    cursor.execute("DROP TABLE IF EXISTS remembered_jobs;")
    cursor.execute("DROP TABLE IF EXISTS jobs;")
    cursor.execute("SET FOREIGN_KEY_CHECKS = 1;")
    cursor.execute("DROP TABLE IF EXISTS locations")