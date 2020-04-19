class Utils:
    def __init__(self):
        # Levels
        self.Manager = set(["manager"])
        self.Senior = set(["principle", "senior"])
        self.Mid = set(["mid"])
        self.Entry = set(["entry", "junior", "new", "grad", "associate"])
        self.Intern = set(["internship", "intern"])
    
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

        if "intern" in job_title:
            pass
        elif "manager" in job_title:
            experience_level = 5
        elif "senior" in job_title:
            experience_level = 4
        elif "mid" in job_title:
            experience_level = 3
        else:
            experience_level = 2

        return experience_level
