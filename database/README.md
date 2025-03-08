# Database Package

This package provides database interaction functionalities for the Dracory framework. It offers a set of tools for interacting with various database systems.

## Usage

This package provides functionalities for opening database connections, executing queries, inserting data, and more. It can be used to interact with various database systems.

## Example

- Example of opening a database connection

```go
db, err := database.Open(database.Options().
     SetDatabaseType(DbDriver).
     SetDatabaseHost(DbHost).
     SetDatabasePort(DbPort).
     SetDatabaseName(DbName).
     SetCharset(`utf8mb4`).
     SetUserName(DbUser).
     SetPassword(DbPass))

if err != nil {
     return err
}

if db == nil {
     return errors.New("db is nil")
}

defer db.Close()
```

- Example of executing a raw query

```go
// using DB
dbCtx := Context(context.Background(), db)
rows, err := Query(dbCtx, "SELECT * FROM users")
if err != nil {
     log.Fatalf("Failed to execute query: %v", err)
}
defer rows.Close()

// using transaction
txCtx := Context(context.Background(), tx)
rows, err := Query(txCtx, "SELECT * FROM users")
if err != nil {
     log.Fatalf("Failed to execute query: %v", err)
}
defer rows.Close()
```

- Example of inserting data

```go
// using DB
dbCtx := Context(context.Background(), db)
_, err := Execute(dbCtx, "INSERT INTO users (name, email) VALUES (?, ?)", "John Doe", "a4lGw@example.com")
if err != nil {
     log.Fatalf("Failed to insert data: %v", err)
}

// using transaction
txCtx := Context(context.Background(), tx)
_, err := Execute(txCtx, "INSERT INTO users (name, email) VALUES (?, ?)", "John Doe", "a4lGw@example.com")
if err != nil {
     log.Fatalf("Failed to insert data: %v", err)
}
```

- Select rows (as map[string]string)

```go
mappedRows, err := database.SelectToMapString(store.toQuerableContext(ctx), sqlStr, params...)
if err != nil {
     log.Fatalf("Failed to select rows: %v", err)
}
```

- Select rows (as map[string]any)

```go
mappedRows, err := database.SelectToMapAny(store.toQuerableContext(ctx), sqlStr, params...)
if err != nil {
     log.Fatalf("Failed to select rows: %v", err)
}
```
