import protos.job_pb2

def create_job(data):
    job = protos.job_pb2.Job()
    job.JobUUID = data[0]
    job.CompanyUUID = data[1]
    job.JobLink = data[2]
    job.DefaultLink = data[3]
    job.ProvidedID = data[4]
    job.CompanyName = data[5]
    job.ExperienceLevel = data[6]
    job.Active = data[7]
    job.JobPosted = data[8]
    job.Location = data[9]
    return job