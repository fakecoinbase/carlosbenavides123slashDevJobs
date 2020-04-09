import requests
from bs4 import BeautifulSoup
import re
import json
from datetime import date

def Honey(UUID, Name, Website, query, utils):
    page = requests.get(Website)
    json_page = page.json()
    _reduce = json_page["departments"]

    # Levels
    Manager = set(["manager"])
    Senior = set(["principle", "senior"])
    Mid = set(["mid"])
    Entry = set(["entry", "junior", "new", "grad", "associate"])
    Intern = set(["internship", "intern"])
    wanted = ["Browser Extension", "Core", "Discovery", "Engineering", "Frontends", "Internship"]

    active_jobs = query.get_active_remembered_jobs(UUID)
    check_job_list = utils.convert_active_jobs_to_dict(active_jobs)

    for item in _reduce:
        if item["name"] in wanted:
            for job in item["jobs"]:
                print(job)
                reduce_date = [ int(x) for x in job["updated_at"].split("T")[0].split("-") ]
                company_listing_date = date(*reduce_date)
                today = date.today()
                delta = (today - company_listing_date).days
                if delta > 30:
                    print("forget about it")
                    break
                location = job["location"]["name"].replace(" ", "%")
                title = job["title"].replace(" ","%")

                job_id = UUID + "_%_" + title + "_%_" + location

                # test deactivating a row
                # first get all relational jobs where isActive == 1
                # then make a list [0] * length of 

                isActive = query.check_active_job(job_id, UUID)

                print(isActive)
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
                    provided_id = job["id"]
                    is_manager = int(bool(intersect_manager))
                    is_senior = int(bool(intersect_senior))
                    is_mid = int(bool(intersect_mid))
                    is_entry = int(bool(intersect_entry))
                    is_intern = int(bool(intersect_intern))
                    active = 1
                    
                    Joblink = job["absolute_url"]
                    query.insert_new_job( (job_id, UUID, Joblink, Website, provided_id, is_intern, is_entry, is_mid, is_senior, is_manager, active) )
                    query.insert_new_remembered_job((job_id, UUID, 1))
                else:
                    break
