from concurrent import futures
import grpc
import chatbot_pb2
import chatbot_pb2_grpc
import googleGenAI

class ChatbotServiceServicer(chatbot_pb2_grpc.ChatbotServiceServicer):
    def SayHello(self, request, context):
        return chatbot_pb2.HelloReply(message='Ping! ' + request.name)

    def GetBotData(self, request, context):
        return chatbot_pb2.ChatBotResponse(response=googleGenAI.generateText(request.request))

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    chatbot_pb2_grpc.add_ChatbotServiceServicer_to_server(ChatbotServiceServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print('Server running on port 50051.')
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
