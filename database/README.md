# Database Package

This package provides database interaction functionalities for the Dracory
framework. It offers a set of tools for interacting with various database systems.

## Usage

This package provides functionalities for opening database connections,
executing queries, inserting data, managing transactions, and more.
It can be used to interact with various database systems.

## Core Functions

### Context and ContextOr

The database package provides two key functions for working with database contexts:

- **Context**: Creates a new queryable context with the given database connection or transaction
  ```go
  // Create a queryable context with a database connection
  dbCtx := database.Context(context.Background(), db)
  
  // Create a queryable context with a transaction
  txCtx := database.Context(context.Background(), tx)
  ```

- **ContextOr**: Intelligently handles contexts that may or may not already be queryable contexts
  ```go
  // If ctx is already a queryable context, it will be returned as is
  // If not, a new queryable context with db will be created
  qCtx := database.ContextOr(ctx, db)
  ```

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

- Example of inserting data with DB connection

```go
dbCtx := Context(context.Background(), db)
result, err := Execute(dbCtx, "INSERT INTO users (name, email) VALUES (?, ?)", "John Doe", "john@example.com")
if err != nil {
     log.Fatalf("Failed to insert data: %v", err)
}
```

- Example of inserting data with transaction

```go
// Begin a transaction
tx, err := db.BeginTx(context.Background(), nil)
if err != nil {
     log.Fatalf("Failed to begin transaction: %v", err)
}

// Create a queryable context with the transaction
txCtx := Context(context.Background(), tx)

// Execute the query within the transaction
result, err := Execute(txCtx, "INSERT INTO users (name, email) VALUES (?, ?)", "Jane Doe", "jane@example.com")
if err != nil {
     // With transactions, you typically want to roll back on error
     tx.Rollback()
     log.Fatalf("Failed to insert data: %v", err)
}

// If successful, commit the transaction
err = tx.Commit()
if err != nil {
     tx.Rollback()
     log.Fatalf("Failed to commit transaction: %v", err)
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

## Transactions

The database package supports transactions through the standard Go `database/sql` package.
The `QueryableInterface` can work with `*sql.DB`, `*sql.Conn`, or `*sql.Tx` (transaction)
objects through a context-based approach.

### Starting a Transaction

```go
// Get a database connection
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
defer db.Close()

// Begin a transaction
tx, err := db.BeginTx(context.Background(), nil)
if err != nil {
     return err
}

// Create a queryable context with the transaction
txCtx := database.Context(context.Background(), tx)
```

### Using Transactions

Once you have a transaction context, you can use it with any of the database functions:

```go
// Execute a query within the transaction
result, err := database.Execute(txCtx, "INSERT INTO users (name, email) VALUES (?, ?)", "John Doe", "john@example.com")
if err != nil {
     // Roll back the transaction if there's an error
     tx.Rollback()
     return err
}

// Query data within the transaction
rows, err := database.Query(txCtx, "SELECT * FROM users WHERE name = ?", "John Doe")
if err != nil {
     tx.Rollback()
     return err
}
defer rows.Close()

// Process the query results...
```

### Committing or Rolling Back Transactions

```go
// If all operations are successful, commit the transaction
err = tx.Commit()
if err != nil {
     // If commit fails, try to roll back
     tx.Rollback()
     return err
}

// If an error occurs during any operation, roll back the transaction
// tx.Rollback()
```

### Using ContextOr with Transactions

The `ContextOr` function provides a convenient way to work with contexts that
may or may not already be queryable contexts. This is especially useful in
functions that might receive either a regular context or a transaction context:

```go
// Function that can work with either a regular context or a transaction context
func GetUserByID(ctx context.Context, db *sql.DB, userID int) (map[string]any, error) {
    // Convert the context to a queryable context if it isn't already one
    qCtx := database.ContextOr(ctx, db)
    
    // If ctx was already a transaction context, it will be used as is
    // If not, a new queryable context with db will be created
    return database.SelectToMapAny(qCtx, "SELECT * FROM users WHERE id = ?", userID)
}
```

This allows you to write functions that can participate in larger transactions
when needed, but can also work independently with a direct database connection.

#### Using ContextOr in Data Stores

In the Dracory framework, data stores are kept in independent packages with
public interfaces and private implementations. The `ContextOr` function is
particularly useful in these store implementations:

```go
// UserStore interface in the users package
type UserStoreInterface interface {
    FindByID(ctx context.Context, id int) (*User, error)
    Create(ctx context.Context, user *User) error
    // Other methods...
}

// Private implementation
type userStore struct {
    db *sql.DB
}

// Implementation using ContextOr to support both regular and transaction contexts
func (store *userStore) FindByID(ctx context.Context, id int) (*User, error) {
    // Convert to queryable context, preserving any existing transaction
    qCtx := database.ContextOr(ctx, store.db)
    
    // Use the queryable context for database operations
    rows, err := database.Query(qCtx, "SELECT * FROM users WHERE id = ?", id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    // Process results...
}
```

This pattern allows store methods to be called either with a regular context for independent operations or with a transaction context when multiple operations need to be atomic.

### Transaction Best Practices

1. **Error Handling**: Always check for errors after each database operation
and roll back the transaction if an error occurs.

2. **Defer Rollback**: Consider using a deferred rollback that is ignored
if the transaction is committed successfully:

   ```go
   tx, err := db.BeginTx(context.Background(), nil)
   if err != nil {
        return err
   }
   
   // Defer a rollback in case anything fails
   defer func() {
        // The rollback will be ignored if the tx has been committed
        tx.Rollback()
   }()
   
   // Perform transaction operations...
   
   // If successful, commit
   return tx.Commit()
   ```

3. **Transaction Isolation**: Be aware of the default transaction isolation
level of your database. You can specify a different isolation level when
beginning a transaction:

   ```go
   tx, err := db.BeginTx(context.Background(), &sql.TxOptions{
        Isolation: sql.LevelSerializable,
        ReadOnly: false,
   })
   ```

4. **Keep Transactions Short**: Long-running transactions can cause performance
issues and deadlocks. Keep transactions as short as possible.
