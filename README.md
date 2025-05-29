# StackForge CLI Command Spec

> A full-featured CLI tool to generate, manage, and scale microservice architectures with Node.js, Go, MongoDB, and Docker.

---

## 🧱 `stackforge init`

**Description**: Initializes a new full-stack microservice project.

**Usage**:

```bash
stackforge init <project-name> [flags]
```

**Flags**:

* `--with-frontend` – include React + Tailwind frontend
* `--with-auth` – include prebuilt auth service
* `--lang=node|go` – preferred backend language
* `--docker` – generate docker-compose + Dockerfiles
* `--ci` – add CI/CD pipeline setup (e.g. GitHub Actions)

**Result**:
Creates:

```
<project-name>/
├── backend/
│   ├── services/
│   ├── gateway/
│   ├── docker/
│   ├── deployments/
├── frontend/ (optional)
├── docker-compose.yml
└── README.md
```

---

## 💠 `stackforge create-service`

**Description**: Generates a new microservice from a template.

**Usage**:

```bash
stackforge create-service <name> [--from=template] [--lang=node|go]
```

**Flags**:

* `--from=default|auth-service|custom-path`
* `--lang=node|go`
* `--port=3001`
* `--with-db` – scaffold DB config

**Result**:
Creates service in `backend/services/<name>-service`.

---

## 🧩 `stackforge add-endpoint`

**Description**: Adds a new API endpoint to a service.

**Usage**:

```bash
stackforge add-endpoint --service=user --method=POST --path=/register --name=RegisterUser
```

**Flags**:

* `--service=<name>`
* `--method=GET|POST|PUT|DELETE`
* `--path=/api/route`
* `--name=HandlerName`
* `--auth` – add auth middleware
* `--validate=SchemaName` – add validation
* `--test` – auto-generate test

**Result**:

* Adds controller function
* Adds route entry
* Adds service method
* Optionally adds test and middleware

---

## 📦 `stackforge add-model`

**Description**: Adds a database model with optional CRUD endpoints.

**Usage**:

```bash
stackforge add-model --service=product --name=Product --fields=name:string,price:number
```

**Flags**:

* `--service=<name>`
* `--name=<ModelName>`
* `--fields=<name:type,...>`
* `--crud` – scaffold CRUD routes

**Result**:

* Adds Mongoose/Go struct
* Adds optional CRUD controller + routes

---

## 🌐 `stackforge link-gateway`

**Description**: Connect a new service to the API gateway routing table.

**Usage**:

```bash
stackforge link-gateway --service=user --route=/api/user --port=3001
```

**Result**:

* Updates `gateway/routes.json` or `gateway.js`
* Optionally applies reverse proxy rule

---

## 📆 `stackforge generate` (alias: `g`)

**Subcommands**:

```bash
stackforge g controller --service=user --name=RegisterUser
stackforge g service --service=user --name=RegisterUser
stackforge g middleware --service=auth --name=verifyToken
```

---

## ✅ `stackforge test`

**Description**: Run unit/integration tests for a service.

**Usage**:

```bash
stackforge test --service=auth
```

**Flags**:

* `--watch` – watch mode
* `--coverage` – show test coverage

---

## 🚓 `stackforge docker-sync`

**Description**: Add the service to `docker-compose.yml` automatically.

**Usage**:

```bash
stackforge docker-sync --service=auth --port=3002
```

---

## 📦 `stackforge upgrade`

**Description**: Updates StackForge CLI or templates.

**Usage**:

```bash
stackforge upgrade
```

---

## 🧠 Future Ideas

* `stackforge deploy` – deploy locally or to cloud
* `stackforge publish` – upload/share service templates
* `stackforge config` – view/edit global CLI config

---

## 📘 About StackForge

StackForge is a CLI-first framework generator for fullstack microservice projects. It includes:

* ⚙️ Go + Node.js microservice scaffolding
* 🚪 Preconfigured API gateway (Express)
* 🎨 Optional React + Tailwind frontend
* 🛢 MongoDB support
* 🐳 Docker-ready templates
* 🧱 CLI tools to speed up service, route, model, and test generation

Built to help teams scale fast with clean architecture and automation.
