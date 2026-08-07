# Go Email Service

A robust, lightweight gRPC microservice built in Go for managing email templates, sending transactional emails via the Brevo API, and logging outgoing email events to a PostgreSQL database using GORM.

---

## 📑 Table of Contents

- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Architecture & Directory Structure](#-architecture--directory-structure)
- [Prerequisites](#-prerequisites)
- [Environment Variables](#-environment-variables)
- [Getting Started](#-getting-started)
  - [1. Clone Repository & Install Dependencies](#1-clone-repository--install-dependencies)
  - [2. Environment Configuration](#2-environment-configuration)
  - [3. Running the Server](#3-running-the-server)
- [gRPC API Specifications](#-grpc-api-specifications)
  - [Services & RPC Methods](#services--rpc-methods)
  - [Message Schemas](#message-schemas)
- [Testing with `grpcurl`](#-testing-with-grpcurl)
  - [Add an Email Template](#1-add-an-email-template)
  - [Send an Email](#2-send-an-email)
- [Database Models](#-database-models)
- [License](#-license)

---

## ✨ Features

- **gRPC API**: Fast, strongly-typed gRPC endpoints for adding templates and sending email notifications.
- **Template Management**: Store static and dynamic HTML email templates in PostgreSQL.
- **Transactional Email Integration**: Sends emails using [Brevo](https://www.brevo.com/) (formerly Sendinblue) HTTP API.
- **Outgoing Email Audit Log**: Automatically records sent emails and recipient metrics in PostgreSQL.
- **Input Validation**: Enforces string and email format constraints using [`buf.validate`](https://github.com/bufbuild/protovalidate).
- **Clean Architecture**: Separated into API, business service logic, data models, repositories, and configuration layers.

---

## 🛠 Tech Stack

- **Language**: [Go 1.25+](https://go.dev/)
- **API Protocol**: [gRPC](https://grpc.io/) & [Protocol Buffers (v3)](https://protobuf.dev/)
- **Database / ORM**: [PostgreSQL](https://www.postgresql.org/) & [GORM](https://gorm.io/)
- **Email Service Provider**: [Brevo SDK / REST API](https://github.com/getbrevo/brevo-go)
- **Validation**: [buf.build protovalidate](https://github.com/bufbuild/protovalidate)
- **Messaging (Optional Config)**: [RabbitMQ (amqp091-go)](https://github.com/rabbitmq/amqp091-go)

---

## 🏗 Architecture & Directory Structure

```text
go-email-service/
├── api/                        # Generated gRPC & Protobuf Go code
│   └── email/
│       ├── emailservice.pb.go
│       └── emailservice_grpc.pb.go
├── cmd/                        # Application entry points
│   └── server/
│       └── main.go             # Main server bootstrap
├── internal/                   # Private application and business logic
│   ├── config/                 # Service dependencies (DB, Brevo, RabbitMQ, Interceptor)
│   │   ├── brevo.go
│   │   ├── database.go
│   │   ├── loggingInterceptor.go
│   │   └── rabbitmq.go
│   ├── data/                   # Persistence layer
│   │   ├── models/             # GORM database entities (EmailTemplate, OutgoingEmail)
│   │   └── repositories/       # Data access interfaces and implementations
│   └── service/                # gRPC server implementation (EmailServer)
├── proto/                      # Protocol Buffer definitions
│   └── email/
│       └── emailservice.proto
├── third_party/                # Protobuf third-party dependencies (protovalidate)
├── .env.example                # Sample environment configuration file
├── Dockerfile                  # Container build file
├── go.mod                      # Go module definition
└── go.sum                      # Go dependency lock file
```

---

## 📋 Prerequisites

Before running the application, ensure you have the following installed:

- **Go**: `v1.25` or higher
- **PostgreSQL**: A running instance of PostgreSQL database
- **Brevo Account**: Active Brevo account and an API key with sending capabilities

---

## ⚙️ Environment Variables

Create a `.env` file in the root of the project using `.env.example` as a template:

```bash
cp .env.example .env
```

Configure the following variables in `.env`:

| Variable | Description | Example / Default |
| :--- | :--- | :--- |
| `DB_HOST` | PostgreSQL host address | `localhost` |
| `DB_USER` | PostgreSQL user | `postgres` |
| `DB_PASSWORD` | PostgreSQL user password | `yourpassword` |
| `DB_NAME` | Database name | `email_service_db` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `BREVO_API_KEY` | Brevo API key | `xkeysib-xxxxxxxx...` |
| `BREVO_URL` | Brevo transactional email endpoint URL | `https://api.brevo.com/v3/smtp/email` |
| `SENDER_EMAIL` | Verified sender email address | `no-reply@yourdomain.com` |
| `RABBITMQ_URL` *(optional)* | RabbitMQ connection string | `amqp://guest:guest@localhost:5672/` |

---

## 🚀 Getting Started

### 1. Clone Repository & Install Dependencies

```bash
git clone https://github.com/TheDurodola/go-email-service.git
cd go-email-service
go mod download
```

### 2. Environment Configuration

Populate your `.env` file with your database credentials and Brevo API details.

### 3. Running the Server

Start the gRPC server:

```bash
go run cmd/server/main.go
```

By default, the gRPC server listens on TCP port `:50051`. Upon startup, GORM auto-migrates the database schemas (`EmailTemplate` and `OutgoingEmail`).

---

## 🔌 gRPC API Specifications

The gRPC definitions are declared in [`proto/email/emailservice.proto`](file:///home/bolaji/Documents/GITHUB/go-email-service/proto/email/emailservice.proto).

### Services & RPC Methods

#### `EmailService`

| Method | Request | Response | Description |
| :--- | :--- | :--- | :--- |
| `AddEmailTemplate` | [`EmailTemplateRequest`](#emailtemplaterequest) | [`EmailTemplateResponse`](#emailtemplateresponse) | Registers a new HTML email template in the database. |
| `SendEmail` | [`SendEmailRequest`](#sendemailrequest) | [`SendEmailResponse`](#sendemailresponse) | Retrieves a template by name, dispatches the email via Brevo, and logs the outgoing email. |

---

### Message Schemas

#### `EmailTemplateRequest`
```protobuf
enum TemplateType {
  STATIC = 0;
  DYNAMIC = 1;
}

message EmailTemplateRequest {
  string template_name = 1;
  TemplateType template_type = 2;
  string template_body = 3;
  string template_subject = 4;
}
```

#### `EmailTemplateResponse`
```protobuf
message EmailTemplateResponse {
  bool is_added = 1;
  string message = 2;
}
```

#### `SendEmailRequest`
```protobuf
message SendEmailRequest {
  string firstname = 1;
  string to = 2;              // Must be a valid email format
  string template_name = 3;
  string app_name = 4;
}
```

#### `SendEmailResponse`
```protobuf
message SendEmailResponse {
  bool is_sent = 1;
}
```

---

## 🧪 Testing with `grpcurl`

You can test the gRPC server directly using [`grpcurl`](https://github.com/fullstorydev/grpcurl).

### 1. Add an Email Template

```bash
grpcurl -plaintext -d '{
  "template_name": "welcome_email",
  "template_type": "STATIC",
  "template_subject": "Welcome to Our Platform!",
  "template_body": "<h1>Welcome</h1><p>Thank you for signing up!</p>"
}' localhost:50051 emailservice.EmailService/AddEmailTemplate
```

### 2. Send an Email

```bash
grpcurl -plaintext -d '{
  "firstname": "John",
  "to": "user@example.com",
  "template_name": "welcome_email",
  "app_name": "MyAwesomeApp"
}' localhost:50051 emailservice.EmailService/SendEmail
```

---

## 🗄 Database Models

### `EmailTemplate`
Stores reusable HTML email templates.

| Field | Type | Description |
| :--- | :--- | :--- |
| `id` | `uint` | Primary key (GORM) |
| `name` | `string` | Unique template identifier (e.g. `welcome_email`) |
| `type` | `string` | Template type (`STATIC` or `DYNAMIC`) |
| `subject` | `string` | Email subject line |
| `body` | `string` | HTML content of the email |

### `OutgoingEmail`
Logs outgoing emails for auditing and analytics.

| Field | Type | Description |
| :--- | :--- | :--- |
| `id` | `uint` | Primary key (GORM) |
| `template_name` | `string` | Template used for sending |
| `app_name` | `string` | Application triggering the email |
| `recipient` | `string` | Target email recipient |
| `first_name` | `string` | Recipient's first name |

---

## 📄 License

This project is open source and available under the standard MIT license.
