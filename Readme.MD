
# Microservices Project

This project demonstrates contract testing between microservices using Keploy. It consists of four microservices:

User Service – Manages user data

Order Service – Handles order processing

Payment Service – Processes payments

Notification Service – Sends order and payment status notifications to users





## Tech Stack

**Go (Golang), PostgreSQL
, Docker , Keploy**


## Setup

```bash
git clone https://github.com/Lokeshranjan8/microservices.git
cd microservices

```

    1) Ensure PostgreSQL databases (user_db, order_db, etc.) are     created via CLI or pgAdmin.

    2) Run each service individually:

```bash
cd user-svc && keploy record -c "go run main.go"
# Repeat for each service
```
    Test endpoints using Postman or curl.

    Keploy will auto-generate test cases in /keploy/test-set-*.

## Running Tests

To run tests, run the following command

```bash
  keploy test -c "go run main.go"

```

