# Chatbot

```For Generating Python Proto Files: python -m grpc_tools.protoc -I./grpcProto --python_out=./slmServer --grpc_python_out=./slmServer ./grpcProto/chatbot.proto
For Generating Go Proto Files: protoc --proto_path=.  --go_out=./server/client --go_opt=paths=source_relative   --go-grpc_out=./server/client --go-grpc_opt=paths=source_relative grpcProto/chatbot.proto```

#Window
```python -m venv venv                 
venv\Scripts\activate```
