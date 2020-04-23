from datetime import date
import time
from protos.create_job import create_job

def extract_job_details_lever(job, company_uuid, company_website_scrape, company_name, check_job_list, query, utils, kafka):
    page = requests.get(company_website_scrape)
    soup = BeautifulSoup(page.text, 'html')
    res = soup.find_all("div", class_="posting")
    for item in res:
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
            
        # print(item)