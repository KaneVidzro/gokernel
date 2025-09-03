# GoKernel

A containerized Go-based platform with REST API, gRPC, PostgreSQL, Redis, and Nginx reverse proxy with automated TLS certificates via Let’s Encrypt.

## 🚀 Features

- REST API (`/`) served over HTTPS
- Native gRPC (`/grpc`) and gRPC-Web (`/grpc-web`) support
- PostgreSQL + Redis integration
- Dockerized environment with `docker compose`
- Automatic TLS certificates via **Certbot + Nginx**
- Developer-friendly **Makefile** for builds, migrations, and quality checks

---

## 📦 Requirements

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go 1.22+](https://go.dev/dl/) (for local builds/tests)
- [protoc](https://grpc.io/docs/protoc-installation/) (for regenerating gRPC stubs)
- [golangci-lint](https://golangci-lint.run/) (for linting)

---

## ⚙️ Setup

1. **Clone the repo**

   ```sh
   git clone https://github.com/kanevidzro/gokernel.git
   ```

2. **Copy `.env.example` to `.env`**

   ```sh
   cp .env.example .env
   ```

3. **Build and start the services**
   ```sh
   make up
   ```
4. **Access the API**

   ```sh
   API → https://api.example.com/
   gRPC → https://api.example.com/grpc
   gRPC-Web → https://api.example.com/grpc-web
   ```

   **Note:** The API and gRPC endpoints are served over HTTPS, so you’ll need to accept the certificate in your browser.

## 🔧 Useful Commands

```sh
make help        # Show available commands
make up          # Start all containers
make down        # Stop and remove containers
make restart     # Restart containers
make logs        # Tail logs
make ps          # Show running containers

make db-shell    # Connect to Postgres
make proto       # Regenerate protobuf/gRPC code
make lint        # Run lint checks
make test        # Run Go tests

make db-migrate  # Run database migrations

make cert-init   # Request initial Let's Encrypt certificate
make cert-renew  # Force renew all certificates

```

## 🔨 Contributing

```sh
1.  Fork the repo and create a feature branch
2. Run `make lint` and `make test` before pushing
3. Submit a pull request
```

## 📝 License

This project is licensed under the [MIT License](LICENSE).
