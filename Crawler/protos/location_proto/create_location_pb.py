import protos.location_proto.location_pb2

def create_location_pb(data):
    location_pb = protos.location_proto.location_pb2.Location()
    location_pb.Location = data[0]
    location_pb.CompanyName = data[1]
    return location_pb