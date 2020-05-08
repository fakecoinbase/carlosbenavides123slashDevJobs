import protos.company.company_cms_pb.company_cms_details_pb2 as company_cms_details_pb2

def create_company_cms_pb(data):
    company_cms_details = company_cms_details_pb2.CompanyCmsDetails()
    company_cms_details.CompanyName = data[0]
    company_cms_details.CompanyWebsite = data[1]
    company_cms_details.WantedDepartments = data[2]
    company_cms_details.WantedLocations = data[3]
    company_cms_details.GreenHouse = data[4]
    company_cms_details.Lever = data[5]
    company_cms_details.Other = data[6]
    return company_cms_details