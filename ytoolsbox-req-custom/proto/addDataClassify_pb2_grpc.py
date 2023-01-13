# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import addDataClassify_pb2 as addDataClassify__pb2


class AddDataClassifyServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.AddDataClassify = channel.unary_unary(
            '/proto.AddDataClassifyService/AddDataClassify',
            request_serializer=addDataClassify__pb2.AddDataClassifyRequest.SerializeToString,
            response_deserializer=addDataClassify__pb2.AddDataClassifyResponse.FromString,
        )


class AddDataClassifyServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def AddDataClassify(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AddDataClassifyServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
        'AddDataClassify': grpc.unary_unary_rpc_method_handler(
            servicer.AddDataClassify,
            request_deserializer=addDataClassify__pb2.AddDataClassifyRequest.FromString,
            response_serializer=addDataClassify__pb2.AddDataClassifyResponse.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        'proto.AddDataClassifyService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))

 # This class is part of an EXPERIMENTAL API.


class AddDataClassifyService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def AddDataClassify(request,
                        target,
                        options=(),
                        channel_credentials=None,
                        call_credentials=None,
                        insecure=False,
                        compression=None,
                        wait_for_ready=None,
                        timeout=None,
                        metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.AddDataClassifyService/AddDataClassify',
                                             addDataClassify__pb2.AddDataClassifyRequest.SerializeToString,
                                             addDataClassify__pb2.AddDataClassifyResponse.FromString,
                                             options, channel_credentials,
                                             insecure, call_credentials, compression, wait_for_ready, timeout, metadata)