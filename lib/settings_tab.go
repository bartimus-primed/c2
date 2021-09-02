package lib

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Main Page Layout
func Get_Settings_Tab(c2_status binding.String, app_status binding.String) *container.TabItem {
	settings_tab := container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), widget.NewLabel("Settings Page"))
	return settings_tab
}
