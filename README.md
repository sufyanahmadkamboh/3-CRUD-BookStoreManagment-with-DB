# 📚 3-CRUD-BookStoreManagment-with-DB

A simple **Go (Golang)** project demonstrating **CRUD operations** for a Book Store Management System using **Gorilla Mux** and **GORM** with support for **PostgreSQL** and **MySQL** databases.

---

## 🚀 Project Setup

### 1. **Initialize Go Module**

```sh
go mod init github.com/sufyanahmadkamboh/3-CRUD-BookStoreManagment-with-DB
```

---

## 📦 Install Dependencies

### Gorilla Mux (Router)

```sh
go get github.com/gorilla/mux
```

### GORM ORM

```sh
go get gorm.io/gorm
```

### PostgreSQL Driver

```sh
go get gorm.io/driver/postgres
```

### MySQL Driver

```sh
go get gorm.io/driver/mysql
```

---

## 🗄️ Database Table Structure

Create a table named **Books** with three fields:

```sql
CREATE TABLE Books (
    id SERIAL PRIMARY KEY,           -- Or AUTO_INCREMENT for MySQL
    name VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    publication VARCHAR(255) NOT NULL
);
```

**Fields:**

* `Name` – Name of the book
* `Author` – Author of the book
* `Publication` – Publication/Publisher name

---

## 📚 Features (CRUD)

This project covers:

* **Create** a new Book
* **Read** all Books or a single Book
* **Update** a Book by ID
* **Delete** a Book by ID

---

## ▶️ How to Run

```sh
go run .\cmd\main\main.go
```

API will start on your configured port (commonly `localhost:8080`).

---

## 🛠 Tech Stack

| Tool               | Purpose              |
| ------------------ | -------------------- |
| Go                 | Programming language |
| Gorilla Mux        | HTTP router          |
| GORM               | ORM for Go           |
| PostgreSQL / MySQL | Database             |

---

## 🤝 Contributing

Feel free to fork the repo, create issues, and submit PRs.

---

## 📄 License

This project is licensed under the MIT License.

---

Happy Coding! 🚀📚
