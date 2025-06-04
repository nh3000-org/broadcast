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

//var myrow int
//var mydow int

func DaysScreen(win fyne.Window) fyne.CanvasObject {

	larow := widget.NewLabel("Row: ")
	edrow := widget.NewEntry()
	edrow.SetPlaceHolder("Automatically Assigned")

	laday := widget.NewLabel("Day: ")
	edday := widget.NewRadioGroup([]string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN", "VID"}, func(string) {})
	edday.Horizontal = true
	ladesc := widget.NewLabel("Description: ")
	eddesc := widget.NewEntry()
	ladow := widget.NewLabel("Day of Week: ")
	eddow := widget.NewRadioGroup([]string{"1", "2", "3", "4", "5", "6", "7"}, func(string) {})
	eddow.Horizontal = true

	gridrow := container.New(layout.NewGridLayoutWithRows(2), larow, edrow)
	gridday := container.New(layout.NewGridLayoutWithRows(2), laday, edday)
	griddesc := container.New(layout.NewGridLayoutWithRows(2), ladesc, eddesc)
	griddow := container.New(layout.NewGridLayoutWithRows(2), ladow, eddow)
	saveaddbutton := widget.NewButtonWithIcon("Add Day of Week", theme.ContentCopyIcon(), func() {
		mydow, _ := strconv.Atoi(eddow.Selected)

		config.DaysAdd(edday.Selected, eddesc.Text, mydow)
		config.DaysGet()
		config.FyneDaysList.Refresh()
	})
	List := widget.NewTable(func() (int, int) {
		return len(config.DaysStore), 4
	}, func() fyne.CanvasObject {
		return container.NewMax(widget.NewLabel("template11"), widget.NewIcon(nil))
	}, func(id widget.TableCellID, o fyne.CanvasObject) {
		l := o.(*fyne.Container).Objects[0].(*widget.Label)
		l.Show()
		switch id.Col {

		case 0: // rowid
			l.SetText(strconv.Itoa(config.DaysStore[id.Row].Row))
		case 1: // dats
			l.SetText(config.DaysStore[id.Row].Day)
		case 2: // hour
			l.SetText(config.DaysStore[id.Row].Desc)
		case 3: // rowid
			l.SetText(strconv.Itoa(config.DaysStore[id.Row].Dow))
		}
	})
	List.SetColumnWidth(0, 64)
	List.SetColumnWidth(1, 132)
	List.SetColumnWidth(2, 132)
	config.FyneDaysList = List
	List.OnSelected = func(id widget.TableCellID) {
		config.SelectedDay = id.Row

		edrow.SetText(strconv.Itoa(config.DaysStore[id.Row].Row))
		edrow.Disable()

		edday.SetSelected(config.DaysStore[id.Row].Day)

		eddesc.SetText(config.DaysStore[id.Row].Desc)

		eddow.SetSelected(strconv.Itoa(config.DaysStore[id.Row].Dow))

		deletebutton := widget.NewButtonWithIcon("Delete Day of Week", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			config.DaysDelete(myrow)
			config.DaysGet()
		})
		savebutton := widget.NewButtonWithIcon("Save Day of Week", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			mydow, _ := strconv.Atoi(eddow.Selected)
			config.DaysUpdate(myrow, edday.Selected, eddesc.Text, mydow)
			config.DaysGet()

		})
		databox := container.NewVBox(
			deletebutton,
			gridrow,
			gridday,
			griddesc,
			griddow,
			savebutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow("Manage Days")
		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()
		List.Unselect(id)
	}
	addbutton := widget.NewButtonWithIcon("Add New Day of Week", theme.ContentCopyIcon(), func() {
		databox := container.NewVBox(
			gridrow,
			gridday,
			griddesc,
			griddow,
			saveaddbutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow("Manage Days")
		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()
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
