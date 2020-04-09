class Query(object):
    def __init__(self, cursor):
        self.cursor = cursor
    
    def get_all_companies(self):
        SQL_SELECT_ALL_COMPANIES = "SELECT * FROM Companies"
        self.cursor.execute(SQL_SELECT_ALL_COMPANIES)
        return self.cursor.fetchall()

    def get_active_remembered_jobs(self, company_uuid):
        SQL_SELECT_ALL_ACTIVE = "SELECT * FROM Remembered_Jobs WHERE Company_UUID=%s AND isActive=%s"
        self.cursor.execute(SQL_SELECT_ALL_ACTIVE, (company_uuid, 1))
        return self.cursor.fetchall()

    def check_active_job(self, job_id, company_uuid):
        SQL_CHECK_REMEMBERED = "SELECT isActive FROM Remembered_Jobs WHERE JobID=%s AND Company_UUID=%s"
        self.cursor.execute(SQL_CHECK_REMEMBERED, (job_id, company_uuid))
        return self.cursor.fetchall()

    def insert_new_job(self, values):
        SQL_INSERT_INTO_JOBS = "INSERT INTO Jobs \
                                (JobID, Company_UUID, Joblink, \
                                DefaultLink, ProvidedID, Internship, Entry, \
                                Mid, Senior, Manager, Active) \
                                VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"
        # print(SQL_INSERT_INTO_JOBS)
        self.cursor.execute(SQL_INSERT_INTO_JOBS, values)

    def insert_new_remembered_job(self, values):
        SQL_INSERT_INTO_REMEMBERED_JOB = "INSERT INTO \
                                            Remembered_Jobs \
                                            (JobID, \
                                            Company_UUID, \
                                            ProvidedID, \
                                            isActive) VALUES (%s, %s, %s, %s)"
        self.cursor.execute(SQL_INSERT_INTO_REMEMBERED_JOB, values)

    def deactivate_job(self, job_id, company_uuid, provided_id):
        SQL_SOFT_DELETE_JOB = "UPDATE %s \
                                SET Active=%s \
                                WHERE JobID=%s, \
                                Company_UUID=%s \
                                ProvidedID=%s"
        job_table_data = ("Jobs", 0, job_id, company_uuid, provided_id)
        self.cursor.execute(SQL_SOFT_DELETE_JOB, job_table_data)
        rememb_table_data = ("Remembered Jobs", 0, job_id, company_uuid)
        self.cursor.execute(SQL_SOFT_DELETE_JOB, rememb_table_data, provided_id)

