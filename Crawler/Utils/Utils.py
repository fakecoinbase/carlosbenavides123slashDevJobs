class Utils:
    def __init__(self):
        # Levels
        self.Manager = ["manager", "lead", "ceo", "chief", "officer"]
        self.Senior = ["principal", "senior", "sr."]
        self.Mid = ["mid", "engineerii", "engineeriii", "engineer2", "engineer3"]
        self.Entry = ["entry", "junior", "new", "grad", "associate", "engineeri", "engineer1"]
        self.Intern = ["intern", "internship"]
    
    def convert_active_jobs_to_dict(self, active_jobs):
        if not active_jobs:
            return {}
        hmap = {}
        for job_id, company_uuid, provided_id, active in active_jobs:
            if job_id in hmap:
                print("LOG", job_id, " in hmap??")
            else:
                hmap[job_id] = [job_id, company_uuid, provided_id]
        return hmap
    
    def determine_experience_level(self, job_title):
        experience_level = 1
        job_title = job_title.strip()
        job_title = job_title.lower()
        job_title = job_title.replace(" ", "")

        for name in self.Intern:
            if name in job_title:
                return 1
        for name in self.Mid:
            if name in job_title:
                return 3
        for name in self.Senior:
            if name in job_title:
                return 4
        for name in self.Manager:
            if name in job_title:
                return 5
        # default entry
        return 2
