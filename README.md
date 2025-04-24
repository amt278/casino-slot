# 🎰 Casino Slot Machine (Go)

A lightweight command-line slot machine game written in Go. Spin the reels and test your luck!

## Features

- **Simple Gameplay**: Engage with a classic slot machine experience directly from your terminal.
- **Randomized Outcomes**: Utilizes Go's `math/rand` package to generate random symbols for each spin.
- **Modular Codebase**: Organized into separate files for main execution (`main.go`), spinning logic (`spin.go`), and utility functions (`util.go`).

## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/amt278/casino-slot.git
   cd casino-slot
   ```

2. **Build the Application**:

   Ensure you have Go installed on your system. Then, build the application using:

   ```bash
   go build -o casino-slot
   ```

3. **Run the Game**:

   Execute the compiled binary:

   ```bash
   ./casino-slot
   ```

## Usage

Upon running the game, follow the on-screen prompts to spin the slot machine. The game will display the outcome of each spin, indicating whether you've won or should try again.

## Project Structure

- `main.go`: Initializes the game and handles user interaction.
- `spin.go`: Contains the logic for spinning the reels and determining outcomes.
- `util.go`: Provides utility functions to support the game's operations.

## Contributing

Contributions are welcome! If you'd like to enhance the game or fix issues, please fork the repository and submit a pull request.
