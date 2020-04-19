import requests
from bs4 import BeautifulSoup
import re
import json
from datetime import date
import time
from protos.create_job import create_job

WANTED = ["Browser Extension", "Core", "Discovery", "Engineering", "Frontends", "Internship"]
def Honey(uuid, company_name, company_website_scrape, query, utils, kafka):
    page = requests.get(company_website_scrape)
    json_page = page.json()
    _reduce = json_page["departments"]

    active_jobs = query.get_active_remembered_jobs(uuid)
    check_job_list = utils.convert_active_jobs_to_dict(active_jobs)

    # print(json_page)
    for item in _reduce:
        if item["name"] in WANTED:
            for job in item["jobs"]:
                reduce_date = [ int(x) for x in job["updated_at"].split("T")[0].split("-") ]
                company_listing_date = date(*reduce_date)
                today = date.today()
                delta = (today - company_listing_date).days

                if delta > 30:
                    print("forget about it")
                    break
                job_location = job["location"]["name"]
                location = job["location"]["name"].replace(" ", "%")
                title = job["title"].replace(" ","%")
                job_id = uuid + "_%_" + title + "_%_" + location
                is_active = query.check_active_job(job_id, uuid)

                if len(is_active) == 0:
                    # send to Jobs Table
                    experience_level = utils.determine_experience_level(job["title"])                    
                    provided_id = str(job["id"])
                    job_link = job["absolute_url"]
                    time_posted = int(time.mktime(company_listing_date.timetuple()))
                    active = 1

                    data = [job_id, uuid, job_link, company_website_scrape, provided_id, company_name, experience_level, active, time_posted, job_location]
                    job = create_job(data)
                    query.insert_new_job( job )
                    query.insert_new_remembered_job( job )
                    
                    kafka.send_protobuf_message("new_job", job)
                else:
                    del check_job_list[job_id]
    # delete the stragglers in here
    # these are the jobs not found in the website
    # so these jobs are probably donezo
    if check_job_list:
        print(check_job_list)
        for key, value in check_job_list.items():
            query.deactivate_job(value[0], value[1], value[2])
