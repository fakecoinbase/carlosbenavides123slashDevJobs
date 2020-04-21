import requests
from bs4 import BeautifulSoup
import re
import json
from datetime import date
import time

page = requests.get("https://api.greenhouse.io/v1/boards/twilio/offices")
json_page = page.json()
names = set(["San Francisco", "Redwood City", "New York City", "Mountain View", "Irvine", "Denver"])
departments = set(["Engineering", "Students"])
for obj in json_page["offices"]:
    location = obj["name"]
    if location in names:
        for department in obj["departments"]:
            if department["name"] in departments:
                for job in department["jobs"]:
                    reduce_date = [ int(x) for x in job["updated_at"].split("T")[0].split("-") ]
                    company_listing_date = date(*reduce_date)
                    today = date.today()
                    delta = (today - company_listing_date).days

                    if delta > 30:
                        continue
                    job_location = job["location"]["name"]
                    print(job_location)




                    print(delta)
        # print(obj["departments"])
        print("#####################################33")
        break