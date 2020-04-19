def run_migrations(cursor):
    cursor.execute("CREATE TABLE IF NOT EXISTS companies \
                        ( company_uuid VARCHAR(255) NOT NULL, \
                        company_name VARCHAR(255), \
                        company_scrape_website VARCHAR(255), \
                        PRIMARY KEY(company_uuid) )")
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
                    FOREIGN KEY(company_uuid) REFERENCES companies(company_uuid))")
