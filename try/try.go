package try

import (
	"fmt"
	"runtime"
	"time"
)

// Try try it
func Try() {
	fmt.Printf("try %s at %s\n", runtime.Version(), time.Now())
}
