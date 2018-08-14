# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import grpcapi_pb2 as grpcapi__pb2


class IAMStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.DeepPing = channel.unary_unary(
        '/iamapi.IAM/DeepPing',
        request_serializer=grpcapi__pb2.PrimitiveVoid.SerializeToString,
        response_deserializer=grpcapi__pb2.PrimitiveString.FromString,
        )
    self.CheckAuthentication = channel.unary_unary(
        '/iamapi.IAM/CheckAuthentication',
        request_serializer=grpcapi__pb2.CheckAuthenticationInput.SerializeToString,
        response_deserializer=grpcapi__pb2.User.FromString,
        )
    self.GetHierarchyRelations = channel.unary_unary(
        '/iamapi.IAM/GetHierarchyRelations',
        request_serializer=grpcapi__pb2.GetHierarchyRelationsInput.SerializeToString,
        response_deserializer=grpcapi__pb2.GetHierarchyRelationsOutput.FromString,
        )
    self.GetEventRecords = channel.unary_unary(
        '/iamapi.IAM/GetEventRecords',
        request_serializer=grpcapi__pb2.GetEventRecordsInput.SerializeToString,
        response_deserializer=grpcapi__pb2.GetEventRecordsOutput.FromString,
        )
    self.GetEventRecordStream = channel.unary_stream(
        '/iamapi.IAM/GetEventRecordStream',
        request_serializer=grpcapi__pb2.PrimitiveVoid.SerializeToString,
        response_deserializer=grpcapi__pb2.PrimitiveBytes.FromString,
        )


class IAMServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def DeepPing(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CheckAuthentication(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetHierarchyRelations(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetEventRecords(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetEventRecordStream(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_IAMServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'DeepPing': grpc.unary_unary_rpc_method_handler(
          servicer.DeepPing,
          request_deserializer=grpcapi__pb2.PrimitiveVoid.FromString,
          response_serializer=grpcapi__pb2.PrimitiveString.SerializeToString,
      ),
      'CheckAuthentication': grpc.unary_unary_rpc_method_handler(
          servicer.CheckAuthentication,
          request_deserializer=grpcapi__pb2.CheckAuthenticationInput.FromString,
          response_serializer=grpcapi__pb2.User.SerializeToString,
      ),
      'GetHierarchyRelations': grpc.unary_unary_rpc_method_handler(
          servicer.GetHierarchyRelations,
          request_deserializer=grpcapi__pb2.GetHierarchyRelationsInput.FromString,
          response_serializer=grpcapi__pb2.GetHierarchyRelationsOutput.SerializeToString,
      ),
      'GetEventRecords': grpc.unary_unary_rpc_method_handler(
          servicer.GetEventRecords,
          request_deserializer=grpcapi__pb2.GetEventRecordsInput.FromString,
          response_serializer=grpcapi__pb2.GetEventRecordsOutput.SerializeToString,
      ),
      'GetEventRecordStream': grpc.unary_stream_rpc_method_handler(
          servicer.GetEventRecordStream,
          request_deserializer=grpcapi__pb2.PrimitiveVoid.FromString,
          response_serializer=grpcapi__pb2.PrimitiveBytes.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'iamapi.IAM', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
