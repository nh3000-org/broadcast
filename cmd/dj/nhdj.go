package main

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
	"context"
	"encoding/json"
	"fmt"
	tc "image/color"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"golang.org/x/crypto/bcrypt"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/nh3000-org/broadcast/config"

	"fyne.io/fyne/v2/widget"
)

var TopWindow fyne.Window

type DJ struct {
	Artist              string
	Album               string
	Song                string
	Length              string
	SchedRow            string
	SchedDay            string
	SchedHour           string
	SchedPosition       string
	SchedCategory       string
	SchedSpinsToPlay    string
	SchedSpinsLefToPlay string
}

var DJJSON = DJ{}
var memoryStats runtime.MemStats
var ctxmain context.Context
var ctxmaincan context.CancelFunc
var a fyne.App
var w fyne.Window
var onairmp3 jetstream.KeyValue

func main() {

	a = app.NewWithID("org.nh3000.nh3000")
	w = a.NewWindow("NH3000")
	config.FyneApp = a
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
	//config.PreferedLanguage = config.Decrypt(config.FyneApp.Preferences().StringWithFallback("PreferedLanguage", config.Encrypt(config.PreferedLanguage, config.MySecret)), config.MySecret)
	MyLogo, iconerr := fyne.LoadResourceFromPath("Icon.png")
	if iconerr != nil {
		log.Println("Icon.png error ", iconerr.Error())
	}
	config.Selected = config.Dark
	a.Settings().SetTheme(config.MyTheme{})
	a.SetIcon(MyLogo)

	TopWindow = w
	w.SetMaster()
	logLifecycle()

	natserr := config.NewNatsJS()
	if natserr != nil {
		log.Fatal("Could not connect to NATS ")
	}

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder(config.GetLangs("ls-password"))
	TPbutton := widget.NewButtonWithIcon(config.GetLangs("ls-trypass"), theme.LoginIcon(), func() {

		var iserrors = false
		ph, _ := config.LoadHashWithDefault("config.hash", "123456")

		//log.Println("pw ", MyPrefs.Password)
		pwh, err := bcrypt.GenerateFromPassword([]byte(password.Text), bcrypt.DefaultCost)
		config.PasswordHash = string(pwh)
		if err != nil {
			iserrors = true

			log.Println("pw cant gen hash")
		}

		// Comparing the password with the hash
		errpw := bcrypt.CompareHashAndPassword([]byte(ph), []byte(password.Text))
		if errpw != nil {
			iserrors = true

			log.Println("pw bad hash ", errpw, "ph", ph, "pt", password.Text)
		}
		if !iserrors {
			readPreferences()
			config.NewNatsJS()
			config.NewPGSQL()
			ctxmain, ctxmaincan = context.WithCancel(context.Background())
			mp3msg, mp3err := config.NATS.OnAirmp3.Watch(ctxmain, "OnAirmp3")

			if mp3err != nil {
				log.Println("ReceiveONAIRMP3", mp3err)
				config.Send("messages."+"DJ", "Receive On Air mp3 ", "DJ")
			}
			//log.Println("ReceiveONAIRMP3 waiting")

			for {
				runtime.GC()
				runtime.ReadMemStats(&memoryStats)
				kve := <-mp3msg.Updates()
				//log.Println("ReceiveONAIRMP3", kve)
				if kve != nil {
					errum := json.Unmarshal(kve.Value(), &DJJSON)
					if errum != nil {
						log.Println("DJ ReceiveONAIRMP3", errum)
						config.Send("messages."+"DJ", "DJ Receive On Air mp3 ", errum.Error())

					}

					if w != nil {
						w.SetTitle("On Air MP3 " + DJJSON.Artist + " - " + DJJSON.Song + " - " + DJJSON.Album + " " + strconv.FormatUint(memoryStats.Alloc/1024/1024, 10) + " Mib")
					}
					drawgGui(DJJSON)

				}
			}

		}

	})
	vertbox := container.NewVBox(

		widget.NewLabelWithStyle("DJ Console", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		password,
		TPbutton,

		widget.NewLabel(""),
		//		themes,

	)

	w.SetContent(vertbox)

	w.Resize(fyne.NewSize(640, 480))
	w.ShowAndRun()

}
func drawgGui(oa DJ) {
	progress := widget.NewProgressBar()
	progress.Min = 0
	m, _ := strconv.ParseFloat(oa.Length, 64)
	progress.Max = m
	green := tc.RGBA{0, 255, 0, 255}
	//blue := tc.RGBA{0, 0, 255, 255}
	go func() {
		for i := 1.0; i <= m; i++ {
			time.Sleep(time.Second)
			progress.SetValue(i)
		}
	}()
	artisth := canvas.NewText("Artist", tc.White)
	songh := canvas.NewText("Song", tc.White)
	albumh := canvas.NewText("Album", tc.White)
	lengthh := canvas.NewText("Length", tc.White)

	artist := canvas.NewText(oa.Artist, green)
	song := canvas.NewText(oa.Song, green)
	album := canvas.NewText(oa.Album, green)
	length := canvas.NewText(oa.Length, green)

	asalgridhead := container.New(layout.NewGridLayout(4), artisth, songh, albumh, lengthh)
	asalgrid := container.New(layout.NewGridLayout(4), artist, song, album, length)

	categoryh := canvas.NewText("Category", tc.White)
	dayh := canvas.NewText("Day", tc.White)
	hourh := canvas.NewText("Hour", tc.White)
	positionh := canvas.NewText("Position", tc.White)

	category := canvas.NewText(oa.SchedCategory, green)
	day := canvas.NewText(oa.SchedDay, green)
	hour := canvas.NewText(oa.SchedHour, green)
	position := canvas.NewText(oa.SchedPosition, green)
	cdhpgridhead := container.New(layout.NewGridLayout(4), categoryh, dayh, hourh, positionh)
	cdhpgrid := container.New(layout.NewGridLayout(4), category, day, hour, position)

	stph := canvas.NewText("Spins To Play", tc.White)
	sltph := canvas.NewText("Spins Left To Play", tc.White)
	stp := canvas.NewText(oa.SchedSpinsToPlay, green)
	sltp := canvas.NewText(oa.SchedSpinsLefToPlay, green)
	ssgridhead := container.New(layout.NewGridLayout(2), stph, sltph)
	ssgrid := container.New(layout.NewGridLayout(2), stp, sltp)
	comminguph := canvas.NewText("Comming Up", tc.White)
	commingup := container.New(layout.NewGridLayout(1), comminguph)
	vertbox := container.NewVBox(
		widget.NewLabel(" "),
		asalgridhead,
		asalgrid,
		widget.NewLabel(" "),
		cdhpgridhead,
		cdhpgrid,
		widget.NewLabel(" "),
		ssgridhead,
		ssgrid,
		widget.NewLabel(" "),
		container.NewVBox(progress),
		widget.NewLabel(" "),
		comminguph,
		commingup,
	)

	w.SetContent(vertbox)

}

var PreferencesLocation = "/home/oem/.config/fyne/org.nh3000.nh3000/preferences.json"

const MySecret string = "abd&1*~#^2^#s0^=)^^7%c34"

func readPreferences() {
	// read config preferences.json
	jsondata, readerr := os.ReadFile(PreferencesLocation)
	if readerr != nil {
		log.Println("ERROR Preferences readerr ", readerr)
	}
	// parse json
	var cfg map[string]any
	errunmarshal := json.Unmarshal(jsondata, &cfg)
	if errunmarshal != nil {
		log.Println("ERROR Preferences errunmarshal ", errunmarshal)
	}

	config.DBpassword = config.Decrypt(fmt.Sprintf("%v", cfg["DBPASSWORD"]), MySecret)

	config.DBaddress = config.Decrypt(fmt.Sprintf("%v", cfg["DBADDRESS"]), MySecret)
	//log.Println(config.DBaddress)

	config.DBuser = config.Decrypt(fmt.Sprintf("%v", cfg["DBUSER"]), MySecret)

	config.NatsCaroot = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCaroot"]), MySecret)
	config.NatsClientkey = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCakey"]), MySecret)
	config.NatsClientcert = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCaclient"]), MySecret)
	config.NatsQueuePassword = config.Decrypt(fmt.Sprintf("%v", cfg["NatsQueuePassword"]), MySecret)
	//amm := strconv.Itoa(cfg["AdsMaxMinutes"])

	//log.Println("CONFIG AdsMaxMinutes", config.AdsMaxMinutes)
	//log.Println("NATS AUTH user", config.NatsServer, config.NatsUser, config.NatsUserPassword)
	config.NewNatsJS()
	config.NewPGSQL()
}
func logLifecycle() {

	a.Lifecycle().SetOnStopped(func() {
		if config.LoggedOn {
			//config.Send("messages."+config.NatsAlias, config.GetLangs("ls-dis"), config.NatsAlias)
			ctxmaincan()

			//config.DevCancel = true

			//config.DeleteConsumer("MESSAGES", "messages")
			//config.DeleteConsumer("DEVICES", "devices")
		}
	})

}
