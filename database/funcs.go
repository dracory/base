package database

import "context"

// IsQueryableContext checks if the given context is a QueryableContext.
//
// Parameters:
// - ctx: The context to check.
//
// Returns:
// - bool: True if the context is a QueryableContext, false otherwise.
func IsQueryableContext(ctx context.Context) bool {
	if _, ok := ctx.(QueryableContext); ok {
		return true
	}

	return false
}

// Context returns a new context with the given QueryableInterface.
// It is a shortcut for NewQueryableContext.
//
// Example:
// 	ctx := database.Context(context.Background(), tx)
//
// Parameters:
// - ctx: The parent context.
// - queryable: The QueryableInterface to be associated with the context.
//
// Returns:
// - QueryableContext: A new context with the given QueryableInterface.
func Context(ctx context.Context, queryable QueryableInterface) QueryableContext {
	return NewQueryableContext(ctx, queryable)
}

// ContextOr returns the existing QueryableContext if the provided context
// is already a QueryableContext, or creates a new one with the given QueryableInterface.
// This is useful when you want to accept either a regular context or a queryable context.
//
// Example:
// 	// This will use the existing QueryableContext if ctx is already one,
// 	// or create a new one with db if it's not
// 	qCtx := database.ContextOr(ctx, db)
//
// Parameters:
// - ctx: The parent context, which may or may not be a QueryableContext.
// - queryable: The QueryableInterface to be associated with the context if a new one is created.
//
// Returns:
// - QueryableContext: Either the existing QueryableContext or a new one.
func ContextOr(ctx context.Context, queryable QueryableInterface) QueryableContext {
	if IsQueryableContext(ctx) {
		return ctx.(QueryableContext)
	}

	return NewQueryableContext(ctx, queryable)
}
