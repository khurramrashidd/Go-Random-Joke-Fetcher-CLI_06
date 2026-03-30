# 🤣 Go Random Joke Fetcher CLI (Day 6)

A fun and lightweight **CLI-based Joke Fetcher built using Go (Golang)**. This tool connects to a public API over the internet, retrieves a random programming or general joke, and displays it in the terminal.

This project is part of my **Go learning journey (Day 6 Project)**, stepping into the world of Networking and JSON parsing!

---

## 🚀 Features

* Makes real-time HTTP GET requests to a public API
* Parses and extracts specific data from JSON responses
* Handles network errors gracefully
* Fast, lightweight, and entertaining CLI experience

---

## 🛠 Built With

**Go (Golang)**
* Standard Library (`net/http`, `encoding/json`, `io`, `fmt`)

---

## 📂 Project Structure

```text
go_cli_06_joke_fetcher/
│
├── main.go
├── go.mod
├── README.md
```

---

## ⚙️ Installation

### 1. Clone the repository

```bash
git clone [https://github.com/yourusername/go_cli_06_joke_fetcher.git](https://github.com/yourusername/go_cli_06_joke_fetcher.git)
```

### 2. Navigate into the folder

```bash
cd go_cli_06_joke_fetcher
```

---

## ▶️ Usage

Run the application:

```bash
go run main.go
```

The CLI will instantly reach out to the internet, grab a joke, and display the Setup and Punchline.

---

## 📸 Example Output
<img width="879" height="200" alt="image" src="https://github.com/user-attachments/assets/2bb01d53-6527-4049-b4f6-481bfa78c77d" />

---

## 🧠 Concepts Practiced

* **Networking:** Using `net/http` to make HTTP GET requests.
* **Struct Tags:** Mapping JSON keys to Go Struct fields using `` `json:"key"` ``.
* **JSON Parsing:** Using `encoding/json` to Unmarshal byte streams into usable Go data structures.
* **Resource Management:** Using `defer` to ensure network connections are closed properly to avoid memory leaks.
* **Error Handling:** Managing network dropouts and parsing failures.

---

## 🔐 Best Practices Followed

* Structuring data strictly with Go Structs.
* Deferring `.Close()` methods immediately after opening resources.
* Meaningful error outputs instead of silent crashes.

---

## 🔮 Future Improvements

* Allow users to select a specific category of jokes (Programming, Dad jokes, etc.).
* Save favorite jokes to a local `.txt` file.
* Add a loading animation while the HTTP request is pending.

---

## 📜 License

This project is open source and available under the **MIT License**.

---

## 👨‍💻 Author

**Khurram Rashid** B.Tech Computer Science Engineering  
Amity University Mumbai
```
