package lib

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var biggest_label float32 = 0

// ImplantWidget the default widget for each implant that calls back
type ImplantWidget struct {
	widget.BaseWidget
	IP                string
	Last_Check_In     string
	Alive             bool
	Detected_Interval string
	Port              int
	Next_Command_Time string
	bindings          *ImplantBindings
}

// ImplantBindings holds all the fyne related functionality allowing the UI to update as needed
type ImplantBindings struct {
	IP                binding.String
	Last_Check_In     binding.String
	Alive             binding.Bool
	Detected_Interval binding.String
	Port              binding.Int
	Next_Command_Time binding.String
}

type implantWidgetRenderer struct {
	ip                    *canvas.Text
	last_check_in         *canvas.Text
	alive                 *canvas.Text
	detected_interval     *canvas.Text
	port                  *canvas.Text
	next_command_time     *canvas.Text
	lbl_ip                *canvas.Text
	lbl_last_check_in     *canvas.Text
	lbl_alive             *canvas.Text
	lbl_detected_interval *canvas.Text
	lbl_port              *canvas.Text
	lbl_next_command_time *canvas.Text
	objects               []fyne.CanvasObject
	implantWidget         *ImplantWidget
}

func (i *implantWidgetRenderer) Refresh() {
	IP_TEXT, _ := i.implantWidget.bindings.IP.Get()
	i.ip.Text = IP_TEXT
	LAST_CHECK_IN_TEXT, _ := i.implantWidget.bindings.Last_Check_In.Get()
	i.last_check_in.Text = LAST_CHECK_IN_TEXT
	ALIVE_TEXT, _ := binding.BoolToString(i.implantWidget.bindings.Alive).Get()
	i.alive.Text = ALIVE_TEXT
	DETECTED_INTERVAL_TEXT, _ := i.implantWidget.bindings.Detected_Interval.Get()
	i.detected_interval.Text = DETECTED_INTERVAL_TEXT
	PORT_TEXT, _ := binding.IntToString(i.implantWidget.bindings.Port).Get()
	i.port.Text = PORT_TEXT
	NEXT_COMMAND_TIME_TEXT, _ := i.implantWidget.bindings.Next_Command_Time.Get()
	i.next_command_time.Text = NEXT_COMMAND_TIME_TEXT
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
	size = size.Subtract(fyne.NewSize(theme.Padding(), theme.Padding()))
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
func NewImplantWidget(title string, subtitle string) *ImplantWidget {
	implantwidget := &ImplantWidget{
		IP:                title,
		Last_Check_In:     "",
		Alive:             false,
		Detected_Interval: "",
		Port:              50555,
		Next_Command_Time: "",
	}
	implantwidget.bindings = &ImplantBindings{
		IP:                binding.BindString(&implantwidget.IP),
		Last_Check_In:     binding.BindString(&implantwidget.Last_Check_In),
		Alive:             binding.BindBool(&implantwidget.Alive),
		Detected_Interval: binding.BindString(&implantwidget.Detected_Interval),
		Port:              binding.BindInt(&implantwidget.Port),
		Next_Command_Time: binding.BindString(&implantwidget.Next_Command_Time),
	}
	implantwidget.ExtendBaseWidget(implantwidget)
	return implantwidget
}

func (i *ImplantWidget) CreateRenderer() fyne.WidgetRenderer {
	IP_TEXT, _ := i.bindings.IP.Get()
	IP := canvas.NewText(IP_TEXT, theme.ForegroundColor())
	LAST_CHECK_IN_TEXT, _ := i.bindings.Last_Check_In.Get()
	Last_Check_In := canvas.NewText(LAST_CHECK_IN_TEXT, theme.ForegroundColor())
	ALIVE_TEXT, _ := binding.BoolToString(i.bindings.Alive).Get()
	Alive := canvas.NewText(ALIVE_TEXT, theme.ForegroundColor())
	DETECTED_INTERVAL_TEXT, _ := i.bindings.Detected_Interval.Get()
	Detected_Interval := canvas.NewText(DETECTED_INTERVAL_TEXT, theme.ForegroundColor())
	PORT_TEXT, _ := binding.IntToString(i.bindings.Port).Get()
	Port := canvas.NewText(PORT_TEXT, theme.ForegroundColor())
	NEXT_COMMAND_TIME_TEXT, _ := i.bindings.Next_Command_Time.Get()
	Next_Command_Time := canvas.NewText(NEXT_COMMAND_TIME_TEXT, theme.ForegroundColor())
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
	return i.BaseWidget.MinSize()
}

// Update_Field allows you the program to update each field, it also auto calculates the detected interval and sets the next time which the implant will listen for commands
// TODO: #1 Need to calculate an average, to better predict the interval, this should help latency issues.
func (t *ImplantWidget) Update_Field(field string, value string) {
	defer t.Refresh()
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

//TODO: #2 Handle modal details to interact with implant
func (t *ImplantWidget) Tapped(_ *fyne.PointEvent) {
	log.Println("I have been tapped")
}
