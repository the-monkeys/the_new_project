# To build an API

Since we know that this project uses http for communicating with clients and gRPC for communicating with microservices. All the request coming from a client will fall on the `API Gateway` and the `API Gateway` will dial to the services using gRPC protocol. We need to follow the following steps to build and API.

1. Define proto messages and service in `services/api_gateway/pkg` service in the `pb` directory.

2. Define the same proto messages and service in `services/article_and_post/pkg/pb` directory if you are writing an API for arts and service.

3. Run the command `make proto` to generate the boiler place code and implement both the client and server.
    3.1 Implement the client in `services/api_gateway/article`.
    
    3.2 Implement the server in `services/article_and_post/pkg/service`