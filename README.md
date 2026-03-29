⚖️ CLI Weight Converter (Go)

A simple command-line application built with Go that converts weights between kilograms and grams.

This project focuses on core Go fundamentals — control flow, user input handling, and clean function design.

🚀 Features
Convert Kilograms → Grams
Convert Grams → Kilograms
Interactive CLI menu
Continuous loop until user exits
Simple and fast execution
🧠 What This Project Teaches

If you built this properly, you practiced:

switch statements (real control flow, not toy examples)
Handling user input with fmt.Scanln
Function design and separation of concerns
Avoiding variable misuse (or at least confronting it)
Basic CLI program structure
📦 Project Structure
.
├── main.go
├── converter.go   (optional if you separated logic)
└── README.md
⚙️ How to Run

Make sure you have Go installed:

go version

Run the program:

go run main.go

Or build it:

go build -o weight-converter
./weight-converter
🖥️ Usage
=====Modern Weight Converter====
1...continue......2...Exit

After continuing:

Select An Option
1...Convert Kilogram To Gram
2...Convert Gram To Kilogram
3...Exit
🧮 Example
Enter value:
5
→ 5000 Gram
Enter value:
2000
→ 2 Kilogram
🔧 Core Functions
func ConvertKilogramtoGram(kg float64) float64 {
    return kg * 1000
}

func ConvertGramToKilogram(g float64) float64 {
    return g / 1000
}
⚠️ Current Limitations

Let’s be honest — this isn’t production-level yet:

No input validation (user can break it easily)
Accepts invalid types without handling errors
UI is basic and rigid
No support for more units (pounds, ounces, etc.)
🧠 What You Should Improve Next

If you stay at this level, you’re just playing with Go — not mastering it.

Next upgrades:

Input validation
Prevent crashes on bad input
Error handling
Handle Scanln failures properly
Refactor structure
Separate logic into packages
Add more units
kg ↔ lb, oz, etc.
Make it testable
Write unit tests for conversion functions
Improve UX
Clear screen, better prompts, retry on invalid input
🧱 Reality Check

This project is beginner-level.

That’s fine — but don’t pretend it’s more than it is.

What matters is what you do next:

Stack more complexity
Add constraints
Break and rebuild it cleaner