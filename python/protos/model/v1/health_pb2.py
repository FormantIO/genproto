# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/model/v1/health.proto
"""Generated protocol buffer code."""
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
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x1cprotos/model/v1/health.proto\x12\x08v1.model\"L\n\x06Health\x12&\n\x06status\x18\x01 \x01(\x0e\x32\x16.v1.model.HealthStatus\x12\x11\n\x07skew_ms\x18\x02 \x01(\x03H\x00\x42\x07\n\x05\x63lock\"O\n\x07\x42\x61ttery\x12\x12\n\npercentage\x18\x01 \x01(\x01\x12\x0f\n\x07voltage\x18\x02 \x01(\x01\x12\x0f\n\x07\x63urrent\x18\x03 \x01(\x01\x12\x0e\n\x06\x63harge\x18\x04 \x01(\x01\"\xf4\x01\n\x0e\x42ufferMetadata\x12\x12\n\nsession_id\x18\x01 \x01(\t\x12\x30\n(latest_ttl_stream_channel_buffered_bytes\x18\x02 \x01(\x04\x12.\n&reliable_stream_channel_buffered_bytes\x18\x03 \x01(\x04\x12\x35\n-latest_reliable_stream_channel_buffered_bytes\x18\x04 \x01(\x04\x12\x35\n-latest_try_once_stream_channel_buffered_bytes\x18\x05 \x01(\x04*`\n\x0cHealthStatus\x12\x12\n\x0eHEALTH_UNKNOWN\x10\x00\x12\x16\n\x12HEALTH_OPERATIONAL\x10\x01\x12\x12\n\x0eHEALTH_OFFLINE\x10\x02\x12\x10\n\x0cHEALTH_ERROR\x10\x03\x42+Z)github.com/FormantIO/genproto/go/v1/modelb\x06proto3'
)

_HEALTHSTATUS = _descriptor.EnumDescriptor(
  name='HealthStatus',
  full_name='v1.model.HealthStatus',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='HEALTH_UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='HEALTH_OPERATIONAL', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='HEALTH_OFFLINE', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='HEALTH_ERROR', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=448,
  serialized_end=544,
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
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='v1.model.Health.status', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='skew_ms', full_name='v1.model.Health.skew_ms', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
    _descriptor.OneofDescriptor(
      name='clock', full_name='v1.model.Health.clock',
      index=0, containing_type=None,
      create_key=_descriptor._internal_create_key,
    fields=[]),
  ],
  serialized_start=42,
  serialized_end=118,
)


_BATTERY = _descriptor.Descriptor(
  name='Battery',
  full_name='v1.model.Battery',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='percentage', full_name='v1.model.Battery.percentage', index=0,
      number=1, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='voltage', full_name='v1.model.Battery.voltage', index=1,
      number=2, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='current', full_name='v1.model.Battery.current', index=2,
      number=3, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='charge', full_name='v1.model.Battery.charge', index=3,
      number=4, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=120,
  serialized_end=199,
)


_BUFFERMETADATA = _descriptor.Descriptor(
  name='BufferMetadata',
  full_name='v1.model.BufferMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='session_id', full_name='v1.model.BufferMetadata.session_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='latest_ttl_stream_channel_buffered_bytes', full_name='v1.model.BufferMetadata.latest_ttl_stream_channel_buffered_bytes', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='reliable_stream_channel_buffered_bytes', full_name='v1.model.BufferMetadata.reliable_stream_channel_buffered_bytes', index=2,
      number=3, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='latest_reliable_stream_channel_buffered_bytes', full_name='v1.model.BufferMetadata.latest_reliable_stream_channel_buffered_bytes', index=3,
      number=4, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='latest_try_once_stream_channel_buffered_bytes', full_name='v1.model.BufferMetadata.latest_try_once_stream_channel_buffered_bytes', index=4,
      number=5, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=202,
  serialized_end=446,
)

_HEALTH.fields_by_name['status'].enum_type = _HEALTHSTATUS
_HEALTH.oneofs_by_name['clock'].fields.append(
  _HEALTH.fields_by_name['skew_ms'])
_HEALTH.fields_by_name['skew_ms'].containing_oneof = _HEALTH.oneofs_by_name['clock']
DESCRIPTOR.message_types_by_name['Health'] = _HEALTH
DESCRIPTOR.message_types_by_name['Battery'] = _BATTERY
DESCRIPTOR.message_types_by_name['BufferMetadata'] = _BUFFERMETADATA
DESCRIPTOR.enum_types_by_name['HealthStatus'] = _HEALTHSTATUS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Health = _reflection.GeneratedProtocolMessageType('Health', (_message.Message,), {
  'DESCRIPTOR' : _HEALTH,
  '__module__' : 'protos.model.v1.health_pb2'
  # @@protoc_insertion_point(class_scope:v1.model.Health)
  })
_sym_db.RegisterMessage(Health)

Battery = _reflection.GeneratedProtocolMessageType('Battery', (_message.Message,), {
  'DESCRIPTOR' : _BATTERY,
  '__module__' : 'protos.model.v1.health_pb2'
  # @@protoc_insertion_point(class_scope:v1.model.Battery)
  })
_sym_db.RegisterMessage(Battery)

BufferMetadata = _reflection.GeneratedProtocolMessageType('BufferMetadata', (_message.Message,), {
  'DESCRIPTOR' : _BUFFERMETADATA,
  '__module__' : 'protos.model.v1.health_pb2'
  # @@protoc_insertion_point(class_scope:v1.model.BufferMetadata)
  })
_sym_db.RegisterMessage(BufferMetadata)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
