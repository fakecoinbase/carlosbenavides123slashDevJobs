import requests
from bs4 import BeautifulSoup
import re
import json
from datetime import date
import time

page = requests.get("https://api.greenhouse.io/v1/boards/applovin/departments")
json_page = page.json()
wanted_locations = set(["Palo Alto, CA", "New York City"])
wanted_departments = set(["Engineering ", "Internships"])

for item in json_page["departments"]:
    if item["name"] in wanted_departments:
        for job in item["jobs"]:
            if job["location"]["name"] in wanted_locations:
                print(job)

