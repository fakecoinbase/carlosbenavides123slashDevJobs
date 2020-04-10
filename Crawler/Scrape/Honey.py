import requests
from bs4 import BeautifulSoup
import re
import json
from datetime import date
import time
from protos.create_job import create_job

WANTED = ["Browser Extension", "Core", "Discovery", "Engineering", "Frontends", "Internship"]
def Honey(UUID, Name, Website, query, utils, kafka):
    page = requests.get(Website)
    json_page = page.json()
    _reduce = json_page["departments"]

    active_jobs = query.get_active_remembered_jobs(UUID)
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
                location = job["location"]["name"].replace(" ", "%")
                title = job["title"].replace(" ","%")
                job_id = UUID + "_%_" + title + "_%_" + location

                isActive = query.check_active_job(job_id, UUID)
                if len(isActive) == 0:
                    # send to Jobs Table

                    # check the level of the job
                    job_title = set(job["title"].lower().split(" "))
                    experience_level = utils.determine_experience_level(job_title)
                    
                    provided_id = str(job["id"])
                    Joblink = job["absolute_url"]

                    data = (job_id, UUID, Joblink, Website, provided_id, Name) + experience_level

                    query.insert_new_job( data )
                    query.insert_new_remembered_job( (data[0], data[1], data[4], data[-1]) )
                    
                    temp_time_arr = [int(time.mktime(company_listing_date.timetuple()))]
                    job = create_job(tuple(list(data) + temp_time_arr ))
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
