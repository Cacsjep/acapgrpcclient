import grpc
from concurrent import futures
import time

import cctv_pb2
import cctv_pb2_grpc

class CCTVService(cctv_pb2_grpc.CCTVServiceServicer):
    def StreamVideo(self, request_iterator, context):
        for frame in request_iterator:
            print(f"Received frame: {frame.frame}")
            yield cctv_pb2.Acknowledgement(message=f"Received frame with length: {len(frame.frame)}")

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    cctv_pb2_grpc.add_CCTVServiceServicer_to_server(CCTVService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("gRPC server is running on port 50051")
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()