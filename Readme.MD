Microservices Project
A microservices-based architecture consisting of:

user-svc

order-svc

payment-svc

notification-svc

📦 Tech Stack
Go (Golang)

PostgreSQL

Docker (optional)

Keploy – API testing & test data generation

🚀 Services Overview
Service	Description	Port
user-svc	Handles user creation and auth	8001
order-svc	Manages orders	8002
payment-svc	Handles payments	8003
notification-svc	Sends notifications	8004
⚙️ Setup
1️⃣ Clone the Repo
bash
Copy
Edit
git clone https://github.com/<your-org>/microservices.git
cd microservices
2️⃣ Start Each Service (in separate terminals)
bash
Copy
Edit
cd user-svc && keploy record -c "go run main.go"
Repeat the above for:

order-svc

payment-svc

notification-svc

3️⃣ Database Setup
Ensure PostgreSQL databases like user_db, order_db, etc., are created manually using pgAdmin or CLI.

4️⃣ Keploy Test Capture
After hitting APIs via Postman, Keploy test cases are saved here:

bash
Copy
Edit
keploy/test-set-*/tests/
🧪 Running Tests
To replay the captured tests:

bash
Copy
Edit
keploy test -c "go run main.go"
📁 Ignore & Clean-up
These are excluded in .gitignore:

bash
Copy
Edit
keploy/test-set-*
keploy/test-report*
*.pcap
keploy-logs.txt