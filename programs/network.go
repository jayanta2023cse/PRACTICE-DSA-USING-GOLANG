package programs

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/user"
	"runtime"
	"time"
)

type Device struct {
	Host     string   `json:"host"`
	OS       string   `json:"os,omitempty"`
	Username string   `json:"username,omitempty"`
	IPs      []string `json:"ips"`
	Alive    bool     `json:"alive,omitempty"`
}

type NetworkScan struct {
	Local   Device   `json:"local"`
	Network []Device `json:"network"`
}

func getLocalDeviceInfo() (Device, error) {
	host, err := os.Hostname()
	if err != nil {
		return Device{}, fmt.Errorf("failed to get hostname: %v", err)
	}

	usr, err := user.Current()
	if err != nil {
		return Device{}, fmt.Errorf("failed to get current user: %v", err)
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return Device{}, fmt.Errorf("failed to get interfaces: %v", err)
	}

	var ips []string
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	return Device{
		Host:     host,
		OS:       runtime.GOOS,
		Username: usr.Username,
		IPs:      ips,
		Alive:    true,
	}, nil
}

func scanNetwork() ([]Device, error) {
	var devices []Device
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get interfaces: %v", err)
	}

	ports := []string{"80", "443"}
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				ip := ipnet.IP.To4()
				network := ip.Mask(ipnet.Mask)
				// Scan first 10 IPs for simplicity
				for i := 1; i <= 10; i++ {
					targetIP := net.IPv4(network[0], network[1], network[2], byte(i)).String()
					for _, port := range ports {
						target := fmt.Sprintf("%s:%s", targetIP, port)
						conn, err := net.DialTimeout("tcp", target, 200*time.Millisecond)
						if err == nil {
							conn.Close()
							hostname, _ := net.LookupHost(targetIP)
							host := targetIP
							if len(hostname) > 0 {
								host = hostname[0]
							}
							devices = append(devices, Device{
								Host:  host,
								IPs:   []string{targetIP},
								Alive: true,
							})
							break // Stop after first successful port
						}
					}
				}
			}
		}
	}
	return devices, nil
}

func GetNetWork() {
	fmt.Println("Starting network scan...")

	local, err := getLocalDeviceInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting local device info: %v\n", err)
		os.Exit(1)
	}

	network, err := scanNetwork()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning network: %v\n", err)
		os.Exit(1)
	}

	result := NetworkScan{
		Local:   local,
		Network: network,
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile("network_info.json", data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing JSON file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Scan completed. Found %d devices. Results written to network_info.json\n", len(network))
}
