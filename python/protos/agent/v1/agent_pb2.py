# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: protos/agent/v1/agent.proto
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
    'protos/agent/v1/agent.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from protos.model.v1 import datapoint_pb2 as protos_dot_model_dot_v1_dot_datapoint__pb2
from protos.model.v1 import intervention_pb2 as protos_dot_model_dot_v1_dot_intervention__pb2
from protos.model.v1 import commands_pb2 as protos_dot_model_dot_v1_dot_commands__pb2
from protos.model.v1 import config_pb2 as protos_dot_model_dot_v1_dot_config__pb2
from protos.model.v1 import math_pb2 as protos_dot_model_dot_v1_dot_math__pb2
from protos.model.v1 import event_pb2 as protos_dot_model_dot_v1_dot_event__pb2
from protos.model.v1 import health_pb2 as protos_dot_model_dot_v1_dot_health__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bprotos/agent/v1/agent.proto\x12\x08v1.agent\x1a\x1fprotos/model/v1/datapoint.proto\x1a\"protos/model/v1/intervention.proto\x1a\x1eprotos/model/v1/commands.proto\x1a\x1cprotos/model/v1/config.proto\x1a\x1aprotos/model/v1/math.proto\x1a\x1bprotos/model/v1/event.proto\x1a\x1cprotos/model/v1/health.proto\"S\n\'PostGenericAPIUnbufferedRequestResponse\x12\x12\n\nstatusCode\x18\x01 \x01(\r\x12\x14\n\x0cresponseBody\x18\x02 \x01(\t\"\x1f\n\x1dPostGenericAPIRequestResponse\"\x14\n\x12StreamDataResponse\"\x12\n\x10PostDataResponse\"P\n\rPostDataError\x12\r\n\x05index\x18\x01 \x01(\r\x12\x0c\n\x04\x63ode\x18\x02 \x01(\r\x12\x11\n\tretryable\x18\x03 \x01(\x08\x12\x0f\n\x07message\x18\x04 \x01(\t\"?\n\x14PostDataMultiRequest\x12\'\n\ndatapoints\x18\x01 \x03(\x0b\x32\x13.v1.model.Datapoint\"\x17\n\x15PostDataMultiResponse\"=\n\x12PostDataMultiError\x12\'\n\x06\x65rrors\x18\x01 \x03(\x0b\x32\x17.v1.agent.PostDataError\"+\n\x1dGetInterventionRequestRequest\x12\n\n\x02id\x18\x01 \x01(\t\"4\n\x1eGetInterventionResponseRequest\x12\x12\n\nrequest_id\x18\x01 \x01(\t\" \n\x1eGetStreamsConfigurationRequest\"Q\n\x1fGetStreamsConfigurationResponse\x12.\n\x07streams\x18\x01 \x03(\x0b\x32\x1d.v1.model.StreamConfiguration\"$\n\"GetApplicationConfigurationRequest\"`\n#GetApplicationConfigurationResponse\x12\x39\n\rconfiguration\x18\x01 \x01(\x0b\x32\".v1.model.ApplicationConfiguration\"\x1a\n\x18GetConfigBlobDataRequest\"B\n\x19GetConfigBlobDataResponse\x12%\n\tblob_data\x18\x01 \x01(\x0b\x32\x12.v1.model.BlobData\"\x1e\n\x1cGetAgentConfigurationRequest\"T\n\x1dGetAgentConfigurationResponse\x12\x33\n\rconfiguration\x18\x01 \x01(\x0b\x32\x1c.v1.model.AgentConfiguration\"\x1a\n\x18GetBufferMetadataRequest\"h\n\x16QueryDatapointsRequest\x12\x13\n\x0bstream_name\x18\x01 \x01(\t\x12\r\n\x05start\x18\x02 \x01(\t\x12\x0b\n\x03\x65nd\x18\x03 \x01(\t\x12\r\n\x05limit\x18\x04 \x01(\r\x12\x0e\n\x06offset\x18\x05 \x01(\r\"?\n\x17QueryDatapointsResponse\x12$\n\x07results\x18\x01 \x03(\x0b\x32\x13.v1.model.Datapoint\"O\n\x12QueryEventsRequest\x12\r\n\x05start\x18\x01 \x01(\t\x12\x0b\n\x03\x65nd\x18\x02 \x01(\t\x12\r\n\x05limit\x18\x03 \x01(\r\x12\x0e\n\x06offset\x18\x04 \x01(\r\"7\n\x13QueryEventsResponse\x12 \n\x07results\x18\x01 \x03(\x0b\x32\x0f.v1.model.Event\"J\n\x19GetBufferMetadataResponse\x12-\n\x0b\x63onnections\x18\x01 \x03(\x0b\x32\x18.v1.model.BufferMetadata\"\x0f\n\rHealthRequest\"\x10\n\x0eHealthResponse\"2\n\x18GetCommandRequestRequest\x12\x16\n\x0e\x63ommand_filter\x18\x01 \x03(\t\"F\n\x19GetCommandRequestResponse\x12)\n\x07request\x18\x01 \x01(\x0b\x32\x18.v1.model.CommandRequest\"I\n\x1aSendCommandResponseRequest\x12+\n\x08response\x18\x01 \x01(\x0b\x32\x19.v1.model.CommandResponse\"\x1d\n\x1bSendCommandResponseResponse\"8\n\x1eGetCommandRequestStreamRequest\x12\x16\n\x0e\x63ommand_filter\x18\x01 \x03(\t\"L\n\x1fGetCommandRequestStreamResponse\x12)\n\x07request\x18\x01 \x01(\x0b\x32\x18.v1.model.CommandRequest\":\n!GetTeleopControlDataStreamRequest\x12\x15\n\rstream_filter\x18\x01 \x03(\t\"[\n\"GetTeleopControlDataStreamResponse\x12\x35\n\x11\x63ontrol_datapoint\x18\x01 \x01(\x0b\x32\x1a.v1.model.ControlDatapoint\"!\n\x1fGetTeleopHeartbeatStreamRequest\"`\n GetTeleopHeartbeatStreamResponse\x12\x0f\n\x07peer_id\x18\x01 \x01(\t\x12\x15\n\ris_disconnect\x18\x02 \x01(\x08\x12\x14\n\x0csession_type\x18\x03 \x01(\t\":\n!GetTelemetryListenerStreamRequest\x12\x15\n\rstream_filter\x18\x01 \x03(\t\"L\n\"GetTelemetryListenerStreamResponse\x12&\n\tdatapoint\x18\x01 \x01(\x0b\x32\x13.v1.model.Datapoint\"\x1c\n\x1aPostTransformFrameResponse\"#\n\x15SetBaseFrameIDRequest\x12\n\n\x02id\x18\x01 \x01(\t\"\x18\n\x16SetBaseFrameIDResponse\"\x1b\n\x19\x43learTransformTreeRequest\"\x1c\n\x1a\x43learTransformTreeResponse\"4\n\x12\x43reateEventRequest\x12\x1e\n\x05\x65vent\x18\x01 \x01(\x0b\x32\x0f.v1.model.Event\"\x15\n\x13\x43reateEventResponse\"%\n\x10\x43reateEventError\x12\x11\n\tretryable\x18\x01 \x01(\x08\"\x16\n\x14GetTeleopInfoRequest\"1\n\x15GetTeleopInfoResponse\x12\x18\n\x10\x63onnection_count\x18\x01 \x01(\r\"\'\n\x16PostLanRtcOfferRequest\x12\r\n\x05offer\x18\x01 \x01(\t\")\n\x17PostLanRtcOfferResponse\x12\x0e\n\x06\x61nswer\x18\x01 \x01(\t\"G\n\x1eSendOnCustomDataChannelRequest\x12\x14\n\x0c\x63hannel_name\x18\x01 \x01(\t\x12\x0f\n\x07payload\x18\x02 \x01(\x0c\"!\n\x1fSendOnCustomDataChannelResponse\"G\n(GetCustomDataChannelMessageStreamRequest\x12\x1b\n\x13\x63hannel_name_filter\x18\x01 \x03(\t\"c\n)GetCustomDataChannelMessageStreamResponse\x12\x0f\n\x07peer_id\x18\x01 \x01(\t\x12\x14\n\x0c\x63hannel_name\x18\x02 \x01(\t\x12\x0f\n\x07payload\x18\x03 \x01(\x0c\x32\xa2\x17\n\x05\x41gent\x12\x43\n\nStreamData\x12\x13.v1.model.Datapoint\x1a\x1c.v1.agent.StreamDataResponse\"\x00(\x01\x12=\n\x08PostData\x12\x13.v1.model.Datapoint\x1a\x1a.v1.agent.PostDataResponse\"\x00\x12R\n\rPostDataMulti\x12\x1e.v1.agent.PostDataMultiRequest\x1a\x1f.v1.agent.PostDataMultiResponse\"\x00\x12{\n\x1aGetTeleopControlDataStream\x12+.v1.agent.GetTeleopControlDataStreamRequest\x1a,.v1.agent.GetTeleopControlDataStreamResponse\"\x00\x30\x01\x12u\n\x18GetTeleopHeartbeatStream\x12).v1.agent.GetTeleopHeartbeatStreamRequest\x1a*.v1.agent.GetTeleopHeartbeatStreamResponse\"\x00\x30\x01\x12{\n\x1aGetTelemetryListenerStream\x12+.v1.agent.GetTelemetryListenerStreamRequest\x1a,.v1.agent.GetTelemetryListenerStreamResponse\"\x00\x30\x01\x12\x90\x01\n!GetCustomDataChannelMessageStream\x12\x32.v1.agent.GetCustomDataChannelMessageStreamRequest\x1a\x33.v1.agent.GetCustomDataChannelMessageStreamResponse\"\x00\x30\x01\x12L\n\x0b\x43reateEvent\x12\x1c.v1.agent.CreateEventRequest\x1a\x1d.v1.agent.CreateEventResponse\"\x00\x12[\n\x19\x43reateInterventionRequest\x12\x1d.v1.model.InterventionRequest\x1a\x1d.v1.model.InterventionRequest\"\x00\x12\x62\n\x16GetInterventionRequest\x12\'.v1.agent.GetInterventionRequestRequest\x1a\x1d.v1.model.InterventionRequest\"\x00\x12\x65\n\x17GetInterventionResponse\x12(.v1.agent.GetInterventionResponseRequest\x1a\x1e.v1.model.InterventionResponse\"\x00\x12p\n\x17GetStreamsConfiguration\x12(.v1.agent.GetStreamsConfigurationRequest\x1a).v1.agent.GetStreamsConfigurationResponse\"\x00\x12|\n\x1bGetApplicationConfiguration\x12,.v1.agent.GetApplicationConfigurationRequest\x1a-.v1.agent.GetApplicationConfigurationResponse\"\x00\x12^\n\x11GetConfigBlobData\x12\".v1.agent.GetConfigBlobDataRequest\x1a#.v1.agent.GetConfigBlobDataResponse\"\x00\x12j\n\x15GetAgentConfiguration\x12&.v1.agent.GetAgentConfigurationRequest\x1a\'.v1.agent.GetAgentConfigurationResponse\"\x00\x12^\n\x11GetBufferMetadata\x12\".v1.agent.GetBufferMetadataRequest\x1a#.v1.agent.GetBufferMetadataResponse\"\x00\x12X\n\x0fQueryDatapoints\x12 .v1.agent.QueryDatapointsRequest\x1a!.v1.agent.QueryDatapointsResponse\"\x00\x12L\n\x0bQueryEvents\x12\x1c.v1.agent.QueryEventsRequest\x1a\x1d.v1.agent.QueryEventsResponse\"\x00\x12=\n\x06Health\x12\x17.v1.agent.HealthRequest\x1a\x18.v1.agent.HealthResponse\"\x00\x12^\n\x11GetCommandRequest\x12\".v1.agent.GetCommandRequestRequest\x1a#.v1.agent.GetCommandRequestResponse\"\x00\x12r\n\x17GetCommandRequestStream\x12(.v1.agent.GetCommandRequestStreamRequest\x1a).v1.agent.GetCommandRequestStreamResponse\"\x00\x30\x01\x12\x64\n\x13SendCommandResponse\x12$.v1.agent.SendCommandResponseRequest\x1a%.v1.agent.SendCommandResponseResponse\"\x00\x12V\n\x12PostTransformFrame\x12\x18.v1.model.TransformFrame\x1a$.v1.agent.PostTransformFrameResponse\"\x00\x12U\n\x0eSetBaseFrameID\x12\x1f.v1.agent.SetBaseFrameIDRequest\x1a .v1.agent.SetBaseFrameIDResponse\"\x00\x12\x61\n\x12\x43learTransformTree\x12#.v1.agent.ClearTransformTreeRequest\x1a$.v1.agent.ClearTransformTreeResponse\"\x00\x12R\n\rGetTeleopInfo\x12\x1e.v1.agent.GetTeleopInfoRequest\x1a\x1f.v1.agent.GetTeleopInfoResponse\"\x00\x12X\n\x0fPostLanRtcOffer\x12 .v1.agent.PostLanRtcOfferRequest\x1a!.v1.agent.PostLanRtcOfferResponse\"\x00\x12p\n\x17SendOnCustomDataChannel\x12(.v1.agent.SendOnCustomDataChannelRequest\x1a).v1.agent.SendOnCustomDataChannelResponse\"\x00\x12\x61\n\x15PostGenericAPIRequest\x12\x1d.v1.model.GenericAPIDatapoint\x1a\'.v1.agent.PostGenericAPIRequestResponse\"\x00\x12u\n\x1fPostGenericAPIUnbufferedRequest\x12\x1d.v1.model.GenericAPIDatapoint\x1a\x31.v1.agent.PostGenericAPIUnbufferedRequestResponse\"\x00\x42+Z)github.com/FormantIO/genproto/go/v1/agentb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'protos.agent.v1.agent_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z)github.com/FormantIO/genproto/go/v1/agent'
  _globals['_POSTGENERICAPIUNBUFFEREDREQUESTRESPONSE']._serialized_start=259
  _globals['_POSTGENERICAPIUNBUFFEREDREQUESTRESPONSE']._serialized_end=342
  _globals['_POSTGENERICAPIREQUESTRESPONSE']._serialized_start=344
  _globals['_POSTGENERICAPIREQUESTRESPONSE']._serialized_end=375
  _globals['_STREAMDATARESPONSE']._serialized_start=377
  _globals['_STREAMDATARESPONSE']._serialized_end=397
  _globals['_POSTDATARESPONSE']._serialized_start=399
  _globals['_POSTDATARESPONSE']._serialized_end=417
  _globals['_POSTDATAERROR']._serialized_start=419
  _globals['_POSTDATAERROR']._serialized_end=499
  _globals['_POSTDATAMULTIREQUEST']._serialized_start=501
  _globals['_POSTDATAMULTIREQUEST']._serialized_end=564
  _globals['_POSTDATAMULTIRESPONSE']._serialized_start=566
  _globals['_POSTDATAMULTIRESPONSE']._serialized_end=589
  _globals['_POSTDATAMULTIERROR']._serialized_start=591
  _globals['_POSTDATAMULTIERROR']._serialized_end=652
  _globals['_GETINTERVENTIONREQUESTREQUEST']._serialized_start=654
  _globals['_GETINTERVENTIONREQUESTREQUEST']._serialized_end=697
  _globals['_GETINTERVENTIONRESPONSEREQUEST']._serialized_start=699
  _globals['_GETINTERVENTIONRESPONSEREQUEST']._serialized_end=751
  _globals['_GETSTREAMSCONFIGURATIONREQUEST']._serialized_start=753
  _globals['_GETSTREAMSCONFIGURATIONREQUEST']._serialized_end=785
  _globals['_GETSTREAMSCONFIGURATIONRESPONSE']._serialized_start=787
  _globals['_GETSTREAMSCONFIGURATIONRESPONSE']._serialized_end=868
  _globals['_GETAPPLICATIONCONFIGURATIONREQUEST']._serialized_start=870
  _globals['_GETAPPLICATIONCONFIGURATIONREQUEST']._serialized_end=906
  _globals['_GETAPPLICATIONCONFIGURATIONRESPONSE']._serialized_start=908
  _globals['_GETAPPLICATIONCONFIGURATIONRESPONSE']._serialized_end=1004
  _globals['_GETCONFIGBLOBDATAREQUEST']._serialized_start=1006
  _globals['_GETCONFIGBLOBDATAREQUEST']._serialized_end=1032
  _globals['_GETCONFIGBLOBDATARESPONSE']._serialized_start=1034
  _globals['_GETCONFIGBLOBDATARESPONSE']._serialized_end=1100
  _globals['_GETAGENTCONFIGURATIONREQUEST']._serialized_start=1102
  _globals['_GETAGENTCONFIGURATIONREQUEST']._serialized_end=1132
  _globals['_GETAGENTCONFIGURATIONRESPONSE']._serialized_start=1134
  _globals['_GETAGENTCONFIGURATIONRESPONSE']._serialized_end=1218
  _globals['_GETBUFFERMETADATAREQUEST']._serialized_start=1220
  _globals['_GETBUFFERMETADATAREQUEST']._serialized_end=1246
  _globals['_QUERYDATAPOINTSREQUEST']._serialized_start=1248
  _globals['_QUERYDATAPOINTSREQUEST']._serialized_end=1352
  _globals['_QUERYDATAPOINTSRESPONSE']._serialized_start=1354
  _globals['_QUERYDATAPOINTSRESPONSE']._serialized_end=1417
  _globals['_QUERYEVENTSREQUEST']._serialized_start=1419
  _globals['_QUERYEVENTSREQUEST']._serialized_end=1498
  _globals['_QUERYEVENTSRESPONSE']._serialized_start=1500
  _globals['_QUERYEVENTSRESPONSE']._serialized_end=1555
  _globals['_GETBUFFERMETADATARESPONSE']._serialized_start=1557
  _globals['_GETBUFFERMETADATARESPONSE']._serialized_end=1631
  _globals['_HEALTHREQUEST']._serialized_start=1633
  _globals['_HEALTHREQUEST']._serialized_end=1648
  _globals['_HEALTHRESPONSE']._serialized_start=1650
  _globals['_HEALTHRESPONSE']._serialized_end=1666
  _globals['_GETCOMMANDREQUESTREQUEST']._serialized_start=1668
  _globals['_GETCOMMANDREQUESTREQUEST']._serialized_end=1718
  _globals['_GETCOMMANDREQUESTRESPONSE']._serialized_start=1720
  _globals['_GETCOMMANDREQUESTRESPONSE']._serialized_end=1790
  _globals['_SENDCOMMANDRESPONSEREQUEST']._serialized_start=1792
  _globals['_SENDCOMMANDRESPONSEREQUEST']._serialized_end=1865
  _globals['_SENDCOMMANDRESPONSERESPONSE']._serialized_start=1867
  _globals['_SENDCOMMANDRESPONSERESPONSE']._serialized_end=1896
  _globals['_GETCOMMANDREQUESTSTREAMREQUEST']._serialized_start=1898
  _globals['_GETCOMMANDREQUESTSTREAMREQUEST']._serialized_end=1954
  _globals['_GETCOMMANDREQUESTSTREAMRESPONSE']._serialized_start=1956
  _globals['_GETCOMMANDREQUESTSTREAMRESPONSE']._serialized_end=2032
  _globals['_GETTELEOPCONTROLDATASTREAMREQUEST']._serialized_start=2034
  _globals['_GETTELEOPCONTROLDATASTREAMREQUEST']._serialized_end=2092
  _globals['_GETTELEOPCONTROLDATASTREAMRESPONSE']._serialized_start=2094
  _globals['_GETTELEOPCONTROLDATASTREAMRESPONSE']._serialized_end=2185
  _globals['_GETTELEOPHEARTBEATSTREAMREQUEST']._serialized_start=2187
  _globals['_GETTELEOPHEARTBEATSTREAMREQUEST']._serialized_end=2220
  _globals['_GETTELEOPHEARTBEATSTREAMRESPONSE']._serialized_start=2222
  _globals['_GETTELEOPHEARTBEATSTREAMRESPONSE']._serialized_end=2318
  _globals['_GETTELEMETRYLISTENERSTREAMREQUEST']._serialized_start=2320
  _globals['_GETTELEMETRYLISTENERSTREAMREQUEST']._serialized_end=2378
  _globals['_GETTELEMETRYLISTENERSTREAMRESPONSE']._serialized_start=2380
  _globals['_GETTELEMETRYLISTENERSTREAMRESPONSE']._serialized_end=2456
  _globals['_POSTTRANSFORMFRAMERESPONSE']._serialized_start=2458
  _globals['_POSTTRANSFORMFRAMERESPONSE']._serialized_end=2486
  _globals['_SETBASEFRAMEIDREQUEST']._serialized_start=2488
  _globals['_SETBASEFRAMEIDREQUEST']._serialized_end=2523
  _globals['_SETBASEFRAMEIDRESPONSE']._serialized_start=2525
  _globals['_SETBASEFRAMEIDRESPONSE']._serialized_end=2549
  _globals['_CLEARTRANSFORMTREEREQUEST']._serialized_start=2551
  _globals['_CLEARTRANSFORMTREEREQUEST']._serialized_end=2578
  _globals['_CLEARTRANSFORMTREERESPONSE']._serialized_start=2580
  _globals['_CLEARTRANSFORMTREERESPONSE']._serialized_end=2608
  _globals['_CREATEEVENTREQUEST']._serialized_start=2610
  _globals['_CREATEEVENTREQUEST']._serialized_end=2662
  _globals['_CREATEEVENTRESPONSE']._serialized_start=2664
  _globals['_CREATEEVENTRESPONSE']._serialized_end=2685
  _globals['_CREATEEVENTERROR']._serialized_start=2687
  _globals['_CREATEEVENTERROR']._serialized_end=2724
  _globals['_GETTELEOPINFOREQUEST']._serialized_start=2726
  _globals['_GETTELEOPINFOREQUEST']._serialized_end=2748
  _globals['_GETTELEOPINFORESPONSE']._serialized_start=2750
  _globals['_GETTELEOPINFORESPONSE']._serialized_end=2799
  _globals['_POSTLANRTCOFFERREQUEST']._serialized_start=2801
  _globals['_POSTLANRTCOFFERREQUEST']._serialized_end=2840
  _globals['_POSTLANRTCOFFERRESPONSE']._serialized_start=2842
  _globals['_POSTLANRTCOFFERRESPONSE']._serialized_end=2883
  _globals['_SENDONCUSTOMDATACHANNELREQUEST']._serialized_start=2885
  _globals['_SENDONCUSTOMDATACHANNELREQUEST']._serialized_end=2956
  _globals['_SENDONCUSTOMDATACHANNELRESPONSE']._serialized_start=2958
  _globals['_SENDONCUSTOMDATACHANNELRESPONSE']._serialized_end=2991
  _globals['_GETCUSTOMDATACHANNELMESSAGESTREAMREQUEST']._serialized_start=2993
  _globals['_GETCUSTOMDATACHANNELMESSAGESTREAMREQUEST']._serialized_end=3064
  _globals['_GETCUSTOMDATACHANNELMESSAGESTREAMRESPONSE']._serialized_start=3066
  _globals['_GETCUSTOMDATACHANNELMESSAGESTREAMRESPONSE']._serialized_end=3165
  _globals['_AGENT']._serialized_start=3168
  _globals['_AGENT']._serialized_end=6146
# @@protoc_insertion_point(module_scope)
