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

func UsersScreen(win fyne.Window) fyne.CanvasObject {

	larow := widget.NewLabel("Row: ")
	edrow := widget.NewEntry()
	edrow.SetPlaceHolder(config.GetLangs("autoassigned"))

	larole := widget.NewLabel(config.GetLangs("usrole"))
	edusrole := widget.NewEntry()
	edusrole.SetPlaceHolder(config.GetLangs("usrole"))

	lapassword := widget.NewLabel(config.GetLangs("userpassword"))
	eduspassword := widget.NewEntry()
	eduspassword.SetPlaceHolder(config.GetLangs("userpassword"))

	lapasswordhash := widget.NewLabel(config.GetLangs("userpasswordhash"))
	eduspasswordhash := widget.NewEntry()
	eduspasswordhash.SetPlaceHolder(config.GetLangs("userpasswordhash"))

	lacategories := widget.NewLabel(config.GetLangs("userauthcategories"))
	edcategories := widget.NewCheckGroup(config.CategoriesToArray(), func(string) {})
	edcategories.Horizontal = false

	laauthactions := widget.NewLabel(config.GetLangs("userauthactions"))
	edauthactions := widget.NewCheckGroup([]string{"Upload/Download", "Category History", "Chart", "Clear", "Ad History"}, func(string) {})
	edauthactions.Horizontal = false

	gridrow := container.New(layout.NewGridLayoutWithRows(2), larow, edrow)
	gridrole := container.New(layout.NewGridLayoutWithRows(2), larole, edusrole)
	gridpassword := container.New(layout.NewGridLayoutWithRows(2), lapassword, eduspassword)
	gridpasswordhash := container.New(layout.NewGridLayoutWithRows(2), lapasswordhash, eduspasswordhash)
	gridcategories := container.New(layout.NewGridLayoutWithRows(2), lacategories, edcategories)
	gridauthactions := container.New(layout.NewGridLayoutWithRows(2), laauthactions, edauthactions)
	saveaddbutton := widget.NewButtonWithIcon(config.GetLangs("adduser"), theme.ContentCopyIcon(), func() {

		config.UserAdd(edusrole.SelectedText(), eduspassword.Text, eduspasswordhash.Text, edcategories.Selected, edauthactions.Selected)
		config.DaysGet()
		config.FyneDaysList.Refresh()
	})
	List := widget.NewTable(func() (int, int) {
		return len(config.UserStore), 4
	}, func() fyne.CanvasObject {
		return container.NewMax(widget.NewLabel("template11"), widget.NewIcon(nil))
	}, func(id widget.TableCellID, o fyne.CanvasObject) {
		l := o.(*fyne.Container).Objects[0].(*widget.Label)
		l.Show()
		switch id.Col {

		case 0: // rowid
			l.SetText(strconv.Itoa(config.DaysStore[id.Row].Row))
		case 1: // role
			l.SetText(config.UserStore[id.Row].Userrole)
		case 2: // pass
			l.SetText(config.UserStore[id.Row].Userpassword)
		case 3: // hash
			l.SetText(config.UserStore[id.Row].Userpasswordhash)
		case 4: // cats
			l.SetText(config.UserStore[id.Row].Userauthcategories[1])
		case 5: // auth
			l.SetText(config.UserStore[id.Row].Userauthaction[1])
		}
	})
	List.SetColumnWidth(0, 64)
	List.SetColumnWidth(1, 96)
	List.SetColumnWidth(2, 96)
	List.SetColumnWidth(3, 132)
	List.SetColumnWidth(4, 132)
	List.SetColumnWidth(5, 132)
	config.FyneUserList = List
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
	addbutton := widget.NewButtopasswordhashnWithIcon("Add New Day of Week", theme.ContentCopyIcon(), func() {
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
