package util

import (
	"sync"
)

var (
	driversMu sync.RWMutex
)

// Driver 这个是给通用包注册引擎用的
type Driver interface {
	AID() uint
	KEY() string
}

// Drivers 引擎组
type Drivers map[string]Driver

// RegDriver 将引擎注册到引擎组中
func RegDriver(drivers Drivers, driver Driver) {

	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("RegDriver: Register driver is nil")
	}
	key := driver.KEY()
	if _, dup := drivers[key]; dup {
		println("RegDriver: Register called twice for driver " + driver.KEY())
		return
		// panic("RegDriver: Register called twice for driver " + driver.KEY())
	}
	drivers[driver.KEY()] = driver
}
