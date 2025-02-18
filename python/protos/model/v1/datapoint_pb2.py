# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: protos/model/v1/datapoint.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    0,
    '',
    'protos/model/v1/datapoint.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from protos.model.v1 import file_pb2 as protos_dot_model_dot_v1_dot_file__pb2
from protos.model.v1 import health_pb2 as protos_dot_model_dot_v1_dot_health__pb2
from protos.model.v1 import math_pb2 as protos_dot_model_dot_v1_dot_math__pb2
from protos.model.v1 import navigation_pb2 as protos_dot_model_dot_v1_dot_navigation__pb2
from protos.model.v1 import text_pb2 as protos_dot_model_dot_v1_dot_text__pb2
from protos.model.v1 import media_pb2 as protos_dot_model_dot_v1_dot_media__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1fprotos/model/v1/datapoint.proto\x12\x08v1.model\x1a\x1aprotos/model/v1/file.proto\x1a\x1cprotos/model/v1/health.proto\x1a\x1aprotos/model/v1/math.proto\x1a protos/model/v1/navigation.proto\x1a\x1aprotos/model/v1/text.proto\x1a\x1bprotos/model/v1/media.proto\"\xea\x05\n\tDatapoint\x12\x0e\n\x06stream\x18\x01 \x01(\t\x12\x11\n\ttimestamp\x18\x02 \x01(\x03\x12+\n\x04tags\x18\x03 \x03(\x0b\x32\x1d.v1.model.Datapoint.TagsEntry\x12\x1e\n\x04text\x18\x04 \x01(\x0b\x32\x0e.v1.model.TextH\x00\x12$\n\x07numeric\x18\x05 \x01(\x0b\x32\x11.v1.model.NumericH\x00\x12+\n\x0bnumeric_set\x18\x11 \x01(\x0b\x32\x14.v1.model.NumericSetH\x00\x12\"\n\x06\x62itset\x18\x07 \x01(\x0b\x32\x10.v1.model.BitsetH\x00\x12\x1e\n\x04\x66ile\x18\x08 \x01(\x0b\x32\x0e.v1.model.FileH\x00\x12 \n\x05image\x18\t \x01(\x0b\x32\x0f.v1.model.ImageH\x00\x12+\n\x0bpoint_cloud\x18\n \x01(\x0b\x32\x14.v1.model.PointCloudH\x00\x12&\n\x08location\x18\x0b \x01(\x0b\x32\x12.v1.model.LocationH\x00\x12.\n\x0clocalization\x18\x0c \x01(\x0b\x32\x16.v1.model.LocalizationH\x00\x12\"\n\x06health\x18\r \x01(\x0b\x32\x10.v1.model.HealthH\x00\x12\x1e\n\x04json\x18\x0e \x01(\x0b\x32\x0e.v1.model.JsonH\x00\x12$\n\x07\x62\x61ttery\x18\x0f \x01(\x0b\x32\x11.v1.model.BatteryH\x00\x12 \n\x05video\x18\x10 \x01(\x0b\x32\x0f.v1.model.VideoH\x00\x12\x31\n\x0etransform_tree\x18\x12 \x01(\x0b\x32\x17.v1.model.TransformTreeH\x00\x12&\n\x08odometry\x18\x13 \x01(\x0b\x32\x12.v1.model.OdometryH\x00\x12\r\n\x05label\x18\x14 \x01(\t\x1a+\n\tTagsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42\x06\n\x04\x64\x61taJ\x04\x08\x06\x10\x07\"\xea\x02\n\x10\x43ontrolDatapoint\x12\x0e\n\x06stream\x18\x01 \x01(\t\x12\x11\n\ttimestamp\x18\x02 \x01(\x03\x12\"\n\x06\x62itset\x18\x03 \x01(\x0b\x32\x10.v1.model.BitsetH\x00\x12 \n\x05twist\x18\x04 \x01(\x0b\x32\x0f.v1.model.TwistH\x00\x12#\n\x04pose\x18\x05 \x01(\x0b\x32\x13.v1.model.TransformH\x00\x12$\n\x07numeric\x18\x06 \x01(\x0b\x32\x11.v1.model.NumericH\x00\x12<\n\x14pose_with_covariance\x18\x07 \x01(\x0b\x32\x1c.v1.model.PoseWithCovarianceH\x00\x12 \n\x05point\x18\x08 \x01(\x0b\x32\x0f.v1.model.PointH\x00\x12\x1c\n\x03joy\x18\t \x01(\x0b\x32\r.v1.model.JoyH\x00\x12\x1c\n\x03\x62it\x18\n \x01(\x0b\x32\r.v1.model.BitH\x00\x42\x06\n\x04\x64\x61ta\"\x81\x02\n\x13GenericAPIDatapoint\x12\x0e\n\x06Method\x18\x01 \x01(\t\x12\x10\n\x08\x45ndpoint\x18\x02 \x01(\t\x12;\n\x07Headers\x18\x03 \x03(\x0b\x32*.v1.model.GenericAPIDatapoint.HeadersEntry\x12\x0c\n\x04\x42ody\x18\x04 \x01(\t\x12\x13\n\x0bIsRetryable\x18\x05 \x01(\x08\x12\x1a\n\x12RequireFormantAuth\x18\x06 \x01(\x08\x12\x1c\n\x14RetryableStatusCodes\x18\x07 \x03(\x03\x1a.\n\x0cHeadersEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42+Z)github.com/FormantIO/genproto/go/v1/modelb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'protos.model.v1.datapoint_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z)github.com/FormantIO/genproto/go/v1/model'
  _globals['_DATAPOINT_TAGSENTRY']._loaded_options = None
  _globals['_DATAPOINT_TAGSENTRY']._serialized_options = b'8\001'
  _globals['_GENERICAPIDATAPOINT_HEADERSENTRY']._loaded_options = None
  _globals['_GENERICAPIDATAPOINT_HEADERSENTRY']._serialized_options = b'8\001'
  _globals['_DATAPOINT']._serialized_start=223
  _globals['_DATAPOINT']._serialized_end=969
  _globals['_DATAPOINT_TAGSENTRY']._serialized_start=912
  _globals['_DATAPOINT_TAGSENTRY']._serialized_end=955
  _globals['_CONTROLDATAPOINT']._serialized_start=972
  _globals['_CONTROLDATAPOINT']._serialized_end=1334
  _globals['_GENERICAPIDATAPOINT']._serialized_start=1337
  _globals['_GENERICAPIDATAPOINT']._serialized_end=1594
  _globals['_GENERICAPIDATAPOINT_HEADERSENTRY']._serialized_start=1548
  _globals['_GENERICAPIDATAPOINT_HEADERSENTRY']._serialized_end=1594
# @@protoc_insertion_point(module_scope)
