import requests
from bs4 import BeautifulSoup
import json
from Utils.GreenHouse import extract_job_details

# todo
def GreenHouse(company_uuid, company_name, company_website_scrape, query, utils, kafka, wanted_departments, wanted_locations):
    page = requests.get(company_website_scrape)
    json_page = page.json()
    reduce_departments = json_page["departments"]

    active_jobs = query.get_active_remembered_jobs(company_uuid)
    check_job_list = utils.convert_active_jobs_to_dict(active_jobs)

    for department in reduce_departments:
        if department["name"] in wanted_departments:
            for job in department["jobs"]:
                if job["location"]["name"] in wanted_locations:
                    extract_job_details(job, company_uuid, company_website_scrape, company_name, check_job_list, query, utils, kafka)
    if check_job_list:
        for key, value in check_job_list.items():
            query.deactivate_job(value[0], value[2])