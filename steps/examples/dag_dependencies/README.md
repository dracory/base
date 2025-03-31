# Dependencies Example

This example demonstrates how to create steps with complex dependencies and ensure they execute in the correct order, using a real-world scenario of calculating a final price.

## Overview

The example shows:
1. How to create multiple steps with different operations
2. How to set up complex dependencies between steps
3. How the steps are automatically sorted and executed in dependency order

## Key Concepts

- Step Dependencies: Using `AddDependency` to specify that one step depends on another
- Topological Sorting: The steps are automatically sorted to ensure dependencies are respected
- Execution Order: Steps are executed in the order determined by their dependencies
- Complex Dependencies: Steps can depend on multiple other steps

## Running the Example

To run this example:

```bash
# Run the main program
go run main.go

# Run the tests
go test -v
```

## Expected Output

The program will output:
```
Base price: 100
Final price: 108
Steps completed: [SetBasePrice ApplyDiscount AddShipping CalculateTax]
```

## Price Calculation Process

The steps perform the following operations in order:
1. `SetBasePrice`: Sets the base price to 100
2. `ApplyDiscount`: Applies a 20% discount (result: 80)
3. `AddShipping`: Adds a fixed shipping cost of 10 (result: 90)
4. `CalculateTax`: Adds 20% tax (result: 108)

The final price is calculated as:
1. Start with base price: 100
2. Apply 20% discount: 100 * 0.8 = 80
3. Add shipping: 80 + 10 = 90
4. Add 20% tax: 90 * 1.2 = 108
