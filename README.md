# Chat Application

This is a simple chat application built with Go that allows users to send and receive chat messages. The application uses HTTP for communication and is containerized using Docker for easy deployment on Kubernetes.

## Features

- Send and receive chat messages
- Scalable deployment with Kubernetes
- Lightweight and efficient

## Getting Started

### Prerequisites

- Go 1.20 or later
- Docker
- Kubernetes cluster
- kubectl

### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/cezar45/chat-app.git
   cd chat-app
   ```
2. **Build the Docker Image**
    ```bash
   docker build -t chat-app:latest .
   ```
3. **Puse the Docker Image**

   ```bash 
   docker push chat-app:latest 
   ```
### Usage
#### Running Locally
Run the application without Docker:
```bash
go run cmd/main.go
```
#### Deploying to Kubernetes
Deploy the application to a Kubernetes cluster:

```bash
kubectl apply -f kubernetes/deployment.yaml
kubectl apply -f kubernetes/service.yaml
```
#### Access the Application
Use the following command to access the application:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"username":"Alice","content":"Hello, world!"}' http://<external-ip>/chat
```

### Contributing

Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bug fixes.

### License

This project is licensed under the MIT License.