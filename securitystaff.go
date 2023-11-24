package main

import (
	"fmt"
)

// SecurityStaff represents the security staff responsible for handling security-related tasks.
type SecurityStaff struct {
}

var securityStaffInstance = &SecurityStaff{} // Global instance of SecurityStaff

// Notify prints a notification message when the parking lot is full.
func (s *SecurityStaff) Notify() {
	fmt.Println("Security staff notified: Parking lot is full. Additional security measures activated.")
}
