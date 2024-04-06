import grpc
import grpcProto.grpcChatbot_pb2 as grpcChatbot_pb2
import grpcProto.grpcChatbot_pb2_grpc as grpcChatbot_pb2_grpc

class ChatbotServicer(grpcChatbot_pb2_grpc.ImageServiceServicer):
    def SendMessage(self, request, context):
        return grpcChatbot_pb2.MessageResponse(reply="Echo: " + request.message)

def serve():
    server = grpc.server(grpc.ThreadPoolExecutor(max_workers=10))
    grpcChatbot_pb2_grpc.add_ChatbotServicer_to_server(ChatbotServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
