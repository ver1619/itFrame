package main

import (
	"fmt"
	"strings"

	"github.com/ver1619/itFrame/compare"
	"github.com/ver1619/itFrame/core"
	"github.com/ver1619/itFrame/ops"
)

// LogEntry is a parsed log line.
type LogEntry struct {
	Timestamp string
	Level     string
	Endpoint  string
	Message   string
}

// EndpointSummary holds aggregated info about errors on one endpoint.
type EndpointSummary struct {
	Endpoint    string
	ErrorCount  int
	LastMessage string
}

// parseLine turns a raw log string into a LogEntry.
// Format: "TIMESTAMP LEVEL ENDPOINT MESSAGE..."
func parseLine(line string) LogEntry {
	parts := strings.SplitN(line, " ", 4)
	if len(parts) < 4 {
		return LogEntry{Message: line} // malformed line
	}
	return LogEntry{
		Timestamp: parts[0],
		Level:     parts[1],
		Endpoint:  parts[2],
		Message:   parts[3],
	}
}

func main() {
	// ── Raw logs from Server A ──
	serverALogs := []string{
		"2026-04-21T10:05:30 INFO /api/health ok",
		"2026-04-21T10:05:31 ERROR /api/users connection timeout",
		"2026-04-21T10:05:32 WARN /api/orders slow query: 2300ms",
		"2026-04-21T10:05:33 ERROR /api/users database unreachable",
		"2026-04-21T10:05:34 INFO /api/health ok",
		"2026-04-21T10:05:35 ERROR /api/payments gateway rejected",
		"2026-04-21T10:05:36 WARN /api/users rate limit approaching",
		"2026-04-21T10:05:37 ERROR /api/orders inventory lock failed",
	}

	// ── Raw logs from Server B ──
	serverBLogs := []string{
		"2026-04-21T10:05:30 ERROR /api/users auth token expired",
		"2026-04-21T10:05:31 INFO /api/health ok",
		"2026-04-21T10:05:32 ERROR /api/payments duplicate charge",
		"2026-04-21T10:05:33 WARN /api/orders high latency: 1800ms",
		"2026-04-21T10:05:34 ERROR /api/users session not found",
	}

	// ── Step 1: Parse raw lines into LogEntry structs ──
	parsedA := ops.Map(core.Slice(serverALogs), parseLine)
	parsedB := ops.Map(core.Slice(serverBLogs), parseLine)

	// ── Step 2: Keep only WARN and ERROR entries ──
	severeA := ops.Filter(parsedA, func(e LogEntry) bool {
		return e.Level == "ERROR" || e.Level == "WARN"
	})
	severeB := ops.Filter(parsedB, func(e LogEntry) bool {
		return e.Level == "ERROR" || e.Level == "WARN"
	})

	// ── Step 3: Group by endpoint ──
	groupsA := ops.GroupBy(severeA, func(e LogEntry) string { return e.Endpoint })
	groupsB := ops.GroupBy(severeB, func(e LogEntry) string { return e.Endpoint })

	// ── Step 4: Aggregate each group into a summary ──
	summarize := func(g ops.Group[string, LogEntry]) EndpointSummary {
		last := g.Items[len(g.Items)-1]
		return EndpointSummary{
			Endpoint:    g.Key,
			ErrorCount:  len(g.Items),
			LastMessage: last.Message,
		}
	}

	summariesA := ops.Aggregate(groupsA, summarize)
	summariesB := ops.Aggregate(groupsB, summarize)

	// ── Step 5: Collect both and print ──
	resultsA := ops.Collect(summariesA)
	resultsB := ops.Collect(summariesB)

	fmt.Println("=== Server A Error Summary ===")
	for _, s := range resultsA {
		fmt.Printf("  %-20s %d issue(s)  last: %s\n", s.Endpoint, s.ErrorCount, s.LastMessage)
	}

	fmt.Println("\n=== Server B Error Summary ===")
	for _, s := range resultsB {
		fmt.Printf("  %-20s %d issue(s)  last: %s\n", s.Endpoint, s.ErrorCount, s.LastMessage)
	}

	// ── Bonus: Merge both summaries sorted by error count (descending) ──
	bySeverity := compare.LessFunc[EndpointSummary](func(a, b EndpointSummary) bool {
		return a.ErrorCount > b.ErrorCount
	})

	mergedIt := ops.Merge(core.Slice(resultsA), core.Slice(resultsB), bySeverity)
	merged := ops.Collect(mergedIt)

	fmt.Println("\n=== Combined Priority Report (most issues first) ===")
	for i, s := range merged {
		fmt.Printf("  %d. %-20s %d issue(s)  last: %s\n", i+1, s.Endpoint, s.ErrorCount, s.LastMessage)
	}
}
