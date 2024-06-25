import grpc
from concurrent import futures
import time

import cctv_pb2
import cctv_pb2_grpc
import decoder

class CCTVService(cctv_pb2_grpc.CCTVServiceServicer):
    """
    Provides methods that implement functionality of CCTV service.
    """
    
    def StreamVideo(self, request_iterator, context):
        """
        Receives a stream of video frames and processes them.
        
        Args:
            request_iterator (iterator): An iterator of incoming video frames.
            context (grpc.ServicerContext): The context of the gRPC call.
        
        Yields:
            cctv_pb2.Acknowledgement: An acknowledgment message for each received frame.
        """
        # Create an instance of the H264Decoder with preview enabled
        decoder_instance = decoder.H264Decoder(preview=True)
        for frame in request_iterator:
            # Decode the received frame
            decoder_instance.decode(frame)
            # Yield an acknowledgment for the received frame
            yield cctv_pb2.Acknowledgement(message=f"Received frame with length: {len(frame.frame)}")

def serve():
    """
    Starts the gRPC server and listens for incoming requests.
    """
    # Create a gRPC server with a thread pool of 10 workers
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    # Add the CCTVService to the gRPC server
    cctv_pb2_grpc.add_CCTVServiceServicer_to_server(CCTVService(), server)
    # Bind the server to port 50051
    server.add_insecure_port('[::]:50051')
    # Start the gRPC server
    server.start()
    print("gRPC server is running on port 50051")
    
    try:
        # Keep the server running indefinitely
        while True:
            time.sleep(86400)  #TODO: Here we could do other things
    except KeyboardInterrupt:
        # Stop the server gracefully on interrupt
        server.stop(0)

if __name__ == '__main__':
    serve()