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

func CategoriesScreen(win fyne.Window) fyne.CanvasObject {
	st := fyne.TextStyle{
		Monospace: true,
	}
	config.FyneApp.Settings().Theme().Font(st)

	larow := widget.NewLabel("Row: ")
	edrow := widget.NewEntry()
	edrow.SetPlaceHolder("Automatically Assigned")

	laid := widget.NewLabel("Category: ")
	edid := widget.NewEntry()

	ladesc := widget.NewLabel("Description: ")
	eddesc := widget.NewEntry()

	gridrow := container.New(layout.NewGridLayoutWithRows(2), larow, edrow)
	gridday := container.New(layout.NewGridLayoutWithRows(2), laid, edid)
	griddesc := container.New(layout.NewGridLayoutWithRows(2), ladesc, eddesc)
	stubbutton := widget.NewButtonWithIcon("Create STUB of Categories", theme.ContentCopyIcon(), func() {
		var where = config.CategoriesWriteStub(false)
		Errors.SetText(where)
		config.CategoriesGet()
		config.FyneCategoryList.Refresh()
		config.Send("messages.Export", where, config.NatsAlias)

	})
	saveaddbutton := widget.NewButtonWithIcon("Add Category", theme.ContentCopyIcon(), func() {

		config.CategoriesAdd(edid.Text, eddesc.Text)
		config.CategoriesGet()
		config.FyneCategoryList.Refresh()
	})

	List := widget.NewTable(func() (int, int) {
		return len(config.CategoriesStore), 3
	}, func() fyne.CanvasObject {
		return container.NewMax(widget.NewLabel("template11"), widget.NewIcon(nil))
	}, func(id widget.TableCellID, o fyne.CanvasObject) {
		l := o.(*fyne.Container).Objects[0].(*widget.Label)
		l.Show()
		switch id.Col {

		case 0: // rowid
			l.SetText(strconv.Itoa(config.CategoriesStore[id.Row].Row))
		case 1: // dats
			l.SetText(config.CategoriesStore[id.Row].Id)
		case 2: // hour
			l.SetText(config.CategoriesStore[id.Row].Desc)
		}
	})
	List.SetColumnWidth(0, 64)
	List.SetColumnWidth(1, 132)
	List.SetColumnWidth(2, 132)

	config.FyneCategoryList = List
	List.OnSelected = func(id widget.TableCellID) {
		config.SelectedCategory = id.Row

		edrow.SetText(strconv.Itoa(config.CategoriesStore[id.Row].Row))
		edrow.Disable()

		edid.SetText(config.CategoriesStore[id.Row].Id)

		eddesc.SetText(config.CategoriesStore[id.Row].Desc)

		deletebutton := widget.NewButtonWithIcon("Delete Inventory Category", theme.ContentCopyIcon(), func() {
			myrowcat, _ := strconv.Atoi(edrow.Text)
			if config.CategoriesWhereUsed(edid.Text) != 0 {
				config.CategoriesDelete(myrowcat)
				config.CategoriesGet()
			}
		})
		savebutton := widget.NewButtonWithIcon("Save Inventory Category", theme.ContentCopyIcon(), func() {
			myrowcat, _ := strconv.Atoi(edrow.Text)

			config.CategoriesUpdate(myrowcat, edid.Text, eddesc.Text)
			config.CategoriesGet()

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
		dlg := fyne.CurrentApp().NewWindow("Manage Inventory Category")

		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()
		List.Unselect(id)
	}

	addbutton := widget.NewButtonWithIcon("Add New Inventory Category", theme.ContentCopyIcon(), func() {
		databox := container.NewVBox(

			gridrow,
			gridday,
			griddesc,

			saveaddbutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow("Manage Inventory Category")
		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()
	})
	topbox := container.NewBorder(addbutton, nil, nil, stubbutton)

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
