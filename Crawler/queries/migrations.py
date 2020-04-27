def run_migrations(cursor):
    cursor.execute("CREATE TABLE IF NOT EXISTS companies \
                        ( company_uuid VARCHAR(255) NOT NULL, \
                        company_name VARCHAR(255), \
                        company_scrape_website VARCHAR(255), \
                        greenhouse TINYINT(1), \
                        lever TINYINT(1), \
                        other TINYINT(1), \
                        PRIMARY KEY(company_uuid) \
                        )")

    cursor.execute("CREATE TABLE IF NOT EXISTS remembered_jobs \
                        ( job_id VARCHAR(255) NOT NULL, \
                        company_uuid VARCHAR(255) NOT NULL, \
                        provided_id VARCHAR(255), \
                        active TINYINT(1), PRIMARY KEY(job_id), \
                        FOREIGN KEY(company_uuid) REFERENCES companies(company_uuid) \
                        )")

    cursor.execute("CREATE TABLE IF NOT EXISTS jobs \
                    ( job_id VARCHAR(255) NOT NULL, \
                    company_uuid VARCHAR(255) NOT NULL, \
                    job_link VARCHAR(255), default_link VARCHAR(255), \
                    provided_id VARCHAR(255), company_name VARCHAR(255), \
                    experience_level int(12), \
                    active TINYINT(1), PRIMARY KEY(job_id), \
                    FOREIGN KEY(company_uuid) REFERENCES companies(company_uuid) \
                    )")
    cursor.execute("CREATE TABLE IF NOT EXISTS locations \
                    ( location VARCHAR(255) NOT NULL, \
                    company_name VARCHAR(255) NOT NULL \
                    )")

    # TODO
    # cursor.execute("CREATE TABLE IF NOT EXISTS company_scrape_details \
    #                 ( company_uuid VARCHAR(255) NOT NULL, \
    #                 company_name VARCHAR(255) NOT NULL, \
    #                 wanted_departments VARCHAR(255), \
    #                 wanted_locations VARCHAR(255), \
    #                 PRIMARY KEY(company_uuid), \
    #                 FOREIGN KEY(company_uuid) REFERENCES companies(company_uuid), \
    #                 ")