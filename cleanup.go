package main

import (
	"fmt"
	"os"
	"time"
)

// Function to manage disk space and perform cleanup
func cleanup() {
	// Get the disk space statistics
	var stat syscall.Statfs_t
	err := syscall.Getfsstat("/", &stat)
	if err != nil {
		fmt.Println("Error getting disk space stats:", err)
		return
	}

	// Check current disk usage
	availableSpace := stat.Bfree * uint64(stat.Bsize)
	if availableSpace < 100*1024*1024 { // If less than 100MB available
		fmt.Println("Running cleanup...")
		// Perform cleanup actions here
	}
}

// Function to run periodic cleanup every hour
func startPeriodicCleanup() {
	// Run cleanup function immediately on startup
	cleanup()

	// Schedule cleanup every hour
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			cleanup()
		}
	}
}

func main() {
	startPeriodicCleanup()
}