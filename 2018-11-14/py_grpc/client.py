import grpc
import data_pb2_grpc
import data_pb2
def client():
    conn = grpc.insecure_channel('localhost:8080')
    baseClient = data_pb2_grpc.templateStub(channel=conn)
    respon = baseClient.say(data_pb2.Person(name="name", age="age", phone='phone'))
    print(respon)

client()
