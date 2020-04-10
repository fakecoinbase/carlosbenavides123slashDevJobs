import protos.job_pb2

def create_job(data):
    job = protos.job_pb2.Job()
    job.JobID = data[0]
    job.Company_UUID = data[1]
    job.JobLink = data[2]
    job.DefaultLink = data[3]
    job.ProvidedID = data[4]
    job.Internship = data[5]
    job.Entry = data[6]
    job.Mid = data[7]
    job.Senior = data[8]
    job.Manager = data[9]
    job.Active = data[10]
    return job