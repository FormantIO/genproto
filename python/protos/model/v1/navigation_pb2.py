# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: protos/model/v1/navigation.proto
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
    'protos/model/v1/navigation.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from protos.model.v1 import math_pb2 as protos_dot_model_dot_v1_dot_math__pb2
from protos.model.v1 import media_pb2 as protos_dot_model_dot_v1_dot_media__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n protos/model/v1/navigation.proto\x12\x08v1.model\x1a\x1aprotos/model/v1/math.proto\x1a\x1bprotos/model/v1/media.proto\"}\n\x08Location\x12\x10\n\x08latitude\x18\x01 \x01(\x01\x12\x11\n\tlongitude\x18\x02 \x01(\x01\x12\x15\n\x08\x61ltitude\x18\x03 \x01(\x01H\x00\x88\x01\x01\x12\x18\n\x0borientation\x18\x04 \x01(\x01H\x01\x88\x01\x01\x42\x0b\n\t_altitudeB\x0e\n\x0c_orientation\"\xb8\x01\n\x0cLocalization\x12$\n\x08odometry\x18\x01 \x01(\x0b\x32\x12.v1.model.Odometry\x12\x1a\n\x03map\x18\x02 \x01(\x0b\x32\r.v1.model.Map\x12*\n\x0cpoint_clouds\x18\x03 \x03(\x0b\x32\x14.v1.model.PointCloud\x12\x1c\n\x04path\x18\x04 \x01(\x0b\x32\x0e.v1.model.Path\x12\x1c\n\x04goal\x18\x05 \x01(\x0b\x32\x0e.v1.model.Goal\"z\n\x08Odometry\x12!\n\x04pose\x18\x01 \x01(\x0b\x32\x13.v1.model.Transform\x12\x1e\n\x05twist\x18\x02 \x01(\x0b\x32\x0f.v1.model.Twist\x12+\n\x0eworld_to_local\x18\x03 \x01(\x0b\x32\x13.v1.model.Transform\"\xf1\x01\n\x03Map\x12\x0c\n\x04uuid\x18\t \x01(\t\x12\x12\n\nresolution\x18\x01 \x01(\x01\x12\r\n\x05width\x18\x02 \x01(\r\x12\x0e\n\x06height\x18\x03 \x01(\r\x12#\n\x06origin\x18\x04 \x01(\x0b\x32\x13.v1.model.Transform\x12+\n\x0eworld_to_local\x18\x05 \x01(\x0b\x32\x13.v1.model.Transform\x12\x31\n\x0eoccupancy_grid\x18\x06 \x01(\x0b\x32\x17.v1.model.OccupancyGridH\x00\x12\r\n\x03url\x18\x07 \x01(\tH\x00\x12\r\n\x03raw\x18\x08 \x01(\x0cH\x00\x42\x06\n\x04\x64\x61ta\"\x1d\n\rOccupancyGrid\x12\x0c\n\x04\x64\x61ta\x18\x01 \x03(\x05\"W\n\x04Path\x12+\n\x0eworld_to_local\x18\x01 \x01(\x0b\x32\x13.v1.model.Transform\x12\"\n\x05poses\x18\x02 \x03(\x0b\x32\x13.v1.model.Transform\"{\n\nJointState\x12+\n\x0eworld_to_local\x18\x05 \x01(\x0b\x32\x13.v1.model.Transform\x12\x0c\n\x04name\x18\x01 \x03(\t\x12\x10\n\x08position\x18\x02 \x03(\x01\x12\x10\n\x08velocity\x18\x03 \x03(\x01\x12\x0e\n\x06\x65\x66\x66ort\x18\x04 \x03(\x01\"V\n\x04Goal\x12+\n\x0eworld_to_local\x18\x01 \x01(\x0b\x32\x13.v1.model.Transform\x12!\n\x04pose\x18\x02 \x01(\x0b\x32\x13.v1.model.Transform\"\x14\n\x06GoalID\x12\n\n\x02id\x18\x01 \x01(\t\"K\n\x12PoseWithCovariance\x12!\n\x04pose\x18\x01 \x01(\x0b\x32\x13.v1.model.Transform\x12\x12\n\ncovariance\x18\x02 \x03(\x01\"7\n\tColorRGBA\x12\t\n\x01r\x18\x01 \x01(\x02\x12\t\n\x01g\x18\x02 \x01(\x02\x12\t\n\x01\x62\x18\x03 \x01(\x02\x12\t\n\x01\x61\x18\x04 \x01(\x02\"\xa2\x03\n\x08Marker3D\x12+\n\x0eworld_to_local\x18\x01 \x01(\x0b\x32\x13.v1.model.Transform\x12\n\n\x02ns\x18\x02 \x01(\t\x12\n\n\x02id\x18\x03 \x01(\x05\x12\x0c\n\x04type\x18\x04 \x01(\t\x12\x0e\n\x06\x61\x63tion\x18\x05 \x01(\t\x12!\n\x04pose\x18\x06 \x01(\x0b\x32\x13.v1.model.Transform\x12 \n\x05scale\x18\x07 \x01(\x0b\x32\x11.v1.model.Vector3\x12\"\n\x05\x63olor\x18\x08 \x01(\x0b\x32\x13.v1.model.ColorRGBA\x12\x10\n\x08lifetime\x18\t \x01(\x01\x12\x14\n\x0c\x66rame_locked\x18\n \x01(\x08\x12!\n\x06points\x18\x0b \x03(\x0b\x32\x11.v1.model.Vector3\x12#\n\x06\x63olors\x18\x0c \x03(\x0b\x32\x13.v1.model.ColorRGBA\x12\x0c\n\x04text\x18\r \x01(\t\x12\x15\n\rmesh_resource\x18\x0e \x01(\t\x12#\n\x1bmesh_use_embedded_materials\x18\x0f \x01(\x08\x12\x10\n\x08\x66rame_id\x18\x10 \x01(\t\"4\n\rMarker3DArray\x12#\n\x07markers\x18\x01 \x03(\x0b\x32\x12.v1.model.Marker3D\"$\n\x03Joy\x12\x0c\n\x04\x61xes\x18\x01 \x03(\x02\x12\x0f\n\x07\x62uttons\x18\x02 \x03(\x05\x42+Z)github.com/FormantIO/genproto/go/v1/modelb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'protos.model.v1.navigation_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z)github.com/FormantIO/genproto/go/v1/model'
  _globals['_LOCATION']._serialized_start=103
  _globals['_LOCATION']._serialized_end=228
  _globals['_LOCALIZATION']._serialized_start=231
  _globals['_LOCALIZATION']._serialized_end=415
  _globals['_ODOMETRY']._serialized_start=417
  _globals['_ODOMETRY']._serialized_end=539
  _globals['_MAP']._serialized_start=542
  _globals['_MAP']._serialized_end=783
  _globals['_OCCUPANCYGRID']._serialized_start=785
  _globals['_OCCUPANCYGRID']._serialized_end=814
  _globals['_PATH']._serialized_start=816
  _globals['_PATH']._serialized_end=903
  _globals['_JOINTSTATE']._serialized_start=905
  _globals['_JOINTSTATE']._serialized_end=1028
  _globals['_GOAL']._serialized_start=1030
  _globals['_GOAL']._serialized_end=1116
  _globals['_GOALID']._serialized_start=1118
  _globals['_GOALID']._serialized_end=1138
  _globals['_POSEWITHCOVARIANCE']._serialized_start=1140
  _globals['_POSEWITHCOVARIANCE']._serialized_end=1215
  _globals['_COLORRGBA']._serialized_start=1217
  _globals['_COLORRGBA']._serialized_end=1272
  _globals['_MARKER3D']._serialized_start=1275
  _globals['_MARKER3D']._serialized_end=1693
  _globals['_MARKER3DARRAY']._serialized_start=1695
  _globals['_MARKER3DARRAY']._serialized_end=1747
  _globals['_JOY']._serialized_start=1749
  _globals['_JOY']._serialized_end=1785
# @@protoc_insertion_point(module_scope)
