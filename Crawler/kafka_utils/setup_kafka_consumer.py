from time import sleep
import threading
import sys
from kafka import KafkaConsumer

import protos.company.companypb.company_pb2 as company_pb2
import protos.company.companyrequestpb.companyrequestpb_pb2 as companyrequest_pb2
from protos.company.companypb.create_company import create_company_pb
from protos.company.company_cms_pb.create_company_cms_pb import create_company_cms_pb

class KafkaConsumerScheduler:
	def __init__(self, producer, query):
		self.consumer = KafkaConsumer(bootstrap_servers='192.168.1.66:19092')
		self.consumer.subscribe(['RequestCmsHome', 'AddNewCompany', "RequestCMSCompany"])
		self.t = threading.Thread(target=self.read_consumer_messages)
		self.producer = producer
		self.query = query

	def start(self):
		try:
			self.t.start()
		except KeyboardInterrupt:
			self.t.join(0)
			sys.exit()

	def read_consumer_messages(self):
		for msg in self.consumer:
			print(msg)
			topic = msg.topic
			# meta = self.consumer.partitions_for_topic(topic)
			# options = {}
			# options[partition] = OffsetAndMetadata(message.offset + 1, meta)
			# self.consumer.commit(options)

			if msg.topic == "RequestCmsHome":
				self.response_cms_home()
			elif msg.topic == "AddNewCompany":
				self.add_new_company(msg)
			elif msg.topic == "RequestCMSCompany":
				self.response_cms_company(msg)

	def add_new_company(self, msg):
		new_company = company_pb2.Company()
		new_company.ParseFromString(msg.value)
		print(new_company)
		gh, lvr, oth = self.determine_company_website(new_company.CompanyWebsite)
		self.query.add_new_company(new_company, gh, lvr, oth)
		self.query.add_new_scrape_details(new_company.CompanyUUID, new_company.CompanyName)

	def response_cms_home(self):
		print("here")
		companies = self.query.get_companies_from_scrappy()
		company_response_pb = company_pb2.CompanyResponse()
		if companies:
			for company_uuid, company_name, company_website in companies:
				company_pb = create_company_pb(company_uuid, company_name, company_website)
				company_response_pb.companies.extend([company_pb])
			print("ResponseCmsHome", company_response_pb)
			self.producer.send_protobuf_message("ResponseCmsHome", company_response_pb)
		else:
			print("LOG ERROR: companies DNE")
	
	def response_cms_company(self, msg):
		company_request = companyrequest_pb2.CompanyRequest()
		company_request.ParseFromString(msg.value)
		print(company_request)
		res = self.query.get_cms_company_details(company_request.CompanyName)
		print(res)
		if res:
			data = [res[1], res[2], res[8], res[9], res[3], res[4], res[5]]
			if data[2] == None:
				data[2] = ""
			if data[3] == None:
				data[3] = ""
			company_cms = create_company_cms_pb(data)
			self.producer.send_protobuf_message("ResponseCompanyCMS", company_cms)

	def determine_company_website(self, company_website):
		gh = 0
		lvr = 0
		oth = 0
		if "greenhouse" in company_website:
			gh = 1
		elif "lever" in company_website:
			lvr = 1
		else:
			oth = 1
		return gh, lvr, oth

