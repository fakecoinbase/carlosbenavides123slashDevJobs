import requests
from bs4 import BeautifulSoup
import json

page = requests.get("https://api.greenhouse.io/v1/boards/coinbase/departments")
json_page = page.json()

wanted_locations = set(["San Francisco", "New York, NY", "Redwood City", "Remote"])
wanted_departments = set(["Engineering", "Engineering - Backend", "Internships & University Grad Positions", "Engineering - Data", "Engineering - Infrastructure"])

for obj in json_page["departments"]:
    if obj["name"] in wanted_departments:
        for job in obj["jobs"]:
            if job["location"]["name"] in wanted_locations:
                if obj["name"] == "Internships & University Grad Positions":
                    if "Legal" in job["title"]:
                        break
                print(job)
                print("#############################")