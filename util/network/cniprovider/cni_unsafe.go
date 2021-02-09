// +build linux

package cniprovider

import (
	_ "unsafe" // required for go:linkname.
)

//go:linkname beforeFork syscall.runtime__BeforeFork
func beforeFork()

//go:linkname afterFork syscall.runtime__AfterFork
func afterFork()

//go:linkname afterForkInChild syscall.runtime__AfterForkInChild
func afterForkInChild()
