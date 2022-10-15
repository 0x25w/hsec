package common

import (
	"github.com/shirou/gopsutil/process"
)

// GetProcess is a function that returns the current process.
func GetProcess() {
	// getProcess
	AllProcess := scanProcess()
	// match from map
	Match(AllProcess)
}

func scanProcess() []string {
	// 扫描进程
	processes, _ := process.Processes()
	var procs []string
	for _, proc := range processes {
		name, _ := proc.Exe()
		procs = append(procs, name)
	}
	return procs
}
