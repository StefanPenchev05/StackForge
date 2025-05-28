# StackForge ⚙️🔥

> A full-stack monorepo template for building modern, scalable applications using microservices, React, MongoDB, and Docker.

---

## ✨ Features

- 🔧 **Microservices** using **Node.js** and **Golang**
- 🧱 **CLI Tool** (written in Go) to scaffold new backend services
- 🎨 **Frontend** using **React + TailwindCSS**
- 🛢️ **MongoDB** for persistent storage
- 🐳 **Dockerized** for development and deployment
- 📦 Easy to extend and customize

---

## 📁 Folder Structure

```txt
backend/
├── cli/            # Go CLI tool for generating new services
├── gateway/        # Node.js API Gateway
├── logger/         # Shared logging utilities
├── monitor/        # Basic monitoring setup (e.g., Prometheus)
├── services/       # Generated microservices live here

frontend/
├── public/
├── src/
│   ├── components/
│   └── pages/
├── tailwind.config.js
└── vite.config.js

docker/
├── docker-compose.yml
└── service Dockerfiles

.github/
└── template.yml
```
## 🚀 Getting Started

### Clone the Repo

```bash
git clone https://github.com/StefanPenchev05/stackforge.git
cd stackforge
```

### Run the CLI
```bash
cd backend/cli
go run main.go create-service user --lang=go
```

### Start Dev Environment (Docker)
```bash
docker-compose up --build
```

