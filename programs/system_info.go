package programs

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"
)

// RAMInfo holds total, free, used RAM in GB and usage percentage
type RAMInfo struct {
	TotalGB float64
	FreeGB  float64
	UsedGB  float64
	UsedPct float64
}

// GetRAMInfo returns total, free, used RAM in GB and usage percentage (cross-platform)
func GetRAMInfo() (RAMInfo, error) {
	switch runtime.GOOS {
	case "linux":
		return getRAMInfoLinux()
	case "darwin":
		return getRAMInfoMacOS()
	case "windows":
		return getRAMInfoWindows()
	default:
		return RAMInfo{}, fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

// -------------------- Linux --------------------
func getRAMInfoLinux() (RAMInfo, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return RAMInfo{}, err
	}
	defer file.Close()

	var totalRAM, freeRAM int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "MemTotal:") {
			fmt.Sscanf(line, "MemTotal: %d kB", &totalRAM)
			totalRAM *= 1024
		}
		if strings.HasPrefix(line, "MemAvailable:") {
			fmt.Sscanf(line, "MemAvailable: %d kB", &freeRAM)
			freeRAM *= 1024
		}
	}
	if err := scanner.Err(); err != nil {
		return RAMInfo{}, err
	}

	totalGB := float64(totalRAM) / (1024 * 1024 * 1024)
	freeGB := float64(freeRAM) / (1024 * 1024 * 1024)
	usedGB := totalGB - freeGB
	usedPct := (usedGB / totalGB) * 100

	return RAMInfo{TotalGB: totalGB, FreeGB: freeGB, UsedGB: usedGB, UsedPct: usedPct}, nil
}

// -------------------- macOS --------------------
func getRAMInfoMacOS() (RAMInfo, error) {
	// Total RAM
	cmd := exec.Command("sysctl", "-n", "hw.memsize")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return RAMInfo{}, err
	}
	memSize := strings.TrimSpace(out.String())
	var total int64
	fmt.Sscanf(memSize, "%d", &total)

	// Free RAM from "vm_stat"
	cmd = exec.Command("vm_stat")
	out.Reset()
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return RAMInfo{}, err
	}

	var pageSize int64 = 4096
	var freePages, inactivePages int64
	lines := strings.Split(out.String(), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Pages free:") {
			fmt.Sscanf(line, "Pages free: %d.", &freePages)
		}
		if strings.HasPrefix(line, "Pages inactive:") {
			fmt.Sscanf(line, "Pages inactive: %d.", &inactivePages)
		}
		if strings.HasPrefix(line, "page size of") {
			fmt.Sscanf(line, "page size of %d bytes", &pageSize)
		}
	}
	free := (freePages + inactivePages) * pageSize

	totalGB := float64(total) / (1024 * 1024 * 1024)
	freeGB := float64(free) / (1024 * 1024 * 1024)
	usedGB := totalGB - freeGB
	usedPct := (usedGB / totalGB) * 100

	return RAMInfo{TotalGB: totalGB, FreeGB: freeGB, UsedGB: usedGB, UsedPct: usedPct}, nil
}

// -------------------- Windows --------------------
func getRAMInfoWindows() (RAMInfo, error) {
	type memoryStatusEx struct {
		length               uint32
		memoryLoad           uint32
		totalPhys            uint64
		availPhys            uint64
		totalPageFile        uint64
		availPageFile        uint64
		totalVirtual         uint64
		availVirtual         uint64
		availExtendedVirtual uint64
	}

	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	globalMemoryStatusEx := kernel32.NewProc("GlobalMemoryStatusEx")

	var memStatus memoryStatusEx
	memStatus.length = uint32(unsafe.Sizeof(memStatus))

	ret, _, err := globalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memStatus)))
	if ret == 0 {
		return RAMInfo{}, err
	}

	totalGB := float64(memStatus.totalPhys) / (1024 * 1024 * 1024)
	freeGB := float64(memStatus.availPhys) / (1024 * 1024 * 1024)
	usedGB := totalGB - freeGB
	usedPct := (usedGB / totalGB) * 100

	return RAMInfo{TotalGB: totalGB, FreeGB: freeGB, UsedGB: usedGB, UsedPct: usedPct}, nil
}

func PrintRAMInfo() {
	info, err := GetRAMInfo()
	if err != nil {
		fmt.Println("Error fetching RAM:", err)
		return
	}
	fmt.Printf("Total: %.2f GB | Free: %.2f GB | Used: %.2f GB (%.1f%%)\n", info.TotalGB, info.FreeGB, info.UsedGB, info.UsedPct)
}
