package lib

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var biggest_label float32 = 0

// ImplantWidget the default widget for each implant that calls back
type ImplantWidget struct {
	widget.BaseWidget
	IP                string
	Port              int
	Last_Check_In     string
	Alive             bool
	Detected_Interval string
	Next_Command_Time string
}

type implantWidgetRenderer struct {
	ip                    *canvas.Text
	port                  *canvas.Text
	last_check_in         *canvas.Text
	alive                 *canvas.Text
	detected_interval     *canvas.Text
	next_command_time     *canvas.Text
	lbl_ip                *canvas.Text
	lbl_port              *canvas.Text
	lbl_last_check_in     *canvas.Text
	lbl_alive             *canvas.Text
	lbl_detected_interval *canvas.Text
	lbl_next_command_time *canvas.Text
	objects               []fyne.CanvasObject
	implantWidget         *ImplantWidget
}

func (i *implantWidgetRenderer) Refresh() {
	i.ip.Text = i.implantWidget.IP
	i.port.Text = strconv.Itoa(i.implantWidget.Port)
	i.last_check_in.Text = i.implantWidget.Last_Check_In
	i.alive.Text = strconv.FormatBool(i.implantWidget.Alive)
	i.detected_interval.Text = i.implantWidget.Detected_Interval
	i.next_command_time.Text = i.implantWidget.Next_Command_Time
	canvas.Refresh(i.implantWidget)
}
func (i *implantWidgetRenderer) Objects() []fyne.CanvasObject {
	return i.objects
}
func (i *implantWidgetRenderer) Destroy() {
}
func align_label_field(lbl *canvas.Text, val *canvas.Text, pos *fyne.Position) {
	// Move label
	lbl.Move(*pos)
	// Indent
	pos.X += biggest_label
	// Move Entry
	val.Move(*pos)
	// Add Height
	pos.Y += lbl.MinSize().Height
	// Subtract Indent
	pos.X -= biggest_label
}
func (i *implantWidgetRenderer) Layout(size fyne.Size) {
	pos := fyne.NewPos(theme.Padding()/2, theme.Padding()/2)
	// Detect biggest label
	for _, o := range i.Objects() {
		if o.MinSize().Width > biggest_label {
			biggest_label = o.MinSize().Width
		}
	}
	align_label_field(i.lbl_ip, i.ip, &pos)
	align_label_field(i.lbl_port, i.port, &pos)
	align_label_field(i.lbl_last_check_in, i.last_check_in, &pos)
	align_label_field(i.lbl_alive, i.alive, &pos)
	align_label_field(i.lbl_detected_interval, i.detected_interval, &pos)
	align_label_field(i.lbl_next_command_time, i.next_command_time, &pos)
}
func (i *implantWidgetRenderer) MinSize() fyne.Size {
	return fyne.NewSize(100, 100)
}
func NewImplantWidget(ip string) *ImplantWidget {
	implantwidget := &ImplantWidget{
		IP:                ip,
		Port:              50555,
		Last_Check_In:     "",
		Alive:             false,
		Detected_Interval: "",
		Next_Command_Time: "",
	}

	return implantwidget
}

func (i *ImplantWidget) CreateRenderer() fyne.WidgetRenderer {
	i.ExtendBaseWidget(i)
	IP := canvas.NewText(i.IP, theme.ForegroundColor())
	Port := canvas.NewText(strconv.Itoa(i.Port), theme.ForegroundColor())
	Last_Check_In := canvas.NewText(i.Last_Check_In, theme.ForegroundColor())
	Alive := canvas.NewText(strconv.FormatBool(i.Alive), theme.ForegroundColor())
	Detected_Interval := canvas.NewText(i.Detected_Interval, theme.ForegroundColor())
	Next_Command_Time := canvas.NewText(i.Next_Command_Time, theme.ForegroundColor())
	lbl_ip := canvas.NewText("IP Address:", theme.ForegroundColor())
	lbl_last_check_in := canvas.NewText("Last Check In:", theme.ForegroundColor())
	lbl_alive := canvas.NewText("Alive?", theme.ForegroundColor())
	lbl_detected_interval := canvas.NewText("Detected Interval:", theme.ForegroundColor())
	lbl_port := canvas.NewText("Port:", theme.ForegroundColor())
	lbl_next_command_time := canvas.NewText("Next Command Time:", theme.ForegroundColor())
	r_o := []fyne.CanvasObject{
		lbl_ip,
		IP,
		lbl_port,
		Port,
		lbl_last_check_in,
		Last_Check_In,
		lbl_alive,
		Alive,
		lbl_detected_interval,
		Detected_Interval,
		lbl_next_command_time,
		Next_Command_Time}

	r := &implantWidgetRenderer{
		ip:                    IP,
		last_check_in:         Last_Check_In,
		alive:                 Alive,
		detected_interval:     Detected_Interval,
		port:                  Port,
		next_command_time:     Next_Command_Time,
		lbl_ip:                lbl_ip,
		lbl_last_check_in:     lbl_last_check_in,
		lbl_alive:             lbl_alive,
		lbl_detected_interval: lbl_detected_interval,
		lbl_port:              lbl_port,
		lbl_next_command_time: lbl_next_command_time,
		objects:               r_o,
		implantWidget:         i,
	}
	return r
}

func (i *ImplantWidget) MinSize() fyne.Size {
	i.ExtendBaseWidget(i)
	return i.BaseWidget.MinSize()
}

// Update_Field allows you the program to update each field, it also auto calculates the detected interval and sets the next time which the implant will listen for commands
// TODO: #1 Need to calculate an average, to better predict the interval, this should help latency issues.
func (t *ImplantWidget) Update_Field(field string, value string) {
	switch field {
	case "Last_Check_In":
		old_time, e := time.Parse(time.RFC3339, t.Last_Check_In)
		if e != nil {
			t.Detected_Interval = "Unknown"
		} else {
			time_diff := time.Since(old_time)
			halved, _ := time.ParseDuration(fmt.Sprintf("%fs", time_diff.Seconds()/2))
			next_time_division := time.Now().Add(halved)
			t.Detected_Interval = time_diff.Truncate(time.Millisecond).Round(time.Second).String()
			t.Next_Command_Time = next_time_division.Format(time.RFC3339)
		}
		t.Last_Check_In = value
	case "Alive":
		switch value {
		case "true":
			t.Alive = true
		case "false":
			t.Alive = false
		default:
			fmt.Println("Unknown Value for Alive true/false")
		}
	case "Port":
		port, _ := strconv.Atoi(value)
		t.Port = port
	}
}

//TODO: #2 Handle modal details to interact with implant
func (t *ImplantWidget) Tapped(_ *fyne.PointEvent) {
	log.Println("I have been tapped")
}
