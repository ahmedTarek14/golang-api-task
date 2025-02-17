# HTTP Web Server for Users and Products Management

## 🚀 About the Project
This is a backend service built using Go and PostgreSQL for managing users and products. It provides authentication, credit card management, and product purchasing functionalities.

## 🛠️ Technologies Used
- Go (Golang)
- PostgreSQL (Using plain SQL, no ORM)
- GorillaMux (For routing)
- Docker (For database setup)
- JWT (For authentication)

## 📌 Setup Instructions

### ❶️ Clone the repository
```sh
git clone https://github.com/your-username/go-api-task.git
cd go-api-task
```

### ❷️ Install Go dependencies
```sh
go mod tidy
```

### ❸️ Setup the Database (using Docker)
```sh
docker-compose up -d
```

### ❹️ Run the project
```sh
go run main.go
```

or build it:
```sh
go build -o app.exe .
./app.exe
```

## 🔥 Implemented Endpoints
| Endpoint               | Method | Description               | Authentication |
|------------------------|--------|---------------------------|---------------|
| `/api/user/register`   | POST   | Register a new user       | ❌ No         |
| `/api/user/login`      | POST   | Login user & get token    | ❌ No         |
| `/api/user/credit-card/add` | POST   | Add a new credit card     | ✅ Yes        |
| `/api/user/credit-card/delete` | DELETE | Delete a credit card | ✅ Yes        |
| `/api/user/products` | GET | List Existing Products | ❌ No        |

## 📌 Next Steps
- Implement product management (Create, Update, Delete)
- Integrate Stripe for payments
- Add filtering for sales reports

---

## 📝 Notes
- Make sure you have Docker installed to run PostgreSQL easily.
- The authentication system uses JWT, so include the token in `Authorization: Bearer <TOKEN>` for secured endpoints.

