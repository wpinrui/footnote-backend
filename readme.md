# footnote-backend

A backend service for the upcoming Footnote application, built with Go and PostgreSQL.

## To create new db migrations

```
migrate create -ext sql -dir migrations -seq <description>
```

Replace `<description>` with a meaningful name for the migration, such as `create_users_table`.

## Setup and Run Instructions

### 1. Copy `.env.example` to `.env`

### 2. Generate JWT_SECRET

Generate a JWT secret key using OpenSSL:

```bash
openssl rand -hex 32
```

Copy the output and set it as `JWT_SECRET` in your `.env` file. Note that the supplied example is only a placeholder and should be replaced.

### 3. Update your DSN

Update the `DSN` variable in your `.env` file with your PostgreSQL connection string. It should look like this:

```
DSN=postgres://username:password@localhost:5432/db_name?sslmode=disable
```

Replace `username`, `password`, and `db_name` with your PostgreSQL credentials and database name.

### 4. Load environment variables and run migration

Load the `.env` variables into your terminal session, then run the make command to apply migrations:

```bash
export $(grep -v '^#' .env | xargs)
make migrate-up
```

You can check that the tables have been created successfully by running:

```
psql "$DSN" -c "\dt"
```

### 5. Install swag globally (recommended)

```bash
go install github.com/swaggo/swag/cmd/swag@latest
# Add swag to path, then reload
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.zshrc
```

### 6. Run the backend

#### Option 1: Using VSCode launch config

Open VSCode

- Use the `.vscode/launch.json` configuration to start the server
- Note that `launch.json` will attempt to update swagger docs. If you chose not to setup swag, this step can fail but you can still continue debugging.

#### Option 2: Using Go CLI

```bash
# if making changes to swagger
swag init --generalInfo cmd/main.go
go run cmd/main.go
```
