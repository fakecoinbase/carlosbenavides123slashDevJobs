import requests
from bs4 import BeautifulSoup
import re
import json
from datetime import date
import time

page = requests.get("https://api.greenhouse.io/v1/boards/asana/departments")
json_page = page.json()
wanted_locations = set(["San Francisco", "New York City"])
wanted_departments = set(["Interns & Early Career", "University", "Web Development", "Product Engineering"])

for department in json_page["departments"]:
    if department["name"] in wanted_departments:
        for job in department["jobs"]:
            if job["location"]["name"] in wanted_locations:
                print(job)