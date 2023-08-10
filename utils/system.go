package utils

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

// MemUsageReport outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func MemUsageReport() string {
	var sb strings.Builder
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	sb.WriteString(fmt.Sprintf("Alloc = %s\t", humanize.Bytes(m.Alloc)))
	sb.WriteString(fmt.Sprintf("TotalAlloc = %s\t", humanize.Bytes(m.TotalAlloc)))
	sb.WriteString(fmt.Sprintf("Sys = %s\t", humanize.Bytes(m.Sys)))
	sb.WriteString(fmt.Sprintf("NumGC = %v\n", m.NumGC))
	return sb.String()
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
