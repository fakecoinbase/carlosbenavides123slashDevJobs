import requests
from bs4 import BeautifulSoup
import re
import json
from datetime import date
import time

page = requests.get("https://jobs.lever.co/goat?department=GOAT&team=Technology")
# page = requests.get("https://jobs.lever.co/verkada/?department=Engineering")
# json_page = page.content()
soup = BeautifulSoup(page.text, 'html')
#wanted_locations = set(["San Francisco", "New York City"])
#wanted_departments = set(["Interns & Early Career", "University", "Web Development", "Product Engineering"])
# print(soup.prettify())
res = soup.find_all("div", class_="posting")
for item in res:
    # print(item.findChildren("a", recursive=False))
    for child in item.findChildren("a", recursive=False):
        title = ""
        location = ""
        link = child['href']

        h5 = child.findChildren("h5")
        if h5:
            title = h5[0].text

        for idx, baby in enumerate(child.findChildren("span")):
            if idx == 0:
                location = baby.text
            if idx == 1 and not title:
                title = baby.text
        print(title, location, link)
    # print(item)
    print("###")
# print(soup.find("div", class_="posting"))
