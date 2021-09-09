package lib

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
)

var beacon_container *fyne.Container
var beacon_map map[string]*ImplantWidget

// Main Page Layout
func Get_Beacons_Tab(c2_status binding.String, app_status binding.String) *container.TabItem {
	// beacon_list.Refresh()
	beacon_map = make(map[string]*ImplantWidget)
	beacon_container = container.NewGridWrap(fyne.NewSize(300, 300))
	beacon_tab := container.NewTabItemWithIcon("Beacons", theme.ViewFullScreenIcon(), beacon_container)
	return beacon_tab
}

func Add_Beacon(ip string, port string, status string) {
	beacon := beacon_map[ip]
	// Does the beacon exist? If not create one.
	if beacon == nil {
		beacon = NewImplantWidget(ip, "")
		beacon.Update_Field("Last_Check_In", time.Now().Format(time.RFC3339))
		beacon_map[ip] = beacon
		beacon_container.Add(beacon)
	}
	beacon.Update_Field("Port", port)
	beacon.Update_Field("Detected_Interval", "Unknown")
	beacon.Update_Field("Next_Command_Time", "unknown")
	switch status {
	case "beacon":
		beacon.Update_Field("Alive", "true")
		// beacon.SetSubTitle("beaconing...")
	case "kill":
		beacon.Update_Field("Alive", "false")
		// beacon.SetSubTitle("died")
	default:
		beacon.Update_Field("Alive", "unknown")
		// beacon.SetSubTitle("unknown")
	}
	beacon.Update_Field("Last_Check_In", time.Now().Format(time.RFC3339))
}
