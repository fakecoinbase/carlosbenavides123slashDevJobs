from datetime import date
import time
from protos.create_job import create_job

from Utils.LocationUtils import location_builder

def extract_job_details(job, company_uuid, company_website_scrape, company_name, check_job_list, query, utils, kafka):
    reduce_date = [ int(x) for x in job["updated_at"].split("T")[0].split("-") ]
    company_listing_date = date(*reduce_date)
    today = date.today()
    delta = (today - company_listing_date).days
    # if the job has been posted for over 30 days, move onto to the next job...
    if delta > 45:
        return

    job_uuid = company_uuid + "_%_" + job["title"].replace(" ", "%" ) + "_%_" +  job["location"]["name"].replace(" ", "%")
    is_active = query.check_active_job(job_uuid, company_uuid)

    # new job
    if len(is_active) == 0:
        experience_level = utils.determine_experience_level(job["title"])                    
        provided_id = str(job["id"])
        job_link = job["absolute_url"]
        job_location = job["location"]["name"]

        location_builder(company_name, job_location, query, kafka)

        time_posted = int(time.mktime(company_listing_date.timetuple()))
        active = 1

        data = [job_uuid, company_uuid, job_link, company_website_scrape, provided_id, company_name, experience_level, active, time_posted, job_location]
        job = create_job(data)
        query.insert_new_job( job )
        query.insert_new_remembered_job( job )
        
        kafka.send_protobuf_message("new_job", job)
    else:
        if job_uuid in check_job_list:
            del check_job_list[job_uuid]
        else:
            print("ERROR:", company_name, job_uuid)


