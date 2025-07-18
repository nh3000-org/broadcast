package panes

import (
	"bytes"

	"github.com/nh3000-org/broadcast/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//var mymessage = ""
//var mymessageshort = ""

// var mymessagedlg []byte

var Errors = widget.NewLabel("...")

func MessagesScreen(win fyne.Window) fyne.CanvasObject {
	var selectedms int
	var selectedseq uint64
	var selecteduuid string
	Errors := widget.NewLabel("...")
	config.FyneMessageWin = win
	message := widget.NewMultiLineEntry()
	message.SetPlaceHolder(config.GetLangs("ms-mm"))
	message.SetMinRowsVisible(2)

	Details := widget.NewLabel("")
	var DetailsBorder = container.NewBorder(Details, nil, nil, nil, nil)

	DetailsVW := container.NewScroll(DetailsBorder)
	DetailsVW.SetMinSize(fyne.NewSize(300, 240))

	cpybutton := widget.NewButtonWithIcon(config.GetLangs("ms-cpy"), theme.ContentCopyIcon(), func() {
		win.Clipboard().SetContent(Details.Text)
	})
	delbutton := widget.NewButtonWithIcon(config.GetLangs("ms-del"), theme.ContentCopyIcon(), func() {
		config.DeleteNatsMessage(selectedseq)
		delete(config.NatsMessages, selectedms)
		delete(config.NatsMessagesIndex, selecteduuid)
		config.FyneMessageList.Refresh()
	})
	/* 	List := widget.NewList(
	   		func() int {
	   			return len(config.NatsMessages)
	   		},
	   		func() fyne.CanvasObject {
	   			return container.NewHBox(widget.NewLabel("Template Object"))
	   		},
	   		func(id widget.ListItemID, item fyne.CanvasObject) {

	   			mymessage := config.NatsMessages[id].MSmessage
	   			if len(config.NatsMessages[id].MSmessage) > 100 {
	   				mymessageshort := strings.ReplaceAll(config.NatsMessages[id].MSmessage, "\n", ".")
	   				mymessage = mymessageshort[0:100]
	   			}
	   			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(config.NatsMessages[id].MSsubject + "." + mymessage)
	   		},
	   	)
	*/
	List := widget.NewTable(func() (int, int) {
		return len(config.NatsMessages), 3
	}, func() fyne.CanvasObject {
		return container.NewMax(widget.NewLabel("template11"), widget.NewIcon(nil))
	}, func(id widget.TableCellID, o fyne.CanvasObject) {
		l := o.(*fyne.Container).Objects[0].(*widget.Label)
		l.Show()
		switch id.Col {

		case 0: // subject
			l.SetText(config.NatsMessages[id.Row].MSsubject)
		case 1: // message
			l.SetText(config.NatsMessages[id.Row].MSmessage)
		case 2: // hostname
			l.SetText(config.NatsMessages[id.Row].MShostname)

		}
	})
	List.SetColumnWidth(0, 132)
	List.SetColumnWidth(1, 256)
	List.SetColumnWidth(2, 132)

	config.FyneMessageList = List
	List.OnSelected = func(id widget.TableCellID) {
		selectedms = id.Row
		selectedseq = config.NatsMessages[id.Row].MSsequence
		selecteduuid = config.NatsMessages[id.Row].MSnodeuuid
		Details.SetText(config.NatsMessages[id.Row].MSmessage + "\n." + "\n" + config.NatsMessages[id.Row].MSsubject + " " + config.NatsMessages[id.Row].MSos + " " + config.NatsMessages[id.Row].MShostname + " on " + config.NatsMessages[id.Row].MSdate + "\n" + config.NatsMessages[id.Row].MSipadrs + "\n" + config.NatsMessages[id.Row].MSmacid + "\nNode ID: " + config.NatsMessages[id.Row].MSnodeuuid + "\nMsg ID:" + config.NatsMessages[id.Row].MSiduuid)
		dlg := fyne.CurrentApp().NewWindow(msg2dlg(id.Row))
		DetailsVW := container.NewScroll(DetailsBorder)
		DetailsVW.SetMinSize(fyne.NewSize(300, 240))
		DetailsBottom := container.NewBorder(cpybutton, delbutton, nil, nil, nil)
		dlg.SetContent(container.NewBorder(DetailsVW, DetailsBottom, nil, nil, nil))
		dlg.Show()
		List.Unselect(id)
	}
	smbutton := widget.NewButtonWithIcon("", theme.MailSendIcon(), func() {
		if !config.LoggedOn {
			Errors.SetText(config.GetLangs("cs-lf"))
		}
		config.Send("messages."+config.NatsAlias, message.Text, config.NatsAlias)
		message.SetText("")
	})
	topbox := container.NewHSplit(
		message,
		smbutton,
	)
	topbox.SetOffset(.95)
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
func msg2dlg(id int) string {
	//config.NatsMessages[id].MSmessage + "\n.................." + "\n" + config.NatsMessages[id].MSsubject + "\n" + config.NatsMessages[id].MSos + "\n" + config.NatsMessages[id].MShostname + "\n" + config.NatsMessages[id].MSipadrs + "\n" + config.NatsMessages[id].MSnodeuuid + "\n" + config.NatsMessages[id].MSiduuid + "\n" + config.NatsMessages[id].MSdate
	buf := &bytes.Buffer{}
	buf.WriteString(config.NatsMessages[id].MSmessage)
	buf.WriteString("\n[")

	buf.WriteString(config.NatsMessages[id].MSsubject)
	buf.WriteString("] from ")
	buf.WriteString(config.NatsMessages[id].MSalias)
	buf.WriteString(" on ")
	buf.WriteString(config.NatsMessages[id].MSos)
	buf.WriteString(" host ")
	buf.WriteString(config.NatsMessages[id].MShostname)
	buf.WriteString("\n[net] ")
	buf.WriteString(config.NatsMessages[id].MSipadrs)
	buf.WriteString("\n[mac] ")
	buf.WriteString(config.NatsMessages[id].MSmacid)
	buf.WriteString("\n[msgid] ")
	buf.WriteString(config.NatsMessages[id].MSiduuid)
	buf.WriteString("\n[node] ")
	buf.WriteString(config.NatsMessages[id].MSnodeuuid)
	buf.WriteString("\n[on] ")
	buf.WriteString(config.NatsMessages[id].MSdate)
	return buf.String()
}
