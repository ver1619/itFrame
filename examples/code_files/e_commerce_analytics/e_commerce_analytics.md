# E-Commerce Order Analytics

A data analytics pipeline that processes customer orders and product catalogs — joining datasets, computing revenue per category, paginating results, and running validation checks. This demonstrates how itFrame can replace what you'd normally need SQL or a DataFrame library for.

---

## What This Program Does

You have two datasets:

- **Products** — catalog with ID, name, category, and price
- **OrderItems** — individual line items with product ID and quantity

This program:

1. **Joins** order items with product details using an inner join
2. **Computes** revenue (price × quantity) for each line item
3. **Groups** by product category and **aggregates** total revenue per category
4. **Finds** the top 3 categories by revenue
5. **Validates** data with `Any` and `All` checks
6. **Paginates** results using `Skip` + `Take`

---

## Packages Used

| Package | Purpose |
|---------|---------|
| `core` | Create iterators from slices and ranges |
| `ops` | Join, Map, GroupBy, Aggregate, Skip, Take, Any, All, Reduce |
| `stream` | Fluent API for the validation and pagination sections |

---

## Code

```go
package main

import (
	"fmt"

	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
	"github.com/ver1619/itFrame/stream"
)

// ── Data Models ──

type Product struct {
	ID       int
	Name     string
	Category string
	Price    float64
}

type OrderItem struct {
	OrderID   int
	ProductID int
	Qty       int
}

type LineDetail struct {
	OrderID  int
	Product  string
	Category string
	Qty      int
	Revenue  float64
}

type CategoryReport struct {
	Category     string
	TotalRevenue float64
	ItemsSold    int
}

func main() {
	// ── Dataset: Product Catalog ──
	products := []Product{
		{1, "Mechanical Keyboard", "Electronics", 89.99},
		{2, "USB-C Hub", "Electronics", 34.99},
		{3, "Desk Lamp", "Home", 24.99},
		{4, "Ergonomic Chair", "Furniture", 399.99},
		{5, "Monitor Stand", "Furniture", 49.99},
		{6, "Notebook Pack", "Stationery", 12.99},
		{7, "Wireless Mouse", "Electronics", 29.99},
		{8, "Standing Desk Mat", "Furniture", 59.99},
		{9, "Cable Organizer", "Home", 14.99},
		{10, "Whiteboard Markers", "Stationery", 8.99},
	}

	// ── Dataset: Order Items (line items from various orders) ──
	orderItems := []OrderItem{
		{1001, 1, 2}, {1001, 3, 1}, {1001, 6, 5},
		{1002, 4, 1}, {1002, 2, 3}, {1002, 7, 2},
		{1003, 5, 1}, {1003, 8, 1}, {1003, 9, 4},
		{1004, 1, 1}, {1004, 10, 10}, {1004, 4, 1},
		{1005, 7, 3}, {1005, 2, 2}, {1005, 3, 2},
	}

	// ─────────────────────────────────────────────
	// Step 1: Join order items with product details
	// ─────────────────────────────────────────────
	// Inner join on ProductID — each order item gets paired with its product info
	joined := ops.Join(
		core.Slice(orderItems),
		core.Slice(products),
		func(oi OrderItem) int { return oi.ProductID },
		func(p Product) int { return p.ID },
	)

	// ─────────────────────────────────────────────
	// Step 2: Compute revenue for each line item
	// ─────────────────────────────────────────────
	// Map each joined pair into a LineDetail with calculated revenue
	details := ops.Map(joined, func(p ops.Pair[OrderItem, Product]) LineDetail {
		return LineDetail{
			OrderID:  p.First.OrderID,
			Product:  p.Second.Name,
			Category: p.Second.Category,
			Qty:      p.First.Qty,
			Revenue:  float64(p.First.Qty) * p.Second.Price,
		}
	})

	// Collect so we can reuse the data (iterators are single-pass)
	allDetails := ops.Collect(details)

	fmt.Println("=== All Line Items ===")
	for _, d := range allDetails {
		fmt.Printf("  Order #%d | %-22s | %dx | $%.2f\n",
			d.OrderID, d.Product, d.Qty, d.Revenue)
	}

	// ─────────────────────────────────────────────
	// Step 3: Group by category and aggregate
	// ─────────────────────────────────────────────
	groups := ops.GroupBy(
		core.Slice(allDetails),
		func(d LineDetail) string { return d.Category },
	)

	reports := ops.Aggregate(groups, func(g ops.Group[string, LineDetail]) CategoryReport {
		totalRev := 0.0
		totalQty := 0
		for _, item := range g.Items {
			totalRev += item.Revenue
			totalQty += item.Qty
		}
		return CategoryReport{
			Category:     g.Key,
			TotalRevenue: totalRev,
			ItemsSold:    totalQty,
		}
	})

	allReports := ops.Collect(reports)

	fmt.Println("\n=== Revenue by Category ===")
	for _, r := range allReports {
		fmt.Printf("  %-15s $%8.2f  (%d items)\n", r.Category, r.TotalRevenue, r.ItemsSold)
	}

	// ─────────────────────────────────────────────
	// Step 4: Total revenue using Reduce
	// ─────────────────────────────────────────────
	totalRevenue := ops.Reduce(
		core.Slice(allDetails),
		0.0,
		func(acc float64, d LineDetail) float64 { return acc + d.Revenue },
	)
	fmt.Printf("\n=== Total Revenue: $%.2f ===\n", totalRevenue)

	// ─────────────────────────────────────────────
	// Step 5: Pagination with Stream API
	// ─────────────────────────────────────────────
	pageSize := 5
	page := 2 // page 2 (0-indexed would be items 5-9)

	fmt.Printf("\n=== Line Items Page %d (size %d) ===\n", page, pageSize)
	pageItems := stream.Slice(allDetails).
		Skip((page - 1) * pageSize).
		Take(pageSize).
		Collect()

	for _, d := range pageItems {
		fmt.Printf("  Order #%d | %-22s | $%.2f\n", d.OrderID, d.Product, d.Revenue)
	}

	// ─────────────────────────────────────────────
	// Step 6: Validation checks with Any / All
	// ─────────────────────────────────────────────
	fmt.Println("\n=== Validation ===")

	hasLargeOrder := stream.Slice(allDetails).
		Any(func(d LineDetail) bool { return d.Revenue > 300 })
	fmt.Printf("  Has order over $300: %v\n", hasLargeOrder)

	allPositive := stream.Slice(allDetails).
		All(func(d LineDetail) bool { return d.Revenue > 0 })
	fmt.Printf("  All revenues positive: %v\n", allPositive)

	// Count electronics items
	electronicsCount := stream.Slice(allDetails).
		Filter(func(d LineDetail) bool { return d.Category == "Electronics" }).
		Count()
	fmt.Printf("  Electronics line items: %d\n", electronicsCount)

	// Sum of furniture revenue using stream
	furnitureRevenue := stream.Slice(allDetails).
		Filter(func(d LineDetail) bool { return d.Category == "Furniture" }).
		Reduce(LineDetail{}, func(acc, val LineDetail) LineDetail {
			acc.Revenue += val.Revenue
			return acc
		})
	fmt.Printf("  Furniture revenue: $%.2f\n", furnitureRevenue.Revenue)
}
```

---

## Output

```
=== All Line Items ===
  Order #1001 | Mechanical Keyboard    | 2x | $179.98
  Order #1001 | Desk Lamp              | 1x | $24.99
  Order #1001 | Notebook Pack          | 5x | $64.95
  Order #1002 | Ergonomic Chair        | 1x | $399.99
  Order #1002 | USB-C Hub              | 3x | $104.97
  Order #1002 | Wireless Mouse         | 2x | $59.98
  Order #1003 | Monitor Stand          | 1x | $49.99
  Order #1003 | Standing Desk Mat      | 1x | $59.99
  Order #1003 | Cable Organizer        | 4x | $59.96
  Order #1004 | Mechanical Keyboard    | 1x | $89.99
  Order #1004 | Whiteboard Markers     | 10x | $89.90
  Order #1004 | Ergonomic Chair        | 1x | $399.99
  Order #1005 | Wireless Mouse         | 3x | $89.97
  Order #1005 | USB-C Hub              | 2x | $69.98
  Order #1005 | Desk Lamp              | 2x | $49.98

=== Revenue by Category ===
  Electronics     $  594.87  (13 items)
  Home            $  134.93  (7 items)
  Stationery      $  154.85  (15 items)
  Furniture       $  909.96  (4 items)

=== Total Revenue: $1794.61 ===

=== Line Items Page 2 (size 5) ===
  Order #1002 | Wireless Mouse         | $59.98
  Order #1003 | Monitor Stand          | $49.99
  Order #1003 | Standing Desk Mat      | $59.99
  Order #1003 | Cable Organizer        | $59.96
  Order #1004 | Mechanical Keyboard    | $89.99

=== Validation ===
  Has order over $300: true
  All revenues positive: true
  Electronics line items: 6
  Furniture revenue: $909.96
```

---

## How the Pipeline Flows

```
OrderItems []              Products []
     │                          │
     └──────── ops.Join ────────┘
                  │
                  ▼
         ops.Map (compute revenue)
                  │
          ┌───────┼───────────────┐
          ▼       ▼               ▼
     GroupBy   Reduce         stream.Slice
       │         │               │
  Aggregate   Total $      Skip → Take → Collect
       │                    (pagination)
    Category
    Reports
```

---

## Key Takeaways

- **`Join` replaces SQL joins** — link order items to products by ID in a single call
- **`Map` for derived fields** — calculate revenue from raw data without mutating the source
- **`GroupBy` + `Aggregate`** — the SQL `GROUP BY ... SUM()` pattern, done in-memory
- **`Reduce` for totals** — fold all data into a single number
- **`stream` for readability** — the validation section uses the fluent API for clean one-liners
- **`Skip` + `Take` for pagination** — a pattern that works for any dataset, not just databases
- **Single-pass rule** — we `Collect` after joining so we can iterate the results multiple times
