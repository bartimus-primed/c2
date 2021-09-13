package lib

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	check_in_history  []string
	popper            *widget.PopUp
}

type implantWidgetRenderer struct {
	background            *canvas.Rectangle
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
	i.port.Text = strconv.Itoa(i.implantWidget.Port)
	i.last_check_in.Text = i.implantWidget.Last_Check_In
	if i.implantWidget.Alive {
		i.alive.Color = theme.PrimaryColorNamed("green")
	} else {
		i.alive.Color = theme.ErrorColor()
	}
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
	// Setup Background
	i.background.StrokeColor = theme.ShadowColor()
	i.background.StrokeWidth = 2
	i.background.Move(pos)
	size.Width += 5
	i.background.Resize(size)
	pos.X += 5
	pos.Y += 5
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
	amount_width := (biggest_label * 2) + theme.Padding()
	amount_height := (i.lbl_ip.MinSize().Height + theme.Padding()) * 6
	return fyne.NewSize(amount_width+10, amount_height)
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
	IP.TextSize = 15
	Port := canvas.NewText(strconv.Itoa(i.Port), theme.ForegroundColor())
	Port.TextSize = 15
	Last_Check_In := canvas.NewText(i.Last_Check_In, theme.ForegroundColor())
	Last_Check_In.TextSize = 15
	Alive := canvas.NewText(strconv.FormatBool(i.Alive), theme.ForegroundColor())
	Alive.TextSize = 15
	Detected_Interval := canvas.NewText(i.Detected_Interval, theme.ForegroundColor())
	Detected_Interval.TextSize = 15
	Next_Command_Time := canvas.NewText(i.Next_Command_Time, theme.ForegroundColor())
	Next_Command_Time.TextSize = 15
	lbl_ip := canvas.NewText("IP Address:", theme.ForegroundColor())
	lbl_last_check_in := canvas.NewText("Last Check In:", theme.ForegroundColor())
	lbl_alive := canvas.NewText("Alive?", theme.ForegroundColor())
	lbl_detected_interval := canvas.NewText("Detected Interval:", theme.ForegroundColor())
	lbl_port := canvas.NewText("Port:", theme.ForegroundColor())
	lbl_next_command_time := canvas.NewText("Next Command Time:", theme.ForegroundColor())
	BACKGROUND := canvas.NewRectangle(theme.HoverColor())
	r_o := []fyne.CanvasObject{
		BACKGROUND,
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
		background:            BACKGROUND,
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
func (i *ImplantWidget) Update_Field(field string, value string) {
	switch field {
	case "Last_Check_In":
		old_time, e := time.Parse(time.RFC3339, i.Last_Check_In)
		i.check_in_history = append(i.check_in_history, i.Last_Check_In)
		if e != nil {
			i.Detected_Interval = "Unknown"
		} else {
			time_diff := time.Since(old_time)
			halved, _ := time.ParseDuration(fmt.Sprintf("%fs", time_diff.Seconds()/2))
			next_time_division := time.Now().Add(halved)
			i.Detected_Interval = time_diff.Truncate(time.Millisecond).Round(time.Second).String()
			i.Next_Command_Time = next_time_division.Format(time.RFC3339)
		}
		i.Last_Check_In = value
	case "Alive":
		switch value {
		case "true":
			i.Alive = true
		case "false":
			i.Alive = false
		default:
			fmt.Println("Unknown Value for Alive true/false")
		}
	case "Port":
		port, _ := strconv.Atoi(value)
		i.Port = port
	}
}

//TODO: #2 Handle modal details to interact with implant
func (i *ImplantWidget) Tapped(_ *fyne.PointEvent) {
	i.Build_Popup()
	i.popper.Show()
}

func (i *ImplantWidget) History() int {
	return len(i.check_in_history)
}
func (i *ImplantWidget) Create_Item() fyne.CanvasObject {
	return widget.NewLabel("0000000000000000000000000000")
}
func (i *ImplantWidget) Update_Item(item int, lbl fyne.CanvasObject) {
	if l, ok := lbl.(*widget.Label); ok {
		l.Text = i.check_in_history[item]
	}
}
func (i *ImplantWidget) Build_Popup() {
	max_con := container.NewCenter()
	m_can := fyne.CurrentApp().Driver().CanvasForObject(i)
	m_can.Content().Resize(fyne.NewSize(biggest_label*1.2, biggest_label*1.1))
	i.popper = widget.NewModalPopUp(max_con, m_can)
	exit_button := widget.NewButtonWithIcon("Go Back", theme.CancelIcon(), i.Close_Popup)
	go_button := widget.NewButtonWithIcon("Go Hands On", theme.ComputerIcon(), func() { fmt.Println("GOING HANDS ON!") })
	list_of_times := widget.NewList(i.History, i.Create_Item, i.Update_Item)
	lbl_check_in := widget.NewLabel("Check In History:")
	max_con.Add(lbl_check_in)
	max_con.Add(list_of_times)
	max_con.Add(container.NewHBox(go_button, layout.NewSpacer(), exit_button))
}

func (i *ImplantWidget) Close_Popup() {
	i.popper.Hide()
}
