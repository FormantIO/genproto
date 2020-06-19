# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from protos.agent.v1 import agent_pb2 as protos_dot_agent_dot_v1_dot_agent__pb2
from protos.model.v1 import datapoint_pb2 as protos_dot_model_dot_v1_dot_datapoint__pb2
from protos.model.v1 import intervention_pb2 as protos_dot_model_dot_v1_dot_intervention__pb2
from protos.model.v1 import math_pb2 as protos_dot_model_dot_v1_dot_math__pb2


class AgentStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.StreamData = channel.stream_unary(
        '/v1.agent.Agent/StreamData',
        request_serializer=protos_dot_model_dot_v1_dot_datapoint__pb2.Datapoint.SerializeToString,
        response_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.StreamDataResponse.FromString,
        )
    self.PostData = channel.unary_unary(
        '/v1.agent.Agent/PostData',
        request_serializer=protos_dot_model_dot_v1_dot_datapoint__pb2.Datapoint.SerializeToString,
        response_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.PostDataResponse.FromString,
        )
    self.PostDataMulti = channel.unary_unary(
        '/v1.agent.Agent/PostDataMulti',
        request_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.PostDataMultiRequest.SerializeToString,
        response_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.PostDataMultiResponse.FromString,
        )
    self.CreateEvent = channel.unary_unary(
        '/v1.agent.Agent/CreateEvent',
        request_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.CreateEventRequest.SerializeToString,
        response_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.CreateEventResponse.FromString,
        )
    self.CreateInterventionRequest = channel.unary_unary(
        '/v1.agent.Agent/CreateInterventionRequest',
        request_serializer=protos_dot_model_dot_v1_dot_intervention__pb2.InterventionRequest.SerializeToString,
        response_deserializer=protos_dot_model_dot_v1_dot_intervention__pb2.InterventionRequest.FromString,
        )
    self.GetInterventionRequest = channel.unary_unary(
        '/v1.agent.Agent/GetInterventionRequest',
        request_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetInterventionRequestRequest.SerializeToString,
        response_deserializer=protos_dot_model_dot_v1_dot_intervention__pb2.InterventionRequest.FromString,
        )
    self.GetInterventionResponse = channel.unary_unary(
        '/v1.agent.Agent/GetInterventionResponse',
        request_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetInterventionResponseRequest.SerializeToString,
        response_deserializer=protos_dot_model_dot_v1_dot_intervention__pb2.InterventionResponse.FromString,
        )
    self.GetStreamsConfiguration = channel.unary_unary(
        '/v1.agent.Agent/GetStreamsConfiguration',
        request_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetStreamsConfigurationRequest.SerializeToString,
        response_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetStreamsConfigurationResponse.FromString,
        )
    self.GetApplicationConfiguration = channel.unary_unary(
        '/v1.agent.Agent/GetApplicationConfiguration',
        request_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetApplicationConfigurationRequest.SerializeToString,
        response_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetApplicationConfigurationResponse.FromString,
        )
    self.GetAgentConfiguration = channel.unary_unary(
        '/v1.agent.Agent/GetAgentConfiguration',
        request_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetAgentConfigurationRequest.SerializeToString,
        response_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetAgentConfigurationResponse.FromString,
        )
    self.Health = channel.unary_unary(
        '/v1.agent.Agent/Health',
        request_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.HealthRequest.SerializeToString,
        response_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.HealthResponse.FromString,
        )
    self.GetCommandRequest = channel.unary_unary(
        '/v1.agent.Agent/GetCommandRequest',
        request_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetCommandRequestRequest.SerializeToString,
        response_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetCommandRequestResponse.FromString,
        )
    self.GetCommandRequestStream = channel.unary_stream(
        '/v1.agent.Agent/GetCommandRequestStream',
        request_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetCommandRequestStreamRequest.SerializeToString,
        response_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetCommandRequestStreamResponse.FromString,
        )
    self.SendCommandResponse = channel.unary_unary(
        '/v1.agent.Agent/SendCommandResponse',
        request_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.SendCommandResponseRequest.SerializeToString,
        response_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.SendCommandResponseResponse.FromString,
        )
    self.PostTransformFrame = channel.unary_unary(
        '/v1.agent.Agent/PostTransformFrame',
        request_serializer=protos_dot_model_dot_v1_dot_math__pb2.TransformFrame.SerializeToString,
        response_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.PostTransformFrameResponse.FromString,
        )
    self.SetBaseFrameID = channel.unary_unary(
        '/v1.agent.Agent/SetBaseFrameID',
        request_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.SetBaseFrameIDRequest.SerializeToString,
        response_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.SetBaseFrameIDResponse.FromString,
        )


class AgentServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def StreamData(self, request_iterator, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PostData(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PostDataMulti(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateEvent(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateInterventionRequest(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetInterventionRequest(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetInterventionResponse(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetStreamsConfiguration(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetApplicationConfiguration(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetAgentConfiguration(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Health(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetCommandRequest(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetCommandRequestStream(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SendCommandResponse(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PostTransformFrame(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def SetBaseFrameID(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_AgentServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'StreamData': grpc.stream_unary_rpc_method_handler(
          servicer.StreamData,
          request_deserializer=protos_dot_model_dot_v1_dot_datapoint__pb2.Datapoint.FromString,
          response_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.StreamDataResponse.SerializeToString,
      ),
      'PostData': grpc.unary_unary_rpc_method_handler(
          servicer.PostData,
          request_deserializer=protos_dot_model_dot_v1_dot_datapoint__pb2.Datapoint.FromString,
          response_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.PostDataResponse.SerializeToString,
      ),
      'PostDataMulti': grpc.unary_unary_rpc_method_handler(
          servicer.PostDataMulti,
          request_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.PostDataMultiRequest.FromString,
          response_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.PostDataMultiResponse.SerializeToString,
      ),
      'CreateEvent': grpc.unary_unary_rpc_method_handler(
          servicer.CreateEvent,
          request_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.CreateEventRequest.FromString,
          response_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.CreateEventResponse.SerializeToString,
      ),
      'CreateInterventionRequest': grpc.unary_unary_rpc_method_handler(
          servicer.CreateInterventionRequest,
          request_deserializer=protos_dot_model_dot_v1_dot_intervention__pb2.InterventionRequest.FromString,
          response_serializer=protos_dot_model_dot_v1_dot_intervention__pb2.InterventionRequest.SerializeToString,
      ),
      'GetInterventionRequest': grpc.unary_unary_rpc_method_handler(
          servicer.GetInterventionRequest,
          request_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetInterventionRequestRequest.FromString,
          response_serializer=protos_dot_model_dot_v1_dot_intervention__pb2.InterventionRequest.SerializeToString,
      ),
      'GetInterventionResponse': grpc.unary_unary_rpc_method_handler(
          servicer.GetInterventionResponse,
          request_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetInterventionResponseRequest.FromString,
          response_serializer=protos_dot_model_dot_v1_dot_intervention__pb2.InterventionResponse.SerializeToString,
      ),
      'GetStreamsConfiguration': grpc.unary_unary_rpc_method_handler(
          servicer.GetStreamsConfiguration,
          request_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetStreamsConfigurationRequest.FromString,
          response_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetStreamsConfigurationResponse.SerializeToString,
      ),
      'GetApplicationConfiguration': grpc.unary_unary_rpc_method_handler(
          servicer.GetApplicationConfiguration,
          request_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetApplicationConfigurationRequest.FromString,
          response_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetApplicationConfigurationResponse.SerializeToString,
      ),
      'GetAgentConfiguration': grpc.unary_unary_rpc_method_handler(
          servicer.GetAgentConfiguration,
          request_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetAgentConfigurationRequest.FromString,
          response_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetAgentConfigurationResponse.SerializeToString,
      ),
      'Health': grpc.unary_unary_rpc_method_handler(
          servicer.Health,
          request_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.HealthRequest.FromString,
          response_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.HealthResponse.SerializeToString,
      ),
      'GetCommandRequest': grpc.unary_unary_rpc_method_handler(
          servicer.GetCommandRequest,
          request_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetCommandRequestRequest.FromString,
          response_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetCommandRequestResponse.SerializeToString,
      ),
      'GetCommandRequestStream': grpc.unary_stream_rpc_method_handler(
          servicer.GetCommandRequestStream,
          request_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetCommandRequestStreamRequest.FromString,
          response_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.GetCommandRequestStreamResponse.SerializeToString,
      ),
      'SendCommandResponse': grpc.unary_unary_rpc_method_handler(
          servicer.SendCommandResponse,
          request_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.SendCommandResponseRequest.FromString,
          response_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.SendCommandResponseResponse.SerializeToString,
      ),
      'PostTransformFrame': grpc.unary_unary_rpc_method_handler(
          servicer.PostTransformFrame,
          request_deserializer=protos_dot_model_dot_v1_dot_math__pb2.TransformFrame.FromString,
          response_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.PostTransformFrameResponse.SerializeToString,
      ),
      'SetBaseFrameID': grpc.unary_unary_rpc_method_handler(
          servicer.SetBaseFrameID,
          request_deserializer=protos_dot_agent_dot_v1_dot_agent__pb2.SetBaseFrameIDRequest.FromString,
          response_serializer=protos_dot_agent_dot_v1_dot_agent__pb2.SetBaseFrameIDResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'v1.agent.Agent', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
