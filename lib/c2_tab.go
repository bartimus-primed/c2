package lib

import (
	"fmt"
	"net"
	"strings"
	"time"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var c2_addr = ""
var c2_port = 0
var connection *net.UDPConn
var default_port = binding.NewInt()
var serverOutput = make(chan string)
var isRunning = false

// Main Page Layout
func Get_C2_Tab(c2_status binding.String, app_status binding.String, connectedBeacons map[string][]string, beacon_tab *container.TabItem) *container.TabItem {
	default_port.Set(50555)
	listening_port := widget.NewEntryWithData(binding.IntToString(default_port))
	listening_port_label := widget.NewLabel("Listen on Port:")
	listening_addr_label := widget.NewLabel("Listen on Address:")
	listening_addr := widget.NewSelect(GetInterfaceAddresses(), func(a string) { c2_addr = a })

	start_c2_button := widget.NewButtonWithIcon("Start C2 Server", theme.MediaPlayIcon(), func() {
		if !isRunning {
			c2_port, _ = default_port.Get()
			c2_addr = strings.Split(c2_addr, "/")[0]
			c2_status.Set(fmt.Sprintf("Listening on %s:%d", c2_addr, c2_port))
			app_status.Set("Started C2 Server")
			go start_c2_server()
			go func() {
				for a := range serverOutput {
					switch a {
					case "closed":
						isRunning = false
						c2_status.Set("Not Running...")
						app_status.Set("C2 Server Closed")
					default:
						beacon_info := strings.Split(a, ",")
						beacon_ip := strings.Split(beacon_info[0], ":")[0]
						beacon_status := beacon_info[1]
						app_status.Set(fmt.Sprintf("Received Beacon From: %s", beacon_ip))
						if connectedBeacons[beacon_ip] != nil {
							if beacon_status == "dead" {
								app_status.Set(fmt.Sprintf("%s Died", beacon_ip))
							}
							connectedBeacons[beacon_ip] = append(connectedBeacons[beacon_ip], beacon_status)
						} else {
							connectedBeacons[""] = append(connectedBeacons[""], beacon_ip)
							connectedBeacons[beacon_ip] = []string{beacon_status}
						}
						beacon_tab.Content.Refresh()
					}
				}
			}()
		}
	})
	stop_c2_button := widget.NewButtonWithIcon("Stop C2 Server", theme.MediaStopIcon(), func() {
		if isRunning {
			go func() {
				if connection != nil {
					connection.Close()
				}
			}()
		}
	})
	c2_server_form := container.New(layout.NewFormLayout(), listening_addr_label, listening_addr, listening_port_label, listening_port)
	c2_container := container.NewGridWithColumns(1, c2_server_form, start_c2_button, stop_c2_button)
	c2_tab := container.NewTabItemWithIcon("C2", theme.ComputerIcon(), c2_container)
	return c2_tab
}

func GetInterfaceAddresses() []string {
	results := []string{}
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		addr_str := addr.String()
		if !strings.Contains(addr_str, ":") {
			results = append(results, addr.String())
		}
	}
	return results
}

func start_c2_server() {
	c2_address := net.UDPAddr{
		Port: c2_port,
		IP:   net.ParseIP(c2_addr),
	}
	connection, _ = net.ListenUDP("udp", &c2_address)
	isRunning = true
	buffer := make([]byte, 1024)
	for {
		count, remoteAddress, err := connection.ReadFromUDP(buffer)
		if err != nil {
			serverOutput <- "closed"
			return
		}
		mesg := string(buffer[:count])
		beacon_ip := remoteAddress
		switch mesg {
		case "beacon\n":
			beacon_time := time.Now()
			serverOutput <- fmt.Sprintf("%s,%s", beacon_ip, beacon_time)
		case "kill\n":
			serverOutput <- fmt.Sprintf("%s,dead", beacon_ip)
		case "exit\n":
			connection.Close()
		}
	}
}
