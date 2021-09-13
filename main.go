package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/bartimus-primed/c2/lib"
)

var app_status = binding.NewString()
var c2_status = binding.NewString()
var connectedBeacons = map[string][]string{}
var myApp fyne.App
var myWindow fyne.Window

func main() {
	// App Settings
	myApp = app.New()
	myWindow = myApp.NewWindow("Go Hands-On Security Toolkit")

	// Set Default Statuses
	c2_status.Set("Not Running...")
	app_status.Set("GHOST v0.2")
	myWindow.CenterOnScreen()
	// myWindow.SetFullScreen(true)

	// Set Tabs for Navigation
	// Add Tabs to Tab Bar
	// vars for each tab since we might need to pass them around
	connectedBeacons[""] = []string{}
	home_tab := lib.Get_Home_Tab()
	beacons_tab := lib.Get_Beacons_Tab(c2_status, app_status)
	c2_tab := lib.Get_C2_Tab(c2_status, app_status)
	settings_tab := lib.Get_Settings_Tab(c2_status, app_status)
	tabs := container.NewAppTabs(home_tab, c2_tab, beacons_tab, settings_tab)

	// Set Layout for Window
	// Top Bar - Title
	top := canvas.NewText("Go Hands-On Security Toolkit", color.White)
	top_bar := container.NewHBox(layout.NewSpacer(), top, layout.NewSpacer())

	// Bottom Bar - C2 Status/App Info/Exit
	// App Info/Status
	app_status_text := widget.NewLabelWithData(app_status)

	// C2 Status
	c2_status_label := widget.NewLabel("C2 Server Status:")
	c2_status_text := widget.NewLabelWithData(c2_status)
	c2_status_container := container.NewHBox(c2_status_label, c2_status_text)

	// Exit button
	exit_button := widget.NewButtonWithIcon("Exit", theme.CancelIcon(), exit_app)
	exit_button.Importance = widget.HighImportance

	// Container to hold status and exit button
	bottom_bar := container.NewHBox(c2_status_container, layout.NewSpacer(), app_status_text, layout.NewSpacer(), exit_button)

	// adding content
	content := container.New(layout.NewBorderLayout(top_bar, bottom_bar, nil, nil), top_bar, bottom_bar, tabs)

	// Set Window Content
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func exit_app() {
	app_status.Set("Exiting...")
	myWindow.Close()
}
