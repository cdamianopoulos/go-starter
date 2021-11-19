package dbconn

import (
	"fmt"
	"strconv"
)

const (
	maxBitSize  = 12
	parseAsBase = 10
)

// The Set functions are used to validate a string as a quantity of open or idle database connections
// used by envconfig.Setter interface.
type (
	OpenQty int
	IdleQty int
)

// Set expects qty to be an integer between 0 and 4095.
func (o *OpenQty) Set(qty string) error {
	u, err := strconv.ParseUint(qty, parseAsBase, maxBitSize) // Limits values between 0 and 4095 (1<<12 - 1).
	if err != nil {
		return err
	}

	*o = OpenQty(u)
	return nil
}

// Set expects qty to be an integer between -1 and 4095.
func (i *IdleQty) Set(qty string) error {
	u, err := strconv.ParseInt(qty, parseAsBase, maxBitSize+1) // Limits values between -4096 and 4095.
	if err != nil {
		return err
	}

	if u < -1 {
		return fmt.Errorf("value %s is out of range. Expecting value to be between -1 and %d", qty, 1<<maxBitSize-1)
	}

	*i = IdleQty(u)
	return nil
}
