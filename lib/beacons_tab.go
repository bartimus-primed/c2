package lib

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Main Page Layout
func Get_Beacons_Tab(c2_status binding.String, app_status binding.String, beacon_map map[string][]string) *container.TabItem {
	beacon_list := widget.NewTreeWithStrings(beacon_map)
	beacon_container := container.NewMax(beacon_list)
	// beacon_list.Refresh()
	beacon_tab := container.NewTabItemWithIcon("Beacons", theme.ViewFullScreenIcon(), beacon_container)
	return beacon_tab
}
