import requests
from bs4 import BeautifulSoup
import re
import json
from datetime import date
import time

wanted_locations = set(["San Francisco", "New York City", "Venice, CA", "Bellevue, WA", "Denver, CO", "Los Angeles, CA"])
from protos.create_job import create_job

from Utils.LocationUtils import location_builder

def lever(company_uuid, company_name, company_website_scrape, query, utils, kafka):
    print(company_website_scrape)
    page = requests.get(company_website_scrape)
    soup = BeautifulSoup(page.text, 'html.parser')
    res = soup.find_all("div", class_="posting")

    active_jobs = query.get_active_remembered_jobs(company_uuid)
    check_job_list = utils.convert_active_jobs_to_dict(active_jobs)

    for item in res:
        # print(item.findChildren("a", recursive=False))
        for child in item.findChildren("a", recursive=False):
            title = ""
            job_location = ""
            job_link = child['href']

            h5 = child.findChildren("h5")
            if h5:
                title = h5[0].text

            for idx, baby in enumerate(child.findChildren("span")):
                if idx == 0:
                    job_location = baby.text
                if idx == 1 and not title:
                    title = baby.text
            if title and job_location and job_link:
                job_uuid = company_uuid + "_%_" + title.replace(" ", "%" ) + "_%_" +  job_location.replace(" ", "%")
                is_active = query.check_active_job(job_uuid, company_uuid)
                if len(is_active) == 0:
                    location_builder(company_name, job_location, query, kafka)
                    active = 1
                    experience_level = utils.determine_experience_level(title) 
                    data = [job_uuid, company_uuid, job_link, company_website_scrape, "", company_name, experience_level, active, 0, job_location]
                    job = create_job(data)
                    query.insert_new_job( job )
                    query.insert_new_remembered_job( job )
                    
                    kafka.send_protobuf_message("new_job", job)
                else:
                    if job_uuid in check_job_list:
                        del check_job_list[job_uuid]
                    else:
                        print("ERROR", company_name, job_uuid)

    if check_job_list:
        for key, value in check_job_list.items():
            query.deactivate_job(value[0], value[2])
