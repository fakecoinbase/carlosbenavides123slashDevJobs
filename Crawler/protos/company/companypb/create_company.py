import protos.company.companypb.company_pb2 as company_pb2

def create_company_pb(company_uuid, company_name, company_website):
    company_pb = company_pb2.Company()
    company_pb.CompanyUUID = company_uuid
    company_pb.CompanyName = company_name
    company_pb.CompanyWebsite = company_website
    return company_pb