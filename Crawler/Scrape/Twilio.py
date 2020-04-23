import requests
from bs4 import BeautifulSoup
import json
from Utils.GreenHouse import extract_job_details

names = set(["San Francisco", "Redwood City", "New York City", "Mountain View", "Irvine", "Denver"])
departments = set(["Engineering", "Students"])

def Twilio(company_uuid, company_name, company_website_scrape, query, utils, kafka):
    page = requests.get(company_website_scrape)
    json_page = page.json()

    active_jobs = query.get_active_remembered_jobs(company_uuid)
    check_job_list = utils.convert_active_jobs_to_dict(active_jobs)


    for obj in json_page["offices"]:
        location = obj["name"]
        if location in names:
            for department in obj["departments"]:
                if department["name"] in departments:
                    for job in department["jobs"]:
                        extract_job_details(job, company_uuid, company_website_scrape, company_name, check_job_list, query, utils, kafka)

    if check_job_list:
        print("LOG:", check_job_list)
        for key, value in check_job_list.items():
            query.deactivate_job(value[0], value[1], value[2])
