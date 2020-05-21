import requests
from bs4 import BeautifulSoup
import re
import json
from datetime import date
date_format = "%m/%d/%Y"

title = "honey"
website = "https://api.greenhouse.io/v1/boards/honey/departments"
location = "Los Angeles"

page = requests.get(website)
json_page = page.json()
_reduce = json_page["departments"]
wanted = ["Browser Extension", "Core", "Discovery", "Engineering", "Frontends", "Internship"]
unwanted_title = ["Senior", "Principle", "Manager"]

for item in _reduce:
    print(item["name"])
    # for job in item["jobs"]:
    #     print(job["location"]["name"])
    if item["name"] in wanted:
        for job in item["jobs"]:
            # print(job["location"]["name"])
            if location in job["location"]["name"]:
                print("ye")
            else:
                print(job["location"])
            reduce_date = [ int(x) for x in job["updated_at"].split("T")[0].split("-") ]
            company_listing_date = date(*reduce_date)
            today = date.today()
            delta = (today - company_listing_date).days
            if delta > 30:
                print("forget about it")
                break
            print(delta)