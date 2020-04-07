import requests
from bs4 import BeautifulSoup
import re
import json
URL = 'https://api.greenhouse.io/v1/boards/honey/departments'
page = requests.get(URL)
json_page = page.json()
_reduce = json_page["departments"]
wanted = ["Core", "Discovery", "Engineering", "Frontends", "Internship"]
for item in _reduce:
    if item["name"] in wanted:
        print(item["name"])
        print(item["jobs"])
    else:
        print("no")


# soup = BeautifulSoup(page.content, 'html.parser')
# result = soup.findAll(
# results_title = soup.find_all("div")[0].prettify()
# results_location = soup.find_all(_class='title3')

