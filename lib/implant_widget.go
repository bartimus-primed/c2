package lib

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// ImplantWidget the default widget for each implant that calls back
type ImplantWidget struct {
	widget.Card
	details  *ImplantDetails
	bindings *ImplantBindings
}

// ImplantDetails holds the native go functionality and values
type ImplantDetails struct {
	Last_Check_In     string
	Alive             bool
	Detected_Interval string
	Port              int
	Next_Command_Time string
}

// ImplantBindings holds all the fyne related functionality allowing the UI to update as needed
type ImplantBindings struct {
	Last_Check_In     binding.String
	Alive             binding.Bool
	Detected_Interval binding.String
	Port              binding.Int
	Next_Command_Time binding.String
}

// NewImplantWidget takes in the title which is the IP and a subtitle which changes depending on the implant status
func NewImplantWidget(title string, subtitle string) *ImplantWidget {
	implant := &ImplantWidget{}
	implant.ExtendBaseWidget(implant)
	implant.Title = title
	implant.Subtitle = subtitle
	implant.Image = canvas.NewImageFromResource(theme.ComputerIcon())
	// Just some default information that should change as the correct data comes in.
	implant.details = &ImplantDetails{
		Last_Check_In:     "unknown",
		Alive:             false,
		Detected_Interval: "unknown",
		Port:              50555,
		Next_Command_Time: "unknown",
	}
	// Assign the values to the bindings for the ui
	implant.bindings = &ImplantBindings{
		Last_Check_In:     binding.BindString(&implant.details.Last_Check_In),
		Alive:             binding.BindBool(&implant.details.Alive),
		Detected_Interval: binding.BindString(&implant.details.Detected_Interval),
		Port:              binding.BindInt(&implant.details.Port),
		Next_Command_Time: binding.BindString(&implant.details.Next_Command_Time),
	}
	// Labels
	lbl_last_check_in := widget.NewLabel("Last Check In Time:")
	lbl_alive := widget.NewLabel("Alive:")
	lbl_ip_address := widget.NewLabel("Detected_Interval:")
	lbl_port := widget.NewLabel("Port:")
	lbl_next_command_time := widget.NewLabel("Next Command Time:")
	// Values
	val_last_check_in := widget.NewLabelWithData(implant.bindings.Last_Check_In)
	val_alive := widget.NewLabelWithData(binding.BoolToString(implant.bindings.Alive))
	val_ip_address := widget.NewLabelWithData(implant.bindings.Detected_Interval)
	val_port := widget.NewLabelWithData(binding.IntToString(implant.bindings.Port))
	val_next_command_time := widget.NewLabelWithData(implant.bindings.Next_Command_Time)
	// Set up the form.
	implant_form := container.New(layout.NewFormLayout(), lbl_last_check_in, val_last_check_in, lbl_alive, val_alive, lbl_ip_address, val_ip_address, lbl_port, val_port, lbl_next_command_time, val_next_command_time)
	implant.Content = container.NewCenter(implant_form)

	return implant
}

// Update_Field allows you the program to update each field, it also auto calculates the detected interval and sets the next time which the implant will listen for commands
// TODO: #1 Need to calculate an average, to better predict the interval, this should help latency issues.
func (t *ImplantWidget) Update_Field(field string, value string) {
	switch field {
	case "Last_Check_In":
		t.bindings.Last_Check_In.Set(value)
	case "Alive":
		switch value {
		case "true":
			t.bindings.Alive.Set(true)
		case "false":
			t.bindings.Alive.Set(false)
		default:
			fmt.Println("Unknown Value for Alive true/false")
		}
	case "Detected_Interval":
		tm, e := t.bindings.Last_Check_In.Get()
		if e != nil {
			t.bindings.Detected_Interval.Set("N/A")
		} else {
			old_time, e := time.Parse(time.RFC3339, tm)
			if e != nil {
				t.bindings.Detected_Interval.Set("Unknown")
			} else {
				time_diff := time.Since(old_time)
				halved, _ := time.ParseDuration(fmt.Sprintf("%fs", time_diff.Seconds()/2))
				next_time_division := time.Now().Add(halved)
				t.bindings.Detected_Interval.Set(time_diff.Truncate(time.Millisecond).Round(time.Second).String())
				t.bindings.Next_Command_Time.Set(next_time_division.Format(time.RFC3339))
			}
		}
	case "Port":
		port, _ := strconv.Atoi(value)
		t.bindings.Port.Set(port)
	}
}
func (t *ImplantWidget) Tapped(_ *fyne.PointEvent) {
	log.Println("I have been tapped")
	t.Refresh()
}

func (t *ImplantWidget) TappedSecondary(_ *fyne.PointEvent) {
}
