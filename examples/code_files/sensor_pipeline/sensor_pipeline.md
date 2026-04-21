# Sensor Data Pipeline with Error Handling

A real-world IoT data pipeline that ingests raw sensor readings from multiple devices, handles malformed data gracefully using the `errors` package, deduplicates overlapping streams, and produces clean time-series output — all without crashing on bad data.

---

## What This Program Does

Imagine you have IoT sensors sending temperature readings. Some readings are valid, some have errors (sensor disconnected, value out of range, corrupt data). You receive data from **two overlapping sensor feeds** that need to be merged and deduplicated.

This program:

1. **Parses** raw sensor strings into structured readings, wrapping failures in `Result[T]`
2. **Validates** readings — rejects temperatures outside a sane range as errors
3. **Uses the error-aware pipeline** — errors flow through without crashing
4. **Collects** the clean data, reporting the first error if any exist
5. **Merges** two sorted, clean sensor feeds into one deduplicated stream
6. **Computes** statistics (min, max, average) using `Reduce`

---

## Packages Used

| Package | Purpose |
|---------|---------|
| `core` | Create iterators from slices |
| `ops` | Map, Filter, Reduce, MergeDistinct, Collect, Scan |
| `errors` | Result type, error-aware Map/Filter/Collect, Stream API |
| `compare` | Comparator for merging sorted readings |

---

## Code

```go
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ver1619/itFrame/compare"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/errors"
	"github.com/ver1619/itFrame/ops"
)

// ── Data Models ──

type SensorReading struct {
	Timestamp int     // unix seconds (simplified)
	DeviceID  string
	TempC     float64
}

// ── Parser ──

// parseReading turns "timestamp,device,temp" into a Result[SensorReading].
// Bad formats or missing fields become error results instead of panics.
func parseReading(line string) errors.Result[SensorReading] {
	parts := strings.Split(line, ",")
	if len(parts) != 3 {
		return errors.ErrResult[SensorReading](
			fmt.Errorf("malformed line (expected 3 fields): %q", line),
		)
	}

	ts, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return errors.ErrResult[SensorReading](
			fmt.Errorf("bad timestamp %q: %w", parts[0], err),
		)
	}

	temp, err := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)
	if err != nil {
		return errors.ErrResult[SensorReading](
			fmt.Errorf("bad temperature %q: %w", parts[2], err),
		)
	}

	return errors.Ok(SensorReading{
		Timestamp: ts,
		DeviceID:  strings.TrimSpace(parts[1]),
		TempC:     temp,
	})
}

// ── Validator ──

// validateRange rejects temperatures outside -50°C to 60°C as sensor errors.
func validateRange(r SensorReading) errors.Result[SensorReading] {
	if r.TempC < -50 || r.TempC > 60 {
		return errors.ErrResult[SensorReading](
			fmt.Errorf("out-of-range temp %.1f°C from %s at t=%d",
				r.TempC, r.DeviceID, r.Timestamp),
		)
	}
	return errors.Ok(r)
}

func main() {
	// ── Raw data from Sensor Feed A ──
	feedA := []string{
		"1000,sensor-A,22.5",
		"1001,sensor-A,22.8",
		"1002,sensor-A,corrupt_data", // ← bad temperature
		"1003,sensor-A,23.1",
		"1004,sensor-A,23.4",
		"1005,sensor-A,350.0", // ← out-of-range (will be caught by validation)
		"1006,sensor-A,23.9",
	}

	// ── Raw data from Sensor Feed B (overlapping timestamps) ──
	feedB := []string{
		"1002,sensor-B,21.0",
		"1003,sensor-B,21.3",
		"1004,sensor-B", // ← malformed (only 2 fields)
		"1005,sensor-B,21.8",
		"1006,sensor-B,22.1",
		"1007,sensor-B,22.4",
	}

	// ─────────────────────────────────────────────
	// Part 1: Error-Aware Pipeline on Feed A
	// ─────────────────────────────────────────────
	fmt.Println("=== Feed A: Error-Aware Processing ===\n")

	// Step 1: Parse each line into Result[SensorReading]
	parsedA := ops.Map(core.Slice(feedA), parseReading)

	// Step 2: For valid readings, run range validation
	// errors.Map only touches Ok values — errors pass through unchanged
	validatedA := errors.Map(parsedA, func(r SensorReading) SensorReading {
		// This runs only on successfully parsed readings
		return r
	})

	// Step 3: Apply range check using FlatMap-style approach
	// We need to turn Ok values into potentially new errors
	// Use a manual approach: map each Result to a re-validated Result
	revalidatedA := ops.Map(validatedA, func(r errors.Result[SensorReading]) errors.Result[SensorReading] {
		if r.IsError() {
			return r // pass errors through
		}
		return validateRange(r.Value) // might produce a new error
	})

	// Step 4: Collect — stops at first error
	valsA, err := errors.Collect(revalidatedA)
	if err != nil {
		fmt.Printf("  ⚠ Pipeline stopped: %s\n", err)
		fmt.Printf("  ✓ Collected %d valid readings before error\n\n", len(valsA))
	} else {
		fmt.Printf("  ✓ All %d readings valid\n\n", len(valsA))
	}

	// ─────────────────────────────────────────────
	// Part 2: Error-Aware Stream API on Feed B
	// ─────────────────────────────────────────────
	fmt.Println("=== Feed B: Stream API Processing ===\n")

	parsedBResults := make([]errors.Result[SensorReading], len(feedB))
	for i, line := range feedB {
		parsedBResults[i] = parseReading(line)
	}

	valsB, err := errors.FromSlice(parsedBResults).
		Filter(func(r SensorReading) bool {
			return r.TempC > 0 // only positive temps (errors pass through)
		}).
		Map(func(r SensorReading) SensorReading {
			r.TempC = r.TempC + 273.15 // convert to Kelvin
			return r
		}).
		Collect()

	if err != nil {
		fmt.Printf("  ⚠ Pipeline stopped: %s\n", err)
		fmt.Printf("  ✓ Collected %d readings before error\n\n", len(valsB))
	} else {
		fmt.Printf("  ✓ All %d readings processed (in Kelvin)\n", len(valsB))
		for _, v := range valsB {
			fmt.Printf("    t=%d %s: %.2f K\n", v.Timestamp, v.DeviceID, v.TempC)
		}
		fmt.Println()
	}

	// ─────────────────────────────────────────────
	// Part 3: Merge Clean Feeds + Statistics
	// ─────────────────────────────────────────────
	fmt.Println("=== Merged Clean Data + Statistics ===\n")

	// Get clean readings (skip errors this time by filtering manually)
	cleanA := []SensorReading{}
	for _, line := range feedA {
		r := parseReading(line)
		if !r.IsError() {
			v := validateRange(r.Value)
			if !v.IsError() {
				cleanA = append(cleanA, v.Value)
			}
		}
	}

	cleanB := []SensorReading{}
	for _, line := range feedB {
		r := parseReading(line)
		if !r.IsError() {
			cleanB = append(cleanB, r.Value)
		}
	}

	// Merge two sorted feeds by timestamp, removing duplicates at same timestamp
	byTimestamp := compare.LessFunc[SensorReading](func(a, b SensorReading) bool {
		return a.Timestamp < b.Timestamp
	})

	merged := ops.MergeDistinct(
		core.Slice(cleanA),
		core.Slice(cleanB),
		byTimestamp,
	)

	allClean := ops.Collect(merged)

	fmt.Printf("  Clean readings after merge: %d\n", len(allClean))
	for _, r := range allClean {
		fmt.Printf("    t=%d  %-10s  %.1f°C\n", r.Timestamp, r.DeviceID, r.TempC)
	}

	// ── Compute statistics with Reduce ──
	type Stats struct {
		Min, Max, Sum float64
		Count         int
	}

	stats := ops.Reduce(core.Slice(allClean), Stats{
		Min: 999, Max: -999,
	}, func(acc Stats, r SensorReading) Stats {
		if r.TempC < acc.Min {
			acc.Min = r.TempC
		}
		if r.TempC > acc.Max {
			acc.Max = r.TempC
		}
		acc.Sum += r.TempC
		acc.Count++
		return acc
	})

	fmt.Printf("\n  📊 Statistics:\n")
	fmt.Printf("    Min: %.1f°C\n", stats.Min)
	fmt.Printf("    Max: %.1f°C\n", stats.Max)
	fmt.Printf("    Avg: %.1f°C\n", stats.Sum/float64(stats.Count))
	fmt.Printf("    Count: %d\n", stats.Count)

	// ── Running average with Scan ──
	type RunningAvg struct {
		Sum   float64
		Count int
	}

	runningIt := ops.Scan(core.Slice(allClean), RunningAvg{},
		func(acc RunningAvg, r SensorReading) RunningAvg {
			acc.Sum += r.TempC
			acc.Count++
			return acc
		},
	)

	fmt.Println("\n  📈 Running Average:")
	runningAvgs := ops.Collect(runningIt)
	for i, ra := range runningAvgs {
		avg := ra.Sum / float64(ra.Count)
		fmt.Printf("    After reading %d: avg = %.2f°C\n", i+1, avg)
	}
}
```

---

## Output

```
=== Feed A: Error-Aware Processing ===

  Pipeline stopped: bad temperature "corrupt_data": strconv.ParseFloat: parsing "corrupt_data": invalid syntax
  Collected 0 valid readings before error

=== Feed B: Stream API Processing ===

  Pipeline stopped: malformed line (expected 3 fields): "1004,sensor-B"
  Collected 0 readings before error

=== Merged Clean Data + Statistics ===

  Clean readings after merge: 8
    t=1000  sensor-A    22.5C
    t=1001  sensor-A    22.8C
    t=1002  sensor-B    21.0C
    t=1003  sensor-A    23.1C
    t=1004  sensor-A    23.4C
    t=1005  sensor-B    21.8C
    t=1006  sensor-A    23.9C
    t=1007  sensor-B    22.4C

  Statistics:
    Min: 21.0C
    Max: 23.9C
    Avg: 22.6C
    Count: 8

  Running Average:
    After reading 1: avg = 22.50C
    After reading 2: avg = 22.65C
    After reading 3: avg = 22.10C
    After reading 4: avg = 22.35C
    After reading 5: avg = 22.56C
    After reading 6: avg = 22.43C
    After reading 7: avg = 22.64C
    After reading 8: avg = 22.61C
```

---

## How the Pipeline Flows

```
Raw CSV strings (Feed A)              Raw CSV strings (Feed B)
        │                                       │
   parseReading()                          parseReading()
        │                                       │
  Result[SensorReading]                   Result[SensorReading]
        │                                       │
  errors.Map (transform)               errors.Stream.Filter
        │                                       │
  validateRange (→ new errors)          errors.Stream.Map
        │                                       │
  errors.Collect                         errors.Stream.Collect
  (stops at first error)                 (stops at first error)
        │                                       │
        └────── Clean readings ─────────────────┘
                       │
              ops.MergeDistinct (by timestamp)
                       │
              ops.Reduce (min/max/avg)
                       │
              ops.Scan (running average)
```

---

## Key Takeaways

- **`errors.Result[T]`** replaces `(value, error)` tuples — one type that carries either
- **Error propagation is automatic** — `errors.Map` and `errors.Filter` skip error values without you writing `if err != nil` checks
- **`errors.Collect` is fail-fast** — returns all good values up to the first error
- **Two API styles** — use standalone `errors.Map()` functions or the fluent `errors.FromSlice().Map().Collect()` stream
- **Clean data after filtering** — once errors are handled, use regular `ops` for merging and statistics
- **`MergeDistinct`** — combines overlapping sorted feeds, dropping duplicates at the same timestamp
- **`Reduce` for aggregation** — computes min, max, sum in a single pass
- **`Scan` for time-series** — shows how the average evolves as readings arrive
