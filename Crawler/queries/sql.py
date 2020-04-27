class Query(object):
    def __init__(self, cursor):
        self.cursor = cursor
    
    def get_all_companies(self):
        SQL_SELECT_ALL_COMPANIES = "SELECT * FROM companies"
        self.cursor.execute(SQL_SELECT_ALL_COMPANIES)
        return self.cursor.fetchall()

    def get_active_remembered_jobs(self, company_uuid):
        SQL_SELECT_ALL_ACTIVE = "SELECT * FROM remembered_jobs WHERE company_uuid=%s AND active=%s"
        self.cursor.execute(SQL_SELECT_ALL_ACTIVE, (company_uuid, 1))
        return self.cursor.fetchall()

    def check_active_job(self, job_id, company_uuid):
        SQL_CHECK_REMEMBERED = "SELECT active FROM remembered_jobs WHERE job_id=%s AND company_uuid=%s"
        self.cursor.execute(SQL_CHECK_REMEMBERED, (job_id, company_uuid))
        return self.cursor.fetchall()

    def insert_new_job(self, job):
        SQL_INSERT_INTO_JOBS = "INSERT INTO jobs \
                                (job_id, company_uuid, job_link, \
                                default_link, provided_id, company_name, experience_level, \
                                active) VALUES (%s, %s, %s, %s, %s, %s, %s, %s)"
        self.cursor.execute(SQL_INSERT_INTO_JOBS, (job.JobUUID, job.CompanyUUID, job.JobLink,
                            job.DefaultLink, job.ProvidedID, job.CompanyName, job.ExperienceLevel,
                            job.Active))

    def insert_new_remembered_job(self, job):
        SQL_INSERT_INTO_REMEMBERED_JOB = "INSERT INTO \
                                            remembered_jobs \
                                            (job_id, \
                                            company_uuid, \
                                            provided_id, \
                                            active) VALUES (%s, %s, %s, %s)"
        self.cursor.execute(SQL_INSERT_INTO_REMEMBERED_JOB, (job.JobUUID, 
                            job.CompanyUUID, job.ProvidedID, job.Active))

    def deactivate_job(self, job_id, provided_id):
        SQL_SOFT_DELETE_JOB = "UPDATE jobs SET active = %s WHERE job_id = %s AND provided_id = %s"
        self.cursor.execute(SQL_SOFT_DELETE_JOB, (0, job_id, provided_id,))
        SQL_SOFT_DELETE_JOB = "UPDATE remembered_jobs SET active = %s WHERE job_id = %s AND provided_id = %s"
        self.cursor.execute(SQL_SOFT_DELETE_JOB, (0, job_id, provided_id,))

    def check_location_company(self, location, company_name):
        location_sql = "SELECT location FROM locations WHERE location=%s AND company_name=%s"
        self.cursor.execute(location_sql, (location, company_name))
        results = self.cursor.fetchall()
        return len(results) == 0

    def insert_new_location(self, location, company_name):
        location_sql = "INSERT INTO locations (location, company_name) VALUES (%s, %s)"
        self.cursor.execute(location_sql, (location, company_name))
