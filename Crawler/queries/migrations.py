def run_migrations(cursor):
    cursor.execute("CREATE TABLE IF NOT EXISTS Companies \
                        ( UUID VARCHAR(255) NOT NULL, Name VARCHAR(255), \
                        Website VARCHAR(255), PRIMARY KEY(UUID) )")

    cursor.execute("CREATE TABLE IF NOT EXISTS Remembered_Jobs \
                        ( JobID VARCHAR(255) NOT NULL, \
                        Company_UUID VARCHAR(255) NOT NULL, \
                        ProvidedID VARCHAR(255), \
                        isActive TINYINT(1), PRIMARY KEY(JobID), \
                        FOREIGN KEY(Company_UUID) REFERENCES Companies(UUID) \
                        )")

    cursor.execute("CREATE TABLE IF NOT EXISTS Jobs \
                    ( JobID VARCHAR(255) NOT NULL, \
                    Company_UUID VARCHAR(255) NOT NULL, \
                    Joblink VARCHAR(255), DefaultLink VARCHAR(255), \
                    ProvidedID VARCHAR(255), CompanyName VARCHAR(255), \
                    Internship TINYINT(1), \
                    Entry TINYINT(1), Mid TINYINT(1), Senior TINYINT(1), \
                    Manager TINYINT(1), Active TINYINT(1), PRIMARY KEY(JobID), \
                    FOREIGN KEY(Company_UUID) REFERENCES Companies(UUID) )")
