import grpc

from calculator_pb2_grpc import CaculatorStub
from calculator_pb2 import Input

if __name__ == '__main__':
    channel = grpc.insecure_channel('localhost:8081')
    stub = CaculatorStub(channel)
    input = Input(left=1, right=10)
    output = stub.Add(input)
    print("received gRPC response from server, ansewr is {}".format(output.value))
