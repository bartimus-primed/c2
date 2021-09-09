package lib

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

var beacon_container *fyne.Container
var beacon_map map[string]*ImplantWidget

// Main Page Layout
func Get_Beacons_Tab(c2_status binding.String, app_status binding.String) *container.TabItem {
	// beacon_list.Refresh()
	beacon_map = make(map[string]*ImplantWidget)
	beacon_container = container.NewVBox()
	beacon_h_container := container.NewHBox(layout.NewSpacer(), beacon_container, layout.NewSpacer())
	beacon_scroll_box := container.NewVScroll(beacon_h_container)
	beacon_tab := container.NewTabItemWithIcon("Beacons", theme.ViewFullScreenIcon(), beacon_scroll_box)
	return beacon_tab
}

func Add_Beacon(ip string, port string, status string) {
	beacon := beacon_map[ip]
	// Does the beacon exist? If not create one.
	if beacon == nil {
		fmt.Println("Beacon is nil")
		beacon = NewImplantWidget(ip)
		beacon.Update_Field("Last_Check_In", time.Now().Format(time.RFC3339))
		beacon_map[ip] = beacon
		beacon_padding := container.NewPadded(beacon)
		beacon_container.Add(beacon_padding)
	}
	beacon.Update_Field("Port", port)
	switch status {
	case "beacon":
		beacon.Update_Field("Alive", "true")
		// beacon.SetSubTitle("beaconing...")
	case "kill":
		beacon.Update_Field("Alive", "false")
		// beacon.SetSubTitle("died")
	}
	beacon.Update_Field("Last_Check_In", time.Now().Format(time.RFC3339))
	beacon.Refresh()
}
