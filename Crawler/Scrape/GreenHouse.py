import requests
from bs4 import BeautifulSoup
import json
from Utils.GreenHouseUtils import extract_job_details

wanted_locations = set(["San Francisco", "New York, NY", "Redwood City", "Remote", "New York City", "Palo Alto, CA", "Los Angeles, CA", "Boulder, CO", "Mountain View", "Irvine", "Denver", "Los Angeles", "Downtown Los Angeles", "Boulder, Colorado", "San Francisco", "New York, NY", "San Francisco", "Redwood City", "New York City", "Mountain View", "Irvine", "Denver", "Denver, CO", "Remote", "San Francisco, CA"])
wanted_departments = set(["Enterprise Applications", "Engineering - Frontend", "Shopping Features", "Engineering", "Engineering - Backend", "Internships & University Grad Positions", "Engineering - Data", "Engineering - Infrastructure", "Interns & Early Career", "University", "Web Development", "Product Engineering",  "Internships", "Browser Extension", "Core", "Discovery", "Frontends","Students", "University Grads", "Engineering University Grads", "Engineering Interns", "Engineering", "Infrastructure", "Internal Engineering", "Presence", "Product Platform", "Engineering", "Students", "Engineering - Backend"])

# todo
def greenhouse(company_uuid, company_name, company_website_scrape, query, utils, kafka):
    page = requests.get(company_website_scrape)
    json_page = page.json()
    reduce_departments = json_page["departments"]

    active_jobs = query.get_active_remembered_jobs(company_uuid)
    check_job_list = utils.convert_active_jobs_to_dict(active_jobs)

    company_departments, company_locations = query.get_company_scrape_details(company_uuid)[0]
    if company_departments == None:
        company_departments = ""
    if company_locations == None:
        company_locations = ""
    company_departments_set = set(company_departments.split(","))
    company_locations_set = set(company_locations.split("/"))
    print(company_departments_set, company_locations_set)

    for department in reduce_departments:
        if department["name"] in company_departments_set or department["name"] in wanted_departments:
            for job in department["jobs"]:
                if job["location"]["name"] in company_locations_set or job["location"]["name"] in wanted_locations:
                    extract_job_details(job, company_uuid, company_website_scrape, company_name, check_job_list, query, utils, kafka)

    # for department in reduce_departments:
    #     if 

    if check_job_list:
        for key, value in check_job_list.items():
            query.deactivate_job(value[0], value[2])