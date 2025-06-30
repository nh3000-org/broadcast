package panes

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/nh3000-org/broadcast/config"
	//"github.com/nh3000-org/radio/config"
)

var myrowhours int

func HoursScreen(win fyne.Window) fyne.CanvasObject {

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

	List := widget.NewTable(func() (int, int) {
		return len(config.HoursStore), 3
	}, func() fyne.CanvasObject {
		return container.NewMax(widget.NewLabel("template11"), widget.NewIcon(nil))
	}, func(id widget.TableCellID, o fyne.CanvasObject) {
		l := o.(*fyne.Container).Objects[0].(*widget.Label)
		l.Show()
		switch id.Col {

		case 0: // rowid
			l.SetText(strconv.Itoa(config.HoursStore[id.Row].Row))
		case 1: // dats
			l.SetText(config.HoursStore[id.Row].Id)
		case 2: // hour
			l.SetText(config.HoursStore[id.Row].Desc)
		}
	})
	List.SetColumnWidth(0, 64)
	List.SetColumnWidth(1, 132)
	List.SetColumnWidth(2, 132)

	config.FyneDaysList = List
	List.OnSelected = func(id widget.TableCellID) {
		config.SelectedHour = id.Row

		edrow.SetText(strconv.Itoa(config.HoursStore[id.Row].Row))
		edrow.Disable()

		edid.SetSelected(config.HoursStore[id.Row].Id)

		eddesc.SetText(config.HoursStore[id.Row].Desc)

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
