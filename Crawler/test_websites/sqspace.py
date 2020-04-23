import requests
from bs4 import BeautifulSoup
import json

page = requests.get("https://api.greenhouse.io/v1/boards/squarespace/departments/")
json_page = page.json()
wanted_locations = set(["San Francisco", "New York, NY"])
wanted_departments = set(["University Grads", "Engineering University Grads", "Engineering Interns", "Engineering", "Infrastructure", "Internal Engineering", "Presence", "Product Platform"])

# def SquareSpace()

for obj in json_page["departments"]:
    if obj["name"] in wanted_departments:
        for job in obj["jobs"]:
            if job["location"]["name"] in wanted_locations:
                print(job)