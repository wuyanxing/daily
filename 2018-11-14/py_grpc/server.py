import grpc
import data_pb2
import data_pb2_grpc
from concurrent import futures
import time

class person(data_pb2_grpc.templateServicer):
    def say(self, request, context):
        print(request.name)

        return  data_pb2.Status(status = 'status', ok = True)

def server():
    baseServer = grpc.server(futures.ThreadPoolExecutor(4))
    data_pb2_grpc.add_templateServicer_to_server(person(), baseServer)
    baseServer.add_insecure_port('localhost:8080')
    baseServer.start()

    while True:
        time.sleep(1)

server()
