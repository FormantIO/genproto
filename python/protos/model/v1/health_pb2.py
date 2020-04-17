# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/model/v1/health.proto

from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='protos/model/v1/health.proto',
  package='v1.model',
  syntax='proto3',
  serialized_options=b'Z)github.com/FormantIO/genproto/go/v1/model',
  serialized_pb=b'\n\x1cprotos/model/v1/health.proto\x12\x08v1.model\"8\n\x06Health\x12.\n\x06status\x18\x01 \x01(\x0e\x32\x16.v1.model.HealthStatusR\x06status*`\n\x0cHealthStatus\x12\x12\n\x0eHEALTH_UNKNOWN\x10\x00\x12\x16\n\x12HEALTH_OPERATIONAL\x10\x01\x12\x12\n\x0eHEALTH_OFFLINE\x10\x02\x12\x10\n\x0cHEALTH_ERROR\x10\x03\x42+Z)github.com/FormantIO/genproto/go/v1/modelb\x06proto3'
)

_HEALTHSTATUS = _descriptor.EnumDescriptor(
  name='HealthStatus',
  full_name='v1.model.HealthStatus',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='HEALTH_UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HEALTH_OPERATIONAL', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HEALTH_OFFLINE', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HEALTH_ERROR', index=3, number=3,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=100,
  serialized_end=196,
)
_sym_db.RegisterEnumDescriptor(_HEALTHSTATUS)

HealthStatus = enum_type_wrapper.EnumTypeWrapper(_HEALTHSTATUS)
HEALTH_UNKNOWN = 0
HEALTH_OPERATIONAL = 1
HEALTH_OFFLINE = 2
HEALTH_ERROR = 3



_HEALTH = _descriptor.Descriptor(
  name='Health',
  full_name='v1.model.Health',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='v1.model.Health.status', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='status', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=42,
  serialized_end=98,
)

_HEALTH.fields_by_name['status'].enum_type = _HEALTHSTATUS
DESCRIPTOR.message_types_by_name['Health'] = _HEALTH
DESCRIPTOR.enum_types_by_name['HealthStatus'] = _HEALTHSTATUS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Health = _reflection.GeneratedProtocolMessageType('Health', (_message.Message,), {
  'DESCRIPTOR' : _HEALTH,
  '__module__' : 'protos.model.v1.health_pb2'
  # @@protoc_insertion_point(class_scope:v1.model.Health)
  })
_sym_db.RegisterMessage(Health)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
