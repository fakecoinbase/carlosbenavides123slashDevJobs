import requests
from bs4 import BeautifulSoup
import re
import json
from datetime import date
import time

# page = requests.get("https://jobs.lever.co/chownow?team=Engineering")
page = requests.get("https://jobs.lever.co/system1/?team=Engineering")
# json_page = page.content()
soup = BeautifulSoup(page.text, 'html')
#wanted_locations = set(["San Francisco", "New York City"])
#wanted_departments = set(["Interns & Early Career", "University", "Web Development", "Product Engineering"])
# print(soup.prettify())
res = soup.find_all("div", class_="posting")
for item in res:
    # print(item.findChildren("a", recursive=False))
    for child in item.findChildren("a", recursive=False):
        job = []
        job.append(child['href'])
        for baby in child.findChildren("span"):
            job.append(baby.text)
        print(job)
    # print(item)
    print("###")
# print(soup.find("div", class_="posting"))
