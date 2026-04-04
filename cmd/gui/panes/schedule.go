package panes

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/nh3000-org/broadcast/config"
)

func ScheduleScreen(win fyne.Window) fyne.CanvasObject {

	laseld := widget.NewLabel("Day:")
	laselday := widget.NewSelect([]string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN", "VID"}, func(string) {})

	laselh := widget.NewLabel("Hour:")
	laselhour := widget.NewSelect([]string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}, func(string) {})

	selebutton := widget.NewButtonWithIcon("Search", theme.SearchIcon(), func() {

		config.ScheduleSel(laselday.Selected, laselhour.Selected)
		config.FyneScheduleList.Refresh()
	})
	gridsel := container.New(layout.NewGridLayoutWithColumns(5), laseld, laselday, laselh, laselhour, selebutton)

	larow := widget.NewLabel("Row: ")
	edrow := widget.NewEntry()
	edrow.SetPlaceHolder("Automatically Assigned")

	laday := widget.NewLabel("Day: ")
	edday := widget.NewRadioGroup([]string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN", "VID"}, func(string) {})
	edday.Horizontal = true

	lahour := widget.NewLabel("Hour: ")
	edhour := widget.NewSelect([]string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}, func(string) {})

	lapos := widget.NewLabel("Position on Schedule: ")
	edpos := widget.NewSelect([]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21"}, func(string) {})

	lacategory := widget.NewLabel("Category to Pick From: ")
	//edcategory := widget.NewSelect(config.CategoriesToArray(), func(string) {})

	laspins := widget.NewLabel("Spins to Play From Category: ")
	edspins := widget.NewSelect([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}, func(string) {})

	lacpf := widget.NewLabel("From Day: ")
	edcpf := widget.NewSelect([]string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}, func(string) {})

	lacpt := widget.NewLabel("To Day: ")
	edcpt := widget.NewSelect([]string{"*ALL", "MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}, func(string) {})
	lacpfh := widget.NewLabel("From Hour: ")
	edcpfh := widget.NewSelect([]string{"*ALL", "00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}, func(string) {})
	lacpth := widget.NewLabel("To Hour: ")
	edcpth := widget.NewSelect([]string{"*ALL", "00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}, func(string) {})

	var cpyerr = false
	copybutton := widget.NewButtonWithIcon("Copy ", theme.ContentCopyIcon(), func() {
		if edcpf.Selected == "" {
			cpyerr = true
		}
		if edcpt.Selected == "" {
			cpyerr = true
		}
		if edcpfh.Selected == "" {
			edcpth.SetSelected("*ALL")
		}
		if edcpth.Selected == "" {
			edcpfh.SetSelected("*ALL")
		}
		if edcpfh.Selected != "*ALL" && edcpth.Selected != "*ALL" {
			if edcpfh.Selected == edcpth.Selected {
				cpyerr = true
			}
		}
		if edcpfh.Selected == "*ALL" && edcpth.Selected != "*ALL" {

			cpyerr = true

		}
		if edcpfh.Selected != "*ALL" && edcpth.Selected == "*ALL" {

			cpyerr = true

		}

		if edcpf.Selected == edcpt.Selected {
			cpyerr = true
		}
		if !cpyerr {
			config.ScheduleCopy(edcpf.Selected, edcpt.Selected, edcpfh.Selected, edcpth.Selected)
			config.ScheduleSel(laselday.Selected, laselhour.Selected)
		}
	})
	gridcopy := container.New(layout.NewGridLayoutWithColumns(4), lacpf, edcpf, lacpt, edcpt, lacpfh, edcpfh, lacpth, edcpth, copybutton)
	gridrow := container.New(layout.NewGridLayoutWithRows(2), larow, edrow)
	gridday := container.New(layout.NewGridLayoutWithRows(2), laday, edday)
	gridhour := container.New(layout.NewGridLayoutWithRows(2), lahour, edhour)
	gridpos := container.New(layout.NewGridLayoutWithRows(2), lapos, edpos)
	gridcat := container.New(layout.NewGridLayoutWithRows(2), lacategory, ED2category)
	gridspins := container.New(layout.NewGridLayoutWithRows(2), laspins, edspins)
	saveaddbutton := widget.NewButtonWithIcon("Add Schedule Item", theme.ContentCopyIcon(), func() {
		myspins, _ := strconv.Atoi(edspins.Selected)

		config.ScheduleAdd(edday.Selected, edhour.Selected, edpos.Selected, ED2category.Selected, myspins)
		config.ScheduleSel(laselday.Selected, laselhour.Selected)
		config.FyneScheduleList.Refresh()
	})

	List := widget.NewTable(func() (int, int) {
		return len(config.ScheduleStore), 6
	}, func() fyne.CanvasObject {
		return container.NewMax(widget.NewLabel("template11"), widget.NewIcon(nil))
	}, func(id widget.TableCellID, o fyne.CanvasObject) {
		l := o.(*fyne.Container).Objects[0].(*widget.Label)
		l.Show()
		switch id.Col {

		case 0: // rowid
			l.SetText(strconv.Itoa(config.ScheduleStore[id.Row].Row))
		case 1: // dats
			l.SetText(config.ScheduleStore[id.Row].Days)
		case 2: // hour
			l.SetText(config.ScheduleStore[id.Row].Hours)
		case 3: // position
			l.SetText(config.ScheduleStore[id.Row].Position)
		case 4: // category
			l.SetText(config.ScheduleStore[id.Row].Category)
		case 5: // spins
			l.SetText(strconv.Itoa(config.ScheduleStore[id.Row].Spinstoplay))
		}
	})
	List.SetColumnWidth(0, 64)
	List.SetColumnWidth(1, 64)
	List.SetColumnWidth(2, 64)
	List.SetColumnWidth(3, 64)
	List.SetColumnWidth(4, 256)
	List.SetColumnWidth(5, 64)
	config.FyneScheduleList = List
	List.OnSelected = func(id widget.TableCellID) {
		config.SelectedDay = id.Row
		edrow.SetText(strconv.Itoa(config.ScheduleStore[id.Row].Row))
		edday.SetSelected(config.ScheduleStore[id.Row].Days)
		edhour.SetSelected(config.ScheduleStore[id.Row].Hours)
		edpos.SetSelected(config.ScheduleStore[id.Row].Position)
		ED2category.SetSelected(config.ScheduleStore[id.Row].Category)
		edspins.SetSelected(strconv.Itoa(config.ScheduleStore[id.Row].Spinstoplay))
		edrow.SetText(strconv.Itoa(config.ScheduleStore[id.Row].Row))
		edrow.Disable()
		deletebutton := widget.NewButtonWithIcon("Delete Schedule Item", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			config.ScheduleDelete(myrow)
			config.ScheduleSel(laselday.Selected, laselhour.Selected)
			config.FyneScheduleList.Refresh()
		})
		savebutton := widget.NewButtonWithIcon("Save Schedule", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			myspins, _ := strconv.Atoi(edspins.Selected)
			config.ScheduleUpdate(myrow, edday.Selected, edhour.Selected, edpos.Selected, ED2category.Selected, myspins)
			config.ScheduleSel(laselday.Selected, laselhour.Selected)
			config.FyneScheduleList.Refresh()
		})
		gridrow := container.New(layout.NewGridLayoutWithRows(2), larow, edrow)

		databox := container.NewVBox(
			deletebutton,
			gridrow,
			gridday,
			gridhour,
			gridpos,
			gridcat,
			gridspins,
			savebutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow("Manage Schedule")

		//DetailsBottom := container.NewBorder(databox, nil, nil, nil, nil)
		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()
		config.ScheduleSel(laselday.Selected, laselhour.Selected)
		config.FyneScheduleList.Refresh()
		List.Unselect(id)
	}
	addbutton := widget.NewButtonWithIcon("Add New Schedule Item", theme.ContentCopyIcon(), func() {
		databox := container.NewVBox(

			gridrow,
			gridday,
			gridhour,
			gridpos,
			gridcat,
			gridspins,
			saveaddbutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow("Manage Schedule Item")

		//DetailsBottom := container.NewBorder(databox, nil, nil, nil, nil)
		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()
		config.ScheduleSel(laselday.Selected, laselhour.Selected)
		config.FyneScheduleList.Refresh()
	})
	topbox := container.NewBorder(addbutton, gridcopy, nil, nil)

	bottombox := container.NewBorder(
		gridsel,
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
