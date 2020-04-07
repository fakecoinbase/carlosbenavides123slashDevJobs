import requests
from bs4 import BeautifulSoup
import re
import json

title = "criteriacorp"
website = "https://www.criteriacorp.com/company/careers.php"
location = "Los Angeles"

page = requests.get(website).content
soup = BeautifulSoup(page, features='html.parser')

mydivs = soup.findAll("a")

for node in mydivs:
    print(node)

# print(mydivs)
