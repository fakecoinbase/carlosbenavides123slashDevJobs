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
        intersect_manager = job_title.intersection(self.Manager)
        intersect_senior = job_title.intersection(self.Senior)
        intersect_mid = job_title.intersection(self.Mid)
        intersect_entry = job_title.intersection(self.Entry)
        intersect_intern = job_title.intersection(self.Intern)

        if intersect_manager:
            print(intersect_manager)
        elif intersect_senior:
            print(intersect_senior)
        elif intersect_mid and not intersect_entry:
            print(intersect_mid)
        elif intersect_entry and not intersect_mid:
            # check if it has the word new grad
            # only want "software engineer" titles
            # print(intersect_entry)
            if not bool(job_title.intersection(set(["new", "grad", "Junior", "Entry"]))):
                intersect_mid = set(["1"])
        else:
            if not bool(job_title.intersection(set(["new", "grad", "Junior", "Entry"]))):
                intersect_mid = set(["1"])
            intersect_entry = set(["1"])
        is_manager = int(bool(intersect_manager))
        is_senior = int(bool(intersect_senior))
        is_mid = int(bool(intersect_mid))
        is_entry = int(bool(intersect_entry))
        is_intern = int(bool(intersect_intern))

        return [is_intern, is_entry, is_mid, is_senior, is_manager]
