package panes

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/gofrs/uuid/v5"
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

	eduspasswordhash.Disable()

	lacategories := widget.NewLabel(config.GetLangs("userauthcategories"))
	USERcategory := widget.NewCheckGroup([]string{"ALL", "DJAM", "DJPM", "NWS", "CURRENTS", "RECURRENTS", "IMAGINGID", "NEXT", "PROMOS"}, func([]string) {})
	USERcategory.Horizontal = true

	laauthactions := widget.NewLabel(config.GetLangs("userauthactions"))
	edauthactions := widget.NewCheckGroup([]string{"ALL", "Upload/Download", "", "Chart", "Clear", "Traffic"}, func([]string) {})
	edauthactions.Horizontal = true

	gridrow := container.New(layout.NewGridLayoutWithRows(2), larow, edrow)
	gridrole := container.New(layout.NewGridLayoutWithRows(2), larole, edusrole)
	gridpassword := container.New(layout.NewGridLayoutWithRows(2), lapassword, eduspassword)
	gridpasswordhash := container.New(layout.NewGridLayoutWithRows(2), lapasswordhash, eduspasswordhash)

	gridcategories := container.New(layout.NewGridLayoutWithRows(2), lacategories, USERcategory)
	gridauthactions := container.New(layout.NewGridLayoutWithRows(2), laauthactions, edauthactions)
	saveaddbutton := widget.NewButtonWithIcon(config.GetLangs("adduser"), theme.ContentCopyIcon(), func() {
		u, err := uuid.NewV7()
		if err != nil {
			log.Fatalf("failed to generate UUID: %v", err)
		}
		eduspasswordhash.SetText(u.String())
		config.UserAdd(edusrole.Text, eduspassword.Text, eduspasswordhash.Text, USERcategory.Selected, edauthactions.Selected)
		config.UserGet()
		config.FyneUserList.Refresh()
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
			l.SetText(config.ToString(config.UserStore[id.Row].Userauthcategories))
		case 5: // auth
			l.SetText(config.ToString(config.UserStore[id.Row].Userauthaction))
		}
	})
	List.SetColumnWidth(0, 64)
	List.SetColumnWidth(1, 96)
	List.SetColumnWidth(2, 96)
	List.SetColumnWidth(3, 132)
	List.SetColumnWidth(4, 132)

	config.FyneUserList = List
	List.OnSelected = func(id widget.TableCellID) {
		config.SelectedUser = id.Row

		edrow.SetText(strconv.Itoa(config.UserStore[id.Row].Row))
		edrow.Disable()

		edusrole.SetText(config.UserStore[id.Row].Userrole)

		eduspassword.SetText(config.UserStore[id.Row].Userpassword)

		eduspasswordhash.SetText(config.UserStore[id.Row].Userpasswordhash)
		USERcategory.SetSelected(config.UserStore[id.Row].Userauthcategories)
		edauthactions.SetSelected(config.UserStore[id.Row].Userauthaction)
		deletebutton := widget.NewButtonWithIcon(config.GetLangs("eng-usdelete"), theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			config.UserDelete(myrow)
			config.UserGet()
		})
		savebutton := widget.NewButtonWithIcon(config.GetLangs("ussave"), theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)

			config.UserUpdate(myrow, edusrole.Text, eduspassword.Text, eduspasswordhash.Text, USERcategory.Selected, edauthactions.Selected)
			config.UserGet()

		})
		databox := container.NewVBox(
			deletebutton,
			gridrow,
			gridrole,
			gridpassword,
			gridpasswordhash,
			gridcategories,
			gridauthactions,
			savebutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow(config.GetLangs("manuser"))
		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()
		List.Unselect(id)
	}
	addbutton := widget.NewButtonWithIcon(config.GetLangs("adduser"), theme.ContentCopyIcon(), func() {
		databox := container.NewVBox(
			gridrow,
			gridrow,
			gridrole,
			gridpassword,
			gridpasswordhash,
			gridcategories,
			gridauthactions,
			saveaddbutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow(config.GetLangs("manuser"))
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
