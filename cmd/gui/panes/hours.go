package panes

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/nh3000-org/radio/config"
	//"github.com/nh3000-org/radio/config"
)

var myrowhours int

func HoursScreen(win fyne.Window) fyne.CanvasObject {

	//config.HoursGet() moved to logon

	Details := widget.NewLabel("")

	larow := widget.NewLabel("Row: ")
	edrow := widget.NewEntry()
	edrow.SetPlaceHolder("Automatically Assigned")

	laid := widget.NewLabel("Hour: ")
	edid := widget.NewSelect([]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}, func(string) {})

	ladesc := widget.NewLabel("Description: ")
	eddesc := widget.NewEntry()

	gridrow := container.New(layout.NewGridLayoutWithRows(2), larow, edrow)
	gridday := container.New(layout.NewGridLayoutWithRows(2), laid, edid)
	griddesc := container.New(layout.NewGridLayoutWithRows(2), ladesc, eddesc)

	saveaddbutton := widget.NewButtonWithIcon("Add Hour Part", theme.ContentCopyIcon(), func() {

		config.HoursAdd(edid.Selected, eddesc.Text)
		config.HoursGet()
	})
	List := widget.NewList(
		func() int {
			return len(config.HoursStore)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {

			//mymessage = config.HoursStore[id].Desc

			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(config.HoursStore[id].Id + " " + config.HoursStore[id].Desc)
		},
	)
	config.FyneDaysList = List
	List.OnSelected = func(id widget.ListItemID) {
		config.SelectedHour = id

		Details.SetText(config.HoursStore[id].Id)

		edrow.SetText(strconv.Itoa(config.HoursStore[id].Row))
		edrow.Disable()

		edid.SetSelected(config.HoursStore[id].Id)

		eddesc.SetText(config.HoursStore[id].Desc)

		deletebutton := widget.NewButtonWithIcon("Delete Hour Part", theme.ContentCopyIcon(), func() {
			myrowhours, _ = strconv.Atoi(edrow.Text)
			config.HoursDelete(myrowhours)
			config.HoursGet()
		})
		savebutton := widget.NewButtonWithIcon("Save Hour Part", theme.ContentCopyIcon(), func() {
			myrowhours, _ = strconv.Atoi(edrow.Text)

			config.HoursUpdate(myrowhours, edid.Selected, eddesc.Text)
			config.HoursGet()

		})
		databox := container.NewVBox(
			deletebutton,
			gridrow,
			gridday,
			griddesc,
			savebutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow("Manage Hour Parts")

		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()
		List.Unselect(id)
	}
	addbutton := widget.NewButtonWithIcon("Add New Hour Part", theme.ContentCopyIcon(), func() {
		databox := container.NewVBox(

			gridrow,
			gridday,
			griddesc,

			saveaddbutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow("Manage Hours")

		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))

		//DetailsBottom := container.NewBorder(databox, nil, nil, nil, nil)dlg.Show()
	})
	topbox := container.NewBorder(addbutton, nil, nil, nil)

	bottombox := container.NewBorder(
		nil,
		Errors,
		nil,
		nil,
		nil,
	)

	return container.NewBorder(
		topbox,
		bottombox,
		nil,
		nil,
		List,
	)

}
