# Word Sorter Project

A Go application that sorts words based on the number of 'a' characters they contain in descending order. If words have the same number of 'a' characters, they are sorted by length.

## Table of Contents
- [Features](#features)
- [Project Structure](#project-structure)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Development](#development)
- [Testing](#testing)
- [Git Workflow](#git-workflow)
- [Contributing](#contributing)

## Features
- Sorts words by number of 'a' characters in descending order
- Secondary sorting by word length
- Accepts input via command-line arguments or standard input
- Handles various input scenarios and error cases
- Clean and maintainable code structure

## Project Structure
word_sorter_project/
├── cmd/
│   └── main.go          # Application entry point
├── internal/
│   └── sorter/
│       ├── sorter.go    # Core sorting logic
│       └── sorter_test.go # Unit tests
├── go.mod              # Go module file
└── README.md          # Documentation

## Requirements
- Go 1.20 or higher
- Git

## Installation
1. Clone the repository:
git clone https://github.com/yourusername/word_sorter_project.git

2. Navigate to project directory:
cd word_sorter_project

3. Build the project:
go build ./cmd/main.go

## Usage
The program accepts input in two ways:

1. Command-line arguments:
go run cmd/main.go word1 word2 word3
Example: go run cmd/main.go aaaasd a aab aaabcd ef

2. Standard input (interactive mode):
go run cmd/main.go
Then type words (one per line)
Press Ctrl+D (Unix) or Ctrl+Z (Windows) to finish input

Example output:
Input: aaaasd a aab aaabcd ef
Sorted: aaaasd aaabcd aab a ef

## Development
### Code Structure
- All function and variable names follow Go naming conventions
- Comments are in English and follow Go best practices
- Core sorting logic is separated into the internal package
- Input handling and main logic in cmd package

### Naming Conventions
- Functions: PascalCase for exported functions (e.g., `SortByACount`)
- Variables: camelCase for internal variables (e.g., `wordCount`)
- Packages: lowercase, single-word names (e.g., `sorter`)

## Testing
Run all tests:
go test ./...

Run tests with coverage:
go test ./... -cover

## Git Workflow
1. Create a new branch for each feature:
git checkout -b feature/new-feature

2. Make changes and commit:
git add .
git commit -m "Descriptive commit message"

3. Push changes:
git push origin feature/new-feature

### Commit Message Format
- feat: New feature
- fix: Bug fix
- docs: Documentation changes
- test: Adding or modifying tests
- refactor: Code refactoring

## Contributing
1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## Error Handling
- Handles empty input gracefully
- Provides clear error messages
- Validates input format
- Proper exit codes for different scenarios

## Future Improvements
- Support for file input/output
- Case-insensitive sorting option
- Performance optimizations for large inputs
- Additional sorting criteria options