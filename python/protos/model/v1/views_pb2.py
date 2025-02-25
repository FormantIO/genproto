# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/model/v1/views.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from protos.model.v1 import datapoint_pb2 as protos_dot_model_dot_v1_dot_datapoint__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='protos/model/v1/views.proto',
  package='v1.model',
  syntax='proto3',
  serialized_options=b'Z)github.com/FormantIO/genproto/go/v1/model',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x1bprotos/model/v1/views.proto\x12\x08v1.model\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1fprotos/model/v1/datapoint.proto\x1a\x1cgoogle/protobuf/struct.proto\"6\n\rViewsMetadata\x12%\n\x05views\x18\x01 \x03(\x0b\x32\x16.v1.model.ViewMetadata\"\xb5\x04\n\x0cViewMetadata\x12\n\n\x02id\x18\x01 \x01(\t\x12.\n\x04tags\x18\x02 \x03(\x0b\x32 .v1.model.ViewMetadata.TagsEntry\x12\x16\n\x0eorganizationId\x18\x03 \x01(\t\x12\x0c\n\x04name\x18\x04 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x05 \x01(\t\x12\x0b\n\x03url\x18\x06 \x01(\t\x12\x1a\n\x12showOnSingleDevice\x18\x07 \x01(\x08\x12\x19\n\x11showOnMultiDevice\x18\x08 \x01(\x08\x12\x14\n\x0cshowOnTeleop\x18\t \x01(\x08\x12\x14\n\x0cshowTimeline\x18\x0b \x01(\x08\x12 \n\x06\x66ilter\x18\x0c \x01(\x0b\x32\x10.v1.model.Filter\x12\'\n\x06layout\x18\r \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\nlayoutType\x18\x0e \x01(\t\x12\x32\n\rconfiguration\x18\x0f \x03(\x0b\x32\x1b.v1.model.ViewConfiguration\x12\x14\n\x0csmartFleetId\x18\x10 \x01(\t\x12\r\n\x05index\x18\x11 \x01(\x03\x1a+\n\tTagsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"S\n\nLayoutType\x12\r\n\tDASHBOARD\x10\x00\x12\x0b\n\x07OBSERVE\x10\x01\x12\x0e\n\nFULLSCREEN\x10\x02\x12\r\n\tANALYTICS\x10\x03\x12\n\n\x06TELEOP\x10\x04J\x04\x08\n\x10\x0b\"\x94\x02\n\x06\x46ilter\x12(\n\x04tags\x18\x01 \x03(\x0b\x32\x1a.v1.model.Filter.TagsEntry\x12\r\n\x05names\x18\x02 \x03(\t\x12\r\n\x05types\x18\x03 \x03(\t\x12\x11\n\tdeviceIds\x18\x04 \x03(\t\x12.\n\x07notTags\x18\x05 \x03(\x0b\x32\x1d.v1.model.Filter.NotTagsEntry\x12\x10\n\x08notNames\x18\x06 \x03(\t\x12\x10\n\x08\x61gentIds\x18\x07 \x03(\t\x1a+\n\tTagsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x1a.\n\x0cNotTagsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"e\n\x11ViewConfiguration\x12\x12\n\nstreamName\x18\x01 \x01(\t\x12\x0c\n\x04type\x18\x02 \x01(\t\x12.\n\rconfiguration\x18\x03 \x01(\x0b\x32\x17.google.protobuf.StructB+Z)github.com/FormantIO/genproto/go/v1/modelb\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,protos_dot_model_dot_v1_dot_datapoint__pb2.DESCRIPTOR,google_dot_protobuf_dot_struct__pb2.DESCRIPTOR,])



_VIEWMETADATA_LAYOUTTYPE = _descriptor.EnumDescriptor(
  name='LayoutType',
  full_name='v1.model.ViewMetadata.LayoutType',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='DASHBOARD', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='OBSERVE', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='FULLSCREEN', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='ANALYTICS', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='TELEOP', index=4, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=670,
  serialized_end=753,
)
_sym_db.RegisterEnumDescriptor(_VIEWMETADATA_LAYOUTTYPE)


_VIEWSMETADATA = _descriptor.Descriptor(
  name='ViewsMetadata',
  full_name='v1.model.ViewsMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='views', full_name='v1.model.ViewsMetadata.views', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=137,
  serialized_end=191,
)


_VIEWMETADATA_TAGSENTRY = _descriptor.Descriptor(
  name='TagsEntry',
  full_name='v1.model.ViewMetadata.TagsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='v1.model.ViewMetadata.TagsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='v1.model.ViewMetadata.TagsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=b'8\001',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=625,
  serialized_end=668,
)

_VIEWMETADATA = _descriptor.Descriptor(
  name='ViewMetadata',
  full_name='v1.model.ViewMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='v1.model.ViewMetadata.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='tags', full_name='v1.model.ViewMetadata.tags', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='organizationId', full_name='v1.model.ViewMetadata.organizationId', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='v1.model.ViewMetadata.name', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='description', full_name='v1.model.ViewMetadata.description', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='url', full_name='v1.model.ViewMetadata.url', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='showOnSingleDevice', full_name='v1.model.ViewMetadata.showOnSingleDevice', index=6,
      number=7, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='showOnMultiDevice', full_name='v1.model.ViewMetadata.showOnMultiDevice', index=7,
      number=8, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='showOnTeleop', full_name='v1.model.ViewMetadata.showOnTeleop', index=8,
      number=9, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='showTimeline', full_name='v1.model.ViewMetadata.showTimeline', index=9,
      number=11, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='filter', full_name='v1.model.ViewMetadata.filter', index=10,
      number=12, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='layout', full_name='v1.model.ViewMetadata.layout', index=11,
      number=13, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='layoutType', full_name='v1.model.ViewMetadata.layoutType', index=12,
      number=14, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='configuration', full_name='v1.model.ViewMetadata.configuration', index=13,
      number=15, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='smartFleetId', full_name='v1.model.ViewMetadata.smartFleetId', index=14,
      number=16, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='index', full_name='v1.model.ViewMetadata.index', index=15,
      number=17, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_VIEWMETADATA_TAGSENTRY, ],
  enum_types=[
    _VIEWMETADATA_LAYOUTTYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=194,
  serialized_end=759,
)


_FILTER_TAGSENTRY = _descriptor.Descriptor(
  name='TagsEntry',
  full_name='v1.model.Filter.TagsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='v1.model.Filter.TagsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='v1.model.Filter.TagsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=b'8\001',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=625,
  serialized_end=668,
)

_FILTER_NOTTAGSENTRY = _descriptor.Descriptor(
  name='NotTagsEntry',
  full_name='v1.model.Filter.NotTagsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='v1.model.Filter.NotTagsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='v1.model.Filter.NotTagsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=b'8\001',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=992,
  serialized_end=1038,
)

_FILTER = _descriptor.Descriptor(
  name='Filter',
  full_name='v1.model.Filter',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='tags', full_name='v1.model.Filter.tags', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='names', full_name='v1.model.Filter.names', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='types', full_name='v1.model.Filter.types', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='deviceIds', full_name='v1.model.Filter.deviceIds', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='notTags', full_name='v1.model.Filter.notTags', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='notNames', full_name='v1.model.Filter.notNames', index=5,
      number=6, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='agentIds', full_name='v1.model.Filter.agentIds', index=6,
      number=7, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_FILTER_TAGSENTRY, _FILTER_NOTTAGSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=762,
  serialized_end=1038,
)


_VIEWCONFIGURATION = _descriptor.Descriptor(
  name='ViewConfiguration',
  full_name='v1.model.ViewConfiguration',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='streamName', full_name='v1.model.ViewConfiguration.streamName', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='type', full_name='v1.model.ViewConfiguration.type', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='configuration', full_name='v1.model.ViewConfiguration.configuration', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=1040,
  serialized_end=1141,
)

_VIEWSMETADATA.fields_by_name['views'].message_type = _VIEWMETADATA
_VIEWMETADATA_TAGSENTRY.containing_type = _VIEWMETADATA
_VIEWMETADATA.fields_by_name['tags'].message_type = _VIEWMETADATA_TAGSENTRY
_VIEWMETADATA.fields_by_name['filter'].message_type = _FILTER
_VIEWMETADATA.fields_by_name['layout'].message_type = google_dot_protobuf_dot_struct__pb2._STRUCT
_VIEWMETADATA.fields_by_name['configuration'].message_type = _VIEWCONFIGURATION
_VIEWMETADATA_LAYOUTTYPE.containing_type = _VIEWMETADATA
_FILTER_TAGSENTRY.containing_type = _FILTER
_FILTER_NOTTAGSENTRY.containing_type = _FILTER
_FILTER.fields_by_name['tags'].message_type = _FILTER_TAGSENTRY
_FILTER.fields_by_name['notTags'].message_type = _FILTER_NOTTAGSENTRY
_VIEWCONFIGURATION.fields_by_name['configuration'].message_type = google_dot_protobuf_dot_struct__pb2._STRUCT
DESCRIPTOR.message_types_by_name['ViewsMetadata'] = _VIEWSMETADATA
DESCRIPTOR.message_types_by_name['ViewMetadata'] = _VIEWMETADATA
DESCRIPTOR.message_types_by_name['Filter'] = _FILTER
DESCRIPTOR.message_types_by_name['ViewConfiguration'] = _VIEWCONFIGURATION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ViewsMetadata = _reflection.GeneratedProtocolMessageType('ViewsMetadata', (_message.Message,), {
  'DESCRIPTOR' : _VIEWSMETADATA,
  '__module__' : 'protos.model.v1.views_pb2'
  # @@protoc_insertion_point(class_scope:v1.model.ViewsMetadata)
  })
_sym_db.RegisterMessage(ViewsMetadata)

ViewMetadata = _reflection.GeneratedProtocolMessageType('ViewMetadata', (_message.Message,), {

  'TagsEntry' : _reflection.GeneratedProtocolMessageType('TagsEntry', (_message.Message,), {
    'DESCRIPTOR' : _VIEWMETADATA_TAGSENTRY,
    '__module__' : 'protos.model.v1.views_pb2'
    # @@protoc_insertion_point(class_scope:v1.model.ViewMetadata.TagsEntry)
    })
  ,
  'DESCRIPTOR' : _VIEWMETADATA,
  '__module__' : 'protos.model.v1.views_pb2'
  # @@protoc_insertion_point(class_scope:v1.model.ViewMetadata)
  })
_sym_db.RegisterMessage(ViewMetadata)
_sym_db.RegisterMessage(ViewMetadata.TagsEntry)

Filter = _reflection.GeneratedProtocolMessageType('Filter', (_message.Message,), {

  'TagsEntry' : _reflection.GeneratedProtocolMessageType('TagsEntry', (_message.Message,), {
    'DESCRIPTOR' : _FILTER_TAGSENTRY,
    '__module__' : 'protos.model.v1.views_pb2'
    # @@protoc_insertion_point(class_scope:v1.model.Filter.TagsEntry)
    })
  ,

  'NotTagsEntry' : _reflection.GeneratedProtocolMessageType('NotTagsEntry', (_message.Message,), {
    'DESCRIPTOR' : _FILTER_NOTTAGSENTRY,
    '__module__' : 'protos.model.v1.views_pb2'
    # @@protoc_insertion_point(class_scope:v1.model.Filter.NotTagsEntry)
    })
  ,
  'DESCRIPTOR' : _FILTER,
  '__module__' : 'protos.model.v1.views_pb2'
  # @@protoc_insertion_point(class_scope:v1.model.Filter)
  })
_sym_db.RegisterMessage(Filter)
_sym_db.RegisterMessage(Filter.TagsEntry)
_sym_db.RegisterMessage(Filter.NotTagsEntry)

ViewConfiguration = _reflection.GeneratedProtocolMessageType('ViewConfiguration', (_message.Message,), {
  'DESCRIPTOR' : _VIEWCONFIGURATION,
  '__module__' : 'protos.model.v1.views_pb2'
  # @@protoc_insertion_point(class_scope:v1.model.ViewConfiguration)
  })
_sym_db.RegisterMessage(ViewConfiguration)


DESCRIPTOR._options = None
_VIEWMETADATA_TAGSENTRY._options = None
_FILTER_TAGSENTRY._options = None
_FILTER_NOTTAGSENTRY._options = None
# @@protoc_insertion_point(module_scope)
