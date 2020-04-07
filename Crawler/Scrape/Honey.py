import requests
from bs4 import BeautifulSoup
import re
import json
from datetime import date

def Honey(UUID, Name, Website, cursor):
    page = requests.get(Website)
    json_page = page.json()
    _reduce = json_page["departments"]

    # Levels
    Manager = set(["manager"])
    Senior = set(["principle", "senior"])
    Mid = set(["mid"])
    Entry = set(["entry", "junior", "new", "grad"])
    Intern = set(["internship", "intern"])
    wanted = ["Browser Extension", "Core", "Discovery", "Engineering", "Frontends", "Internship"]

    for item in _reduce:
        if item["name"] in wanted:
            for job in item["jobs"]:
                reduce_date = [ int(x) for x in job["updated_at"].split("T")[0].split("-") ]
                company_listing_date = date(*reduce_date)
                today = date.today()
                delta = (today - company_listing_date).days
                if delta > 30:
                    print("forget about it")
                    break
                location = job["location"]["name"].replace(" ", "%")
                title = job["title"].replace(" ","%")
                JobID = UUID + "_%_" + title + "_%_" + location
                print(JobID)
                SQL_CHECK_REMEMBERED = "SELECT isActive FROM Remembered_Jobs WHERE JobID=%s AND Company_UUID=%s"
                cursor.execute(SQL_CHECK_REMEMBERED, (JobID, UUID))
                isActive = cursor.fetchall()
                if len(isActive) == 0:
                    # send to Jobs Table

                    # check the level of the job
                    title_key_word = set(job["title"].lower().split(" "))
                    print(title_key_word)
                    intersect_manager = title_key_word.intersection(Manager)
                    intersect_senior = title_key_word.intersection(Senior)
                    intersect_mid = title_key_word.intersection(Mid)
                    intersect_entry = title_key_word.intersection(Entry)
                    intersect_intern = title_key_word.intersection(Intern)
                    # results = tuple( intersect_senior, intersect_mid, intersect_entry,   )
                    if intersect_manager:
                        print(intersect_manager)
                    elif intersect_senior:
                        print(intersect_senior)
                    elif intersect_mid and not intersect_entry:
                        print(intersect_mid)
                    elif intersect_entry and not intersect_mid:
                        # check if it has the word new grad
                        # only want "software engineer" titles
                        print(intersect_entry)
                        if not bool(title_key_word.intersection(set(["new", "grad", "Junior", "Entry"]))):
                            intersect_mid = set(["1"])
                    else:
                        if not bool(title_key_word.intersection(set(["new", "grad", "Junior", "Entry"]))):
                            intersect_mid = set(["1"])
                        intersect_entry = set(["1"])

                    is_manager = int(bool(intersect_manager))
                    is_senior = int(bool(intersect_senior))
                    is_mid = int(bool(intersect_mid))
                    is_entry = int(bool(intersect_entry))
                    is_intern = int(bool(intersect_intern))

                    Joblink = job["absolute_url"]
                    SQL = "INSERT INTO Jobs (JobID, Company_UUID, Joblink, DefaultLink, Internship, Entry, Mid, Senior, Manager) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s)"
                    print(SQL)
                    print((JobID, UUID, Joblink, Website, is_intern, is_entry, is_mid, is_senior, is_manager))
                    cursor.execute(SQL, (JobID, UUID, Joblink, Website, is_intern, is_entry, is_mid, is_senior, is_manager))
                else:
                    break
