
# Hidden Duck Airdrop Bot

This is a **Telegram bot** built using **Golang** that automates the completion of hidden duck airdrop tasks by sending **HTTP requests**.

## Features

- Automates participation in hidden duck airdrop events.
- Sends HTTP requests to interact with airdrop platforms and complete tasks.
- Fully configurable for different task types.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Golang](https://golang.org/doc/install) (version 1.16 or above)
- A Telegram Bot API Token (created via [BotFather](https://core.telegram.org/bots#creating-a-new-bot) on Telegram)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/itpourya/DuckAirdropHiddenTask.git
   ```

2. Move into the project directory:

   ```bash
   cd DuckAirdropHiddenTask
   ```

3. Install the required dependencies:

   ```bash
   go mod tidy
   ```

## Configuration

1. Open the `cheatbot/env.go` file and set your **Telegram Bot Token** and other necessary configurations:

   ```bash
   TOKEN=your_telegram_bot_token
   ```

   Customize other environment variables as needed to suit the hidden duck airdrop task's requirements.

## Usage

To run the bot:

```bash
go run main.go
```

The bot will start interacting with the hidden duck airdrop platform and completing tasks automatically.

## Project Structure

- `main.go`: The entry point of the bot.
- `handler/`: Contains the logic to handle different Telegram bot events and HTTP requests.
- `process/`: Contains services responsible for interacting with the airdrop API.
- `env.go`: Environment variables like bot tokens and API URLs.

## Contribution

Feel free to contribute to this project by submitting pull requests:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -m "Add new feature"`).
4. Push the branch (`git push origin feature-branch`).
5. Open a Pull Request.
