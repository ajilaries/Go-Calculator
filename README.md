Go Calculator (CLI + Desktop UI)

A simple and modular calculator built using Go, with both a CLI version and a desktop UI powered by Wails.

🚀 Overview

This project showcases how Go can be used to build both backend logic and desktop applications using a clean and structured approach.

✨ Core Features
Modular calculator engine
Supports basic and advanced operations
CLI + Desktop UI support
Real-time calculation results
Custom error handling (e.g., divide by zero)
Clean and scalable project structure
🧮 Supported Operations

Basic:
+ Addition
- Subtraction
* Multiplication
/ Division

Advanced:
^ Power
% Modulus

Scientific:
sqrt Square root
log Natural log

🛠️ Technologies Used
Go (Golang) – Core backend logic
Wails – Desktop UI framework
HTML, CSS, JavaScript – UI design
Regex – Expression parsing
▶️ Running the Project
Desktop UI
wails dev
CLI Version
go run ./cmd/app
⚠️ Error Handling

Handles common errors like:

Invalid input
Division by zero
Unsupported operations
📦 Project Structure (Simplified)
parser → Input processing
engine → Calculation logic
operations → Math functions
ui → Frontend (Wails)
🚀 Future Improvements
Complex expressions (BODMAS)
Parentheses support
Calculation history
Keyboard input
Standalone executable
👨‍💻 Author

Developed by Ajil
