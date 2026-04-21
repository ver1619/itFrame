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

type SensorReading struct {
	Timestamp int
	DeviceID  string
	TempC     float64
}

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
	feedA := []string{
		"1000,sensor-A,22.5",
		"1001,sensor-A,22.8",
		"1002,sensor-A,corrupt_data",
		"1003,sensor-A,23.1",
		"1004,sensor-A,23.4",
		"1005,sensor-A,350.0",
		"1006,sensor-A,23.9",
	}

	feedB := []string{
		"1002,sensor-B,21.0",
		"1003,sensor-B,21.3",
		"1004,sensor-B",
		"1005,sensor-B,21.8",
		"1006,sensor-B,22.1",
		"1007,sensor-B,22.4",
	}

	// Part 1: Error-Aware Pipeline on Feed A
	fmt.Println("=== Feed A: Error-Aware Processing ===")
	fmt.Println()

	parsedA := ops.Map(core.Slice(feedA), parseReading)

	validatedA := errors.Map(parsedA, func(r SensorReading) SensorReading {
		return r
	})

	revalidatedA := ops.Map(validatedA, func(r errors.Result[SensorReading]) errors.Result[SensorReading] {
		if r.IsError() {
			return r
		}
		return validateRange(r.Value)
	})

	valsA, err := errors.Collect(revalidatedA)
	if err != nil {
		fmt.Printf("  Pipeline stopped: %s\n", err)
		fmt.Printf("  Collected %d valid readings before error\n\n", len(valsA))
	} else {
		fmt.Printf("  All %d readings valid\n\n", len(valsA))
	}

	// Part 2: Error-Aware Stream API on Feed B
	fmt.Println("=== Feed B: Stream API Processing ===")
	fmt.Println()

	parsedBResults := make([]errors.Result[SensorReading], len(feedB))
	for i, line := range feedB {
		parsedBResults[i] = parseReading(line)
	}

	valsB, err := errors.FromSlice(parsedBResults).
		Filter(func(r SensorReading) bool {
			return r.TempC > 0
		}).
		Map(func(r SensorReading) SensorReading {
			r.TempC = r.TempC + 273.15
			return r
		}).
		Collect()

	if err != nil {
		fmt.Printf("  Pipeline stopped: %s\n", err)
		fmt.Printf("  Collected %d readings before error\n\n", len(valsB))
	} else {
		fmt.Printf("  All %d readings processed (in Kelvin)\n", len(valsB))
		for _, v := range valsB {
			fmt.Printf("    t=%d %s: %.2f K\n", v.Timestamp, v.DeviceID, v.TempC)
		}
		fmt.Println()
	}

	// Part 3: Merge Clean Feeds + Statistics
	fmt.Println("=== Merged Clean Data + Statistics ===")
	fmt.Println()

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
		fmt.Printf("    t=%d  %-10s  %.1fC\n", r.Timestamp, r.DeviceID, r.TempC)
	}

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

	fmt.Printf("\n  Statistics:\n")
	fmt.Printf("    Min: %.1fC\n", stats.Min)
	fmt.Printf("    Max: %.1fC\n", stats.Max)
	fmt.Printf("    Avg: %.1fC\n", stats.Sum/float64(stats.Count))
	fmt.Printf("    Count: %d\n", stats.Count)

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

	fmt.Println("\n  Running Average:")
	runningAvgs := ops.Collect(runningIt)
	for i, ra := range runningAvgs {
		avg := ra.Sum / float64(ra.Count)
		fmt.Printf("    After reading %d: avg = %.2fC\n", i+1, avg)
	}
}
