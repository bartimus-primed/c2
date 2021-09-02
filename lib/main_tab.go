package lib

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Main Page Layout
func Get_Home_Tab() *container.TabItem {
	changelog_map := map[string][]string{}
	changelog_map[""] = []string{"Change Log", "V0.1", "V0.2"}
	changelog_map["V0.1"] = []string{"GUI Access", "Working C2 Start", "Working Beacon Listing"}
	changelog_map["V0.2"] = []string{"Coming Soon..."}
	changelog_list := widget.NewTreeWithStrings(changelog_map)
	home_container := container.NewMax(changelog_list)
	changelog_list.Refresh()
	main_tab := container.NewTabItemWithIcon("Main", theme.HomeIcon(), home_container)
	return main_tab
}
