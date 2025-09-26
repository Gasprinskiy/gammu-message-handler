# Gammu Message Handler

A Go application that receives SMS messages from a GSM modem via [Gammu SMSD](https://wammu.eu/gammu/), stores them in a database, and forwards them to a Telegram chat.

This repository contains only the **message handler** service.
For a full setup you will also need:

- 📦 [gammu-mysql-db](https://github.com/Gasprinskiy/gammu-mysql-db) — database schema and migrations
- 📦 [gammu-smsd-container](https://github.com/Gasprinskiy/gammu-smsd-container) — containerized Gammu SMSD service

Each repository contains its own **instructions for setup and running**, so please refer to them individually.

---

## Features

- Integration with **Gammu SMSD** for handling incoming SMS
- Stores messages in a relational database (MySQL/MariaDB)
- Forwards SMS to a **Telegram chat or channel**
- Lightweight, written in **Go**

---

## Prerequisites

- Docker and Docker Compose installed
- A GSM modem with a working SIM card
- Cloned repositories:
  - [gammu-mysql-db](https://github.com/Gasprinskiy/gammu-mysql-db)
  - [gammu-smsd-container](https://github.com/Gasprinskiy/gammu-smsd-container)
  - [gammu-message-handler](https://github.com/Gasprinskiy/gammu-message-handler)

---

## Setup & Run


1. **Create a common Docker network** for all SMS services, you can find a script in gammu-mysql-db repository or do it manually:

   ```bash
   docker network create sms_services
   ```

2. **Start the database service** (from the `gammu-mysql-db` repository):
   Instructions are available in that repository.
   The DB container should also join the `sms_services` network.

3. **Start the SMSD service** (from the `gammu-smsd-container` repository):
   Instructions are available in that repository.
   This will run Gammu SMSD inside a container and connect it to the `sms_services` network.
   If you done it correctly you will have all incoming messages in `inbox` table.

4. **Create a `.env` file** in the root of this repository with the following variables:

   ```env
   DB_HOST=gammu-mysql-db // it's a name of db service on docker compose file
   DB_PORT=3306
   DB_USER=yourdbuser
   DB_PASSWORD=yourpassword
   DB_NAME=yourdbname

   TELEGRAM_BOT_TOKEN=your_bot_token
   TELEGRAM_CHAT_ID=your_chat_id
````

5. **Finally, start the message handler** (from this repository):

   ```bash
   ./start_docker.sh
   ```

   It will automatically connect to the database and forward incoming SMS messages to Telegram.

---

## Usage Flow

1. An SMS arrives at the GSM modem.
2. Gammu SMSD writes the message into the database and runs script that makes http request to this service.
3. The handler reads new messages from the DB.
4. The message is forwarded to the configured Telegram chat.