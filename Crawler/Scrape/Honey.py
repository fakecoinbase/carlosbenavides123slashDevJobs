import requests
from bs4 import BeautifulSoup
import json
from Utils.GreenHouse import extract_job_details

WANTED = ["Browser Extension", "Core", "Discovery", "Engineering", "Frontends", "Internship"]
def Honey(company_uuid, company_name, company_website_scrape, query, utils, kafka):
    page = requests.get(company_website_scrape)
    json_page = page.json()
    _reduce = json_page["departments"]

    active_jobs = query.get_active_remembered_jobs(company_uuid)
    check_job_list = utils.convert_active_jobs_to_dict(active_jobs)

    # print(json_page)
    for item in _reduce:
        if item["name"] in WANTED:
            for job in item["jobs"]:
                extract_job_details(job, company_uuid, company_website_scrape, company_name, check_job_list, query, utils, kafka)


    if check_job_list:
        print(check_job_list)
        for key, value in check_job_list.items():
            query.deactivate_job(value[0], value[1], value[2])
