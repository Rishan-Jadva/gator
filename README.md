# Gator CLI

Gator is a command-line interface (CLI) application built with Go that interacts with a PostgreSQL database to manage users, feeds, and aggregation of RSS content. It provides a lightweight tool for scraping and processing RSS feeds in a structured and user-friendly way.


## 🚀 Requirements

To run and build the Gator CLI, you’ll need the following installed on your system:

- **Go** (version 1.20 or later recommended) – [Install Go](https://golang.org/doc/install)
- **PostgreSQL** – [Install PostgreSQL](https://www.postgresql.org/download/)



## 🔧 Installation

You can install the `gator` CLI globally using `go install`.



## ⚙️ Configuration

The CLI requires a configuration file to connect to your database and identify the current user. Create a `.gatorconfig` file in the root directory or your home directory with the following structure:

```json
DB_URL=postgres://username:password@localhost:5432/gator
CURRENT_USER=username123
```

Replace the values as appropriate for your setup.



## 📚 Available Commands

Here are a few example commands you can run after building or installing the CLI:

```bash
    gator register <user-name> – Register a new user

    gator login <user-name> – Log in as an existing user

    gator addfeed <feed-name> <feed-url> – Add an RSS feed to follow

    gator feeds – List available feeds

    gator agg – Run the feed aggregation process (scrape the next feed and display item titles)
```