# Notification Service

## Problem Statement

Develop a notification service that facilitates the sending of various types of notifications, including email, Slack, and in-app notifications.

### Requirements
1. **Notification Types**
   - Send notifications through different channels: Email, Slack, and in-app.
2. **Channel Mapping**
   - Route notifications to the appropriate channel based on type.
3. **Notification Scheduling**
   - Allow notifications to be sent instantly or scheduled for later.
4. **Notification Templates**
   - Support predefined and user-defined templates for notifications.
5. **Flexibility and Creativity**
   - Open-ended design to allow creative solutions and best practices.

---

## Solution Overview

### Architecture
- **Gin Framework** is used for HTTP API endpoints.
- **Modular Structure**: Separate packages for server, scheduler, sender, templates, and clients.
- **Database**: Templates and notification schedules are persisted in a relational database (e.g., PostgreSQL).

### Key Features
- **Multi-Channel Support**: Email, Slack, and in-app notifications via pluggable client interfaces.
- **Channel Mapping**: Each notification type is mapped to one or more delivery channels.
- **Scheduling**: Notifications can be sent immediately or scheduled for future delivery using a scheduler service.
- **Templates**: CRUD APIs for managing notification templates, supporting both system and user-defined templates.
- **Extensible**: Easily add new channels or notification types by implementing the respective interfaces.

### API Endpoints
- `/notifications` (POST): Create a notification (immediate or scheduled)
- `/notifications/{id}` (GET, PUT): Get or update a notification
- `/templates` (POST): Create a template
- `/templates/{id}` (GET, PUT, DELETE): Manage templates

See `swagger.json` for the full OpenAPI specification.

### Technologies Used
- Go (Golang)
- Gin Web Framework
- PostgreSQL (or compatible SQL DB)
- OpenAPI (Swagger) for API documentation

### How to Run
1. Clone the repository
2. Configure your database and environment variables
3. Run `go mod tidy` to install dependencies
4. Start the server: `go run server/main.go`

---

## Folder Structure
- `src/clients/` - Channel clients (email, Slack, app)
- `src/scheduler/` - Scheduling logic and repository
- `src/sender/` - Notification sending logic
- `src/server/` - API handlers, services, and repositories
- `src/templates/` - Template management

---

## License
MIT
