# Gammu Message Handler

A Go application that receives SMS messages from a GSM modem via [Gammu SMSD](https://wammu.eu/gammu/), stores them in a database, and forwards them to a Telegram chat.

This repository contains only the **message handler** service.
For a full setup you will also need:

- 📦 [gammu-mysql-db](https://github.com/Gasprinskiy/gammu-mysql-db) — database schema and migrations
- 📦 [sms-smsd](https://github.com/Gasprinskiy/sms-smsd) — containerized Gammu SMSD service

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
  - [sms-smsd](https://github.com/Gasprinskiy/sms-smsd)
  - [gammu-message-handler](https://github.com/Gasprinskiy/gammu-message-handler)

---

## Setup & Run

1. **Create a `.env` file** in the root of this repository with the following variables:

   ```env
   DB_HOST=sms-db
   DB_PORT=3306
   DB_USER=smsuser
   DB_PASSWORD=yourpassword
   DB_NAME=smsdb

   TELEGRAM_BOT_TOKEN=your_bot_token
   TELEGRAM_CHAT_ID=your_chat_id
````

2. **Create a common Docker network** for all SMS services:

   ```bash
   docker network create sms_services
   ```

3. **Start the SMSD service** (from the `sms-smsd` repository):

   ```bash
   ./start_docker.sh
   ```

   This will run Gammu SMSD inside a container and connect it to the `sms_services` network.

4. **Start the database service** (from the `sms-database` repository):

   ```bash
   ./start_docker.sh
   ```

   The DB container should also join the `sms_services` network.

5. **Finally, start the message handler** (from this repository):

   ```bash
   ./start_docker.sh
   ```

   It will automatically connect to the database and forward incoming SMS messages to Telegram.

---

## Usage Flow

1. An SMS arrives at the GSM modem.
2. Gammu SMSD writes the message into the database.
3. The handler reads new messages from the DB.
4. The message is forwarded to the configured Telegram chat.

---
