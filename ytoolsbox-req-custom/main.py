from utils.LoggingUtils import LoggingHandler
import urllib3
import concurrent
import grpc
from proto import req_pb2_grpc, addDataClassify_pb2_grpc
from grpc_service.DSCService import DSCService
from grpc_service.addDataClassifyService import AddDataClassifyService
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

_HOST = '0.0.0.0'
_PORT = '2468'


def serve():
    # 指定最多可以有 4 个线程来处理请求
    grpcServer = grpc.server(
        concurrent.futures.ThreadPoolExecutor(max_workers=4))

    # 将相应的服务注册到 grpc 中
    req_pb2_grpc.add_DSCServiceServicer_to_server(
        DSCService(), grpcServer)
    addDataClassify_pb2_grpc.add_AddDataClassifyServiceServicer_to_server(
        AddDataClassifyService(), grpcServer)

    # 指定端口并且非 ssl 模式
    LoggingHandler.info(f"Python GRPC服务端启动:{_HOST + ':' + _PORT}")
    grpcServer.add_insecure_port(_HOST + ':' + _PORT)
    grpcServer.start()

    # 等待结束，不然进程会立刻退出
    grpcServer.wait_for_termination()


if __name__ == "__main__":
    serve()
