package utils

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestMemUsage(t *testing.T) {
	// Print our starting memory usage (should be around 0mb)
	fmt.Print(MemUsageReport())

	var overall [][]int
	for i := 0; i < 4; i++ {

		// Allocate memory using make() and append to overall (so it doesn't get
		// garbage collected). This is to create an ever increasing memory usage
		// which we can track. We're just using []int as an example.
		a := make([]int, 0, 999999)
		overall = append(overall, a)

		// Print our memory usage at each interval
		if len(overall) > 0 { // avoids warning
			fmt.Print(MemUsageReport())
			time.Sleep(100 * time.Millisecond)
		}
	}

	// Clear our memory and print usage, unless the GC has run 'Alloc' will remain the same
	overall = nil
	fmt.Print(MemUsageReport())

	// Force GC to clear up, should see a memory drop
	runtime.GC()
	fmt.Print(MemUsageReport())
}

func TestTimeTrack(t *testing.T) {
	defer TimeTrack(time.Now(), "Test")
	time.Sleep(1 * time.Second)
}
