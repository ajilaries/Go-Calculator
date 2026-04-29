🧮 Go Calculator (CLI + Desktop UI)

A modern, modular calculator application built using Go with a clean architecture and a beautiful desktop UI powered by Wails.

🚀 Overview

This project demonstrates real-world Go application development, combining:

⚙️ A scalable backend (calculator engine)
🖥️ A desktop UI (HTML, CSS, JS via Wails)
🧠 Clean architecture (separation of concerns)
✨ Features
🧠 Backend
Modular calculator engine
Plug-and-play operation system
Custom error handling
Expression parsing (supports flexible input)
Unary + binary operations
🎨 UI (Wails)
Modern dark-themed calculator
Button-based input (grid layout)
Smooth interaction with backend
Real-time result display
🧮 Supported Operations
🔹 Basic
+   Addition
-   Subtraction
*   Multiplication
/   Division
🔹 Advanced
^   Power
%   Modulus
🔹 Scientific
sqrt   Square root
log    Natural log

⚙️ Installation
1️⃣ Clone the repository
git clone <your-repo-url>
cd go-calculator
2️⃣ Install dependencies
go mod tidy
3️⃣ Install Wails
go install github.com/wailsapp/wails/v2/cmd/wails@latest

wails doctor
▶️ Running the Application
🖥️ Run Desktop UI
cd G0Calculator
wails dev

👉 Launches the calculator window

💻 Run CLI Version
go run ./cmd/app
💡 Usage
UI Input (Buttons)
Click numbers and operators
Press = to calculate
Press C to clear
CLI Input
10 + 5
sqrt 25
log 10
⚠️ Error Handling

Handles:

Invalid input format
Division by zero
Unsupported operations

Example:

10 / 0 → Error: cannot divide by zero
🧠 Architecture
🔹 Layers
Layer	Responsibility
Parser	Converts input → Expression
Engine	Executes operations
Operations	Contains logic
UI (Wails)	User interaction
🔥 Key Concepts Used
Go modules & project structuring
Function registry pattern
Variadic functions
Expression parsing (regex-based)
Frontend-backend communication (Wails bridge)
🚀 Future Enhancements
 Complex expressions (10+5*2)
 Parentheses support
 Calculation history
 Keyboard input support
 Scientific calculator UI
 Export as standalone .exe
🧪 Example Inputs
7+3
8*5
sqrt25
log10
🤝 Contributing

Contributions are welcome!

Fork the repo
Create a new branch
Commit your changes
Open a Pull Request
📄 License

This project is licensed under the MIT License.

👨‍💻 Author

Developed by Ajil
