from protos.location_proto.create_location_pb import create_location_pb


def location_builder(company_name, company_location, query, kafka):
    locations = ["new york", "san francisco", "los angeles", "venice", "bellevue", "seattle", "san mateo", "denver", "boulder"]
    for location in locations:
        if location in company_location.lower():
            if not query.check_location_company(location.title(), company_name):
                return
            query.insert_new_location(location.title(), company_name)
            data = [location.title(), company_name]
            location_pb = create_location_pb(data)
            kafka.send_protobuf_message("job_location", location_pb)
            break


