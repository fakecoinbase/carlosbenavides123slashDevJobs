# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: company_cms_details.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='company_cms_details.proto',
  package='',
  syntax='proto3',
  serialized_pb=_b('\n\x19\x63ompany_cms_details.proto\"\xa6\x01\n\x11\x43ompanyCmsDetails\x12\x13\n\x0b\x43ompanyName\x18\x01 \x01(\t\x12\x16\n\x0e\x43ompanyWebsite\x18\x02 \x01(\t\x12\x19\n\x11WantedDepartments\x18\x03 \x01(\t\x12\x17\n\x0fWantedLocations\x18\x04 \x01(\t\x12\x12\n\nGreenHouse\x18\x05 \x01(\x08\x12\r\n\x05Lever\x18\x06 \x01(\x08\x12\r\n\x05Other\x18\x07 \x01(\x08\x42\x13Z\x11\x43ompanyCmsDetailsb\x06proto3')
)




_COMPANYCMSDETAILS = _descriptor.Descriptor(
  name='CompanyCmsDetails',
  full_name='CompanyCmsDetails',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='CompanyName', full_name='CompanyCmsDetails.CompanyName', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='CompanyWebsite', full_name='CompanyCmsDetails.CompanyWebsite', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='WantedDepartments', full_name='CompanyCmsDetails.WantedDepartments', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='WantedLocations', full_name='CompanyCmsDetails.WantedLocations', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='GreenHouse', full_name='CompanyCmsDetails.GreenHouse', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Lever', full_name='CompanyCmsDetails.Lever', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Other', full_name='CompanyCmsDetails.Other', index=6,
      number=7, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=30,
  serialized_end=196,
)

DESCRIPTOR.message_types_by_name['CompanyCmsDetails'] = _COMPANYCMSDETAILS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CompanyCmsDetails = _reflection.GeneratedProtocolMessageType('CompanyCmsDetails', (_message.Message,), dict(
  DESCRIPTOR = _COMPANYCMSDETAILS,
  __module__ = 'company_cms_details_pb2'
  # @@protoc_insertion_point(class_scope:CompanyCmsDetails)
  ))
_sym_db.RegisterMessage(CompanyCmsDetails)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\021CompanyCmsDetails'))
# @@protoc_insertion_point(module_scope)
