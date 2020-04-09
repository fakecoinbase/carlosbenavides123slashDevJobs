class Utils:
    def __init__(self):
        pass
    
    def convert_active_jobs_to_dict(self, active_jobs):
        if not active_jobs:
            return {}
        hmap = {}
        for job_id, company_uuid, active in active_jobs:
            if job_id in hmap:
                print("LOG", job_id, " in hmap??")
            else:
                hmap[job_id] = [job_id, company_uuid, active]
        return hmap
