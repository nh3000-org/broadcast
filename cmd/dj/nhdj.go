package dj

// Copyright 2012-2023 The NH3000 Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


import (
	"log"
	"os"

	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/nh3000-org/broadcast/cmd/gui/panes"
	"github.com/nh3000-org/broadcast/config"

	"fyne.io/fyne/v2/widget"
)

var TopWindow fyne.Window

type Pane struct {
	Title, Intro string
	Icon         fyne.Resource
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
}

//var Panes = map[string]Pane{}
//var PanesIndex = map[string][]string{}

func main() {

	var a = app.NewWithID("org.nh3000.nh3000")
	config.FyneApp = a
	var w = a.NewWindow("NH3000")
	config.FyneMainWin = w
	config.PreferedLanguage = "eng"
	if strings.HasPrefix(os.Getenv("LANG"), "en") {
		config.PreferedLanguage = "eng"
	}
	if strings.HasPrefix(os.Getenv("LANG"), "sp") {
		config.PreferedLanguage = "spa"
	}
	if strings.HasPrefix(os.Getenv("LANG"), "hn") {
		config.PreferedLanguage = "hin"
	}
	config.PreferedLanguage = config.Decrypt(config.FyneApp.Preferences().StringWithFallback("PreferedLanguage", config.Encrypt(config.PreferedLanguage, config.MySecret)), config.MySecret)
	MyLogo, iconerr := fyne.LoadResourceFromPath("Icon.png")
	if iconerr != nil {
		log.Println("Icon.png error ", iconerr.Error())
	}
	config.Selected = config.Dark
	config.FyneApp.Settings().SetTheme(config.MyTheme{})
	config.FyneApp.SetIcon(MyLogo)

	TopWindow = w
	w.SetMaster()
	config.NewPGSQL()
	intro := widget.NewLabel(config.GetLangs("mn-intro-1") + "\n" + "nats.io" + config.GetLangs("mn-intro-2"))
	intro.Wrapping = fyne.TextWrapWord
	var Panes = map[string]Pane{
		"logon":      {config.GetLangs("ls-title"), "", theme.LoginIcon(), panes.LogonScreen, true},
		"messages":   {config.GetLangs("ms-title"), "", theme.MailSendIcon(), panes.MessagesScreen, true},
		"reports":    {config.GetLangs("rpts"), "", theme.ListIcon(), panes.ReportsScreen, true},
		"inventory":  {config.GetLangs("ra-inv"), "", theme.ListIcon(), panes.InventoryScreen, true},
		"schedule":   {config.GetLangs("ra-sched"), "", theme.ListIcon(), panes.ScheduleScreen, true},
		"categories": {config.GetLangs("ra-cats"), "", theme.ListIcon(), panes.CategoriesScreen, true},
		"hours":      {config.GetLangs("ra-hours"), "", theme.ListIcon(), panes.HoursScreen, true},
		"days":       {config.GetLangs("ra-days"), "", theme.ListIcon(), panes.DaysScreen, true},
		"settings":   {config.GetLangs("ss-title"), "", theme.SettingsIcon(), panes.SettingsScreen, true},
		"password":   {config.GetLangs("ps-title"), "", theme.DocumentIcon(), panes.PasswordScreen, true},
		"encdec":     {config.GetLangs("es-title"), "", theme.CheckButtonIcon(), panes.EncdecScreen, true},
	}

	config.FyneMainWin.SetContent(container.NewAppTabs(
		container.NewTabItemWithIcon(Panes["logon"].Title, Panes["logon"].Icon, panes.LogonScreen(config.FyneMainWin)),
		container.NewTabItemWithIcon(Panes["messages"].Title, Panes["messages"].Icon, panes.MessagesScreen(config.FyneMainWin)),
		container.NewTabItemWithIcon(Panes["reports"].Title, Panes["reports"].Icon, panes.ReportsScreen(config.FyneMainWin)),
		container.NewTabItemWithIcon(Panes["inventory"].Title, Panes["inventory"].Icon, panes.InventoryScreen(config.FyneMainWin)),
		container.NewTabItemWithIcon(Panes["schedule"].Title, Panes["schedule"].Icon, panes.ScheduleScreen(config.FyneMainWin)),
		container.NewTabItemWithIcon(Panes["categories"].Title, Panes["categories"].Icon, panes.CategoriesScreen(config.FyneMainWin)),
		container.NewTabItemWithIcon(Panes["hours"].Title, Panes["hours"].Icon, panes.HoursScreen(config.FyneMainWin)),
		container.NewTabItemWithIcon(Panes["days"].Title, Panes["days"].Icon, panes.DaysScreen(config.FyneMainWin)),
		container.NewTabItemWithIcon(Panes["encdec"].Title, Panes["encdec"].Icon, panes.EncdecScreen(config.FyneMainWin)),
		container.NewTabItemWithIcon(Panes["settings"].Title, Panes["settings"].Icon, panes.SettingsScreen(config.FyneMainWin)),
		container.NewTabItemWithIcon(Panes["password"].Title, Panes["password"].Icon, panes.PasswordScreen(config.FyneMainWin)),
	))

	config.FyneMainWin.Resize(fyne.NewSize(640, 480))
	config.FyneMainWin.ShowAndRun()
}


