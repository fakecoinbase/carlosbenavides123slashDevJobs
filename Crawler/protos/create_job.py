import protos.job_pb2

def create_job(data):
    job = protos.job_pb2.Job()
    job.JobID = data[0]
    job.Company_UUID = data[1]
    job.JobLink = data[2]
    job.DefaultLink = data[3]
    job.ProvidedID = data[4]
    job.CompanyName = data[5]
    job.Internship = data[6]
    job.Entry = data[7]
    job.Mid = data[8]
    job.Senior = data[9]
    job.Manager = data[10]
    job.Active = data[11]
    job.JobPosted = data[12]
    return job