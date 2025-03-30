# Conditional Logic Example

This example demonstrates how to use conditional logic to skip steps based on the context. It shows how to create different step chains based on the type of order being processed.

## Overview

The example shows:
1. How to create steps with dependencies
2. How to conditionally skip steps based on context
3. How to create different step chains for different scenarios
4. How to maintain proper step ordering

## Key Concepts

- Conditional Skipping: Steps can be skipped based on context conditions
- Step Dependencies: Maintaining proper dependencies between steps
- Different Step Chains: Creating different step chains for different scenarios

## Running the Example

To run this example:

```bash
# Run the tests
go test -v

# Run the main program
go run main.go
```

## Example Scenarios

The example demonstrates three different order types:

1. Digital Order:
   - Processes the order
   - Applies a discount
   - Calculates tax (no shipping needed)

2. Physical Order:
   - Processes the order
   - Applies a discount
   - Adds shipping cost
   - Calculates tax

3. Subscription Order:
   - Processes the order
   - Applies a discount (no shipping or tax needed)
