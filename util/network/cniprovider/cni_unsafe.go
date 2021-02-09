// +build linux

package cniprovider

import (
	_ "unsafe" // required for go:linkname.
)

//go:linkname beforeFork runtime.beforefork
func beforeFork()

//go:linkname afterFork runtime.afterfork
func afterFork()

//go:linkname afterForkInChild syscall.runtime_AfterForkInChild
func afterForkInChild()
