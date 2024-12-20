# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: protos/model/v1/media.proto
# Protobuf Python Version: 5.28.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
    1,
    '',
    'protos/model/v1/media.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from protos.model.v1 import math_pb2 as protos_dot_model_dot_v1_dot_math__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bprotos/model/v1/media.proto\x12\x08v1.model\x1a\x1aprotos/model/v1/math.proto\"C\n\x05Image\x12\x14\n\x0c\x63ontent_type\x18\x01 \x01(\t\x12\r\n\x03url\x18\x02 \x01(\tH\x00\x12\r\n\x03raw\x18\x03 \x01(\x0cH\x00\x42\x06\n\x04\x64\x61ta\"m\n\nPointCloud\x12\x0c\n\x04uuid\x18\x04 \x01(\t\x12\r\n\x03url\x18\x01 \x01(\tH\x00\x12\r\n\x03raw\x18\x02 \x01(\x0cH\x00\x12+\n\x0eworld_to_local\x18\x03 \x01(\x0b\x32\x13.v1.model.TransformB\x06\n\x04\x64\x61ta\"J\n\rRtcPointCloud\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\x0c\x12+\n\x0eworld_to_local\x18\x02 \x01(\x0b\x32\x13.v1.model.Transform\"B\n\x0eH264VideoFrame\x12\r\n\x05index\x18\x01 \x01(\x05\x12\r\n\x05\x66lags\x18\x02 \x01(\x05\x12\x12\n\nframe_data\x18\x03 \x01(\x0c\"?\n\nAudioChunk\x12\r\n\x05index\x18\x01 \x01(\x05\x12\x0e\n\x06\x66ormat\x18\x02 \x01(\t\x12\x12\n\nchunk_data\x18\x03 \x01(\x0c\"R\n\x05Video\x12\x11\n\tmime_type\x18\x01 \x01(\t\x12\x10\n\x08\x64uration\x18\x02 \x01(\x03\x12\r\n\x03url\x18\x03 \x01(\tH\x00\x12\r\n\x03raw\x18\x04 \x01(\x0cH\x00\x42\x06\n\x04\x64\x61ta\"5\n\rTransformTree\x12\r\n\x03url\x18\x01 \x01(\tH\x00\x12\r\n\x03raw\x18\x02 \x01(\x0cH\x00\x42\x06\n\x04\x64\x61taB+Z)github.com/FormantIO/genproto/go/v1/modelb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'protos.model.v1.media_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z)github.com/FormantIO/genproto/go/v1/model'
  _globals['_IMAGE']._serialized_start=69
  _globals['_IMAGE']._serialized_end=136
  _globals['_POINTCLOUD']._serialized_start=138
  _globals['_POINTCLOUD']._serialized_end=247
  _globals['_RTCPOINTCLOUD']._serialized_start=249
  _globals['_RTCPOINTCLOUD']._serialized_end=323
  _globals['_H264VIDEOFRAME']._serialized_start=325
  _globals['_H264VIDEOFRAME']._serialized_end=391
  _globals['_AUDIOCHUNK']._serialized_start=393
  _globals['_AUDIOCHUNK']._serialized_end=456
  _globals['_VIDEO']._serialized_start=458
  _globals['_VIDEO']._serialized_end=540
  _globals['_TRANSFORMTREE']._serialized_start=542
  _globals['_TRANSFORMTREE']._serialized_end=595
# @@protoc_insertion_point(module_scope)
