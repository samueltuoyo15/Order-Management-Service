# **Order Management Service** 📦

Welcome to the **Order Management Service**, a robust and scalable microservice application built with Go. This project showcases a clean architecture for handling order creation and retrieval, leveraging gRPC for efficient inter-service communication and a complementary HTTP API for broader accessibility. It’s designed to be a solid foundation for distributed systems, emphasizing clarity, performance, and maintainability.

---

## 🛠️ Installation

Getting this project up and running on your local machine is straightforward! Follow these steps to set up the environment and run the services:

*   **Clone the Repository**:
    ```bash
    git clone https://github.com/samueltuoyo15/Order-Management-Service.git
    cd Order-Management-Service
    ```
*   **Install Dependencies**:
    Go will automatically download the necessary modules defined in `go.mod`. Just run:
    ```bash
    go mod tidy
    ```
*   **Generate Protocol Buffer Files**:
    This project uses Protocol Buffers for defining service contracts. You'll need to generate the Go code from the `.proto` definitions. We've got a `Makefile` target for that!
    ```bash
    make proto
    ```
    This command creates the necessary Go files in `common/genproto/orders`.

---

## ▶️ Usage

Once installed, you can run the services and interact with them. This project consists of two main services: `order-service` (the backend logic) and `kitchen-service` (a frontend client).

1.  **Start the Order Service**:
    The `order-service` handles all order-related operations and exposes both gRPC and HTTP interfaces. It's designed to run both concurrently.
    ```bash
    make run-order
    ```
    This will start the HTTP server on `http://localhost:8000` and the gRPC server on `http://localhost:9000`. You'll see log messages indicating their readiness.

2.  **Start the Kitchen Service**:
    The `kitchen-service` acts as a simple client. It interacts with the `order-service` via gRPC to create and fetch orders, then displays them in a basic HTML interface.
    ```bash
    make run-kitchen
    ```
    This server will run on `http://localhost:3000`.

3.  **Interact with the Services**:

    *   **Via Kitchen Service (Browser)**:
        Open your web browser and navigate to `http://localhost:3000`. The `kitchen-service` will automatically make a gRPC call to `order-service` to create a sample order and then fetch all existing orders, displaying them on the page. Each refresh will likely create a new test order.

    *   **Directly with Order Service (HTTP API)**:
        You can also send HTTP requests directly to the `order-service`. For instance, to create a new order:
        ```bash
        curl -X POST -H "Content-Type: application/json" -d '{
            "customerId": 101,
            "productId": 202,
            "quantity": 3
        }' http://localhost:8000/orders
        ```
        This will add an order to the in-memory store of the `order-service`.

---

## ✨ Features

This project highlights several key aspects of modern service development:

*   **Microservices Architecture**: Clearly separated `order-service` and `kitchen-service` demonstrate modular design principles.
*   **gRPC Communication**: Utilizes gRPC for high-performance, contract-first inter-service communication, ensuring type safety and efficient data exchange.
*   **HTTP API**: Provides a RESTful HTTP interface for the order service, allowing easy integration with external clients.
*   **Protocol Buffers**: Employs Protocol Buffers for defining robust and backward-compatible service contracts, facilitating seamless data serialization.
*   **Structured Logging**: Integrates `slog` with `tint` for clear, readable, and structured logging, crucial for debugging and monitoring in production environments.
*   **In-Memory Persistence**: For demonstration purposes, orders are stored in a simple in-memory database, making setup quick and easy.

---

## 🚀 Technologies Used

| Technology         | Description                                     | Link                                         |
| :----------------- | :---------------------------------------------- | :------------------------------------------- |
| **Go**             | Core language for service development           | [golang.org](https://golang.org/)            |
| **gRPC**           | High-performance RPC framework                  | [grpc.io](https://grpc.io/)                  |
| **Protocol Buffers** | Language-neutral, platform-neutral extensible mechanism for serializing structured data | [developers.google.com/protocol-buffers](https://developers.google.com/protocol-buffers) |
| **`net/http`**     | Standard Go package for HTTP servers and clients | [pkg.go.dev/net/http](https://pkg.go.dev/net/http) |
| **`log/slog`**     | Structured logging for Go applications          | [pkg.go.dev/log/slog](https://pkg.go.dev/log/slog) |
| **`lmittmann/tint`** | Pretty, colored output for `slog`              | [github.com/lmittmann/tint](https://github.com/lmittmann/tint) |

---

## 📄 License

This project is open-source and available under the [License Name] – *add your license information here*.

---

## ✍️ Author

**Samuel Tuoyo**

- [LinkedIn](https://www.linkedin.com/in/samuel-tuoyo-8568b62b6)
- [X](https://x.com/TuoyoS26091)

---

![Go](https://img.shields.io/badge/Go-1.24.3-00ADD8?style=for-the-badge&logo=go&logoColor=white)

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)