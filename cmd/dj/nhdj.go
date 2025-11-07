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

var a = app.NewWithID("org.nh3000.nh3000")
var w = a.NewWindow("NH3000")

var DJJSON = DJ{}
var memoryStats runtime.MemStats
var ctxmain context.Context
var ctxmaincan context.CancelFunc

//var a fyne.App
//var w fyne.Window

// var onairmp3 jetstream.KeyValue
var errum error
var mp3msg jetstream.KeyWatcher
var mp3err error
var wavmsg jetstream.KeyWatcher
var waverr error
var kve jetstream.KeyValueEntry

func main() {

	//a = app.NewWithID("org.nh3000.nh3000")
	//w = a.NewWindow("NH3000")
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
			if config.NatsBucketType == "mp3" {
				mp3msg, mp3err = config.NATS.OnAirmp3.Watch(ctxmain, "OnAirmp3")
				if mp3err != nil {
					log.Println("ReceiveONAIRMP3", mp3err)
					config.Send("messages."+"DJ", "Receive On Air mp3 ", "DJ")
				}
				for {

					kve = <-mp3msg.Updates()
					//log.Println("ReceiveONAIRMP3", kve)
					if kve != nil {
						errum = json.Unmarshal(kve.Value(), &DJJSON)
						if errum != nil {
							log.Println("DJ ReceiveONAIRMP3", errum)
							config.Send("messages."+"DJ", "DJ Receive On Air mp3 ", errum.Error())

						}
						runtime.GC()
						runtime.ReadMemStats(&memoryStats)
						if w != nil {
							w.SetTitle("On Air MP3 " + DJJSON.Artist + " - " + DJJSON.Song + " - " + DJJSON.Album + " " + strconv.FormatUint(memoryStats.Alloc/1024/1024, 10) + " Mib")
						}
						drawgGui(DJJSON)

					}
				}
			}
			if config.NatsBucketType == "wav" {
				wavmsg, waverr = config.NATS.OnAirwav.Watch(ctxmain, "OnAirwav")
				if waverr != nil {
					log.Println("ReceiveONAIRWAV", waverr)
					config.Send("messages."+"DJ", "Receive On Air wav ", "DJ")
				}
				for {

					kve = <-wavmsg.Updates()
					//log.Println("ReceiveONAIRMP3", kve)
					if kve != nil {
						errum = json.Unmarshal(kve.Value(), &DJJSON)
						if errum != nil {
							log.Println("DJ ReceiveONAIRMP3", errum)
							config.Send("messages."+"DJ", "DJ Receive On Air wav ", errum.Error())

						}
						runtime.GC()
						runtime.ReadMemStats(&memoryStats)
						if w != nil {
							w.SetTitle("On Air WAV " + DJJSON.Artist + " - " + DJJSON.Song + " - " + DJJSON.Album + " " + strconv.FormatUint(memoryStats.Alloc/1024/1024, 10) + " Mib")
						}
						drawgGui(DJJSON)

					}
				}
			}
			//log.Println("ReceiveONAIRMP3 waiting")
			/*
				for {

					kve = <-mp3msg.Updates()
					//log.Println("ReceiveONAIRMP3", kve)
					if kve != nil {
						errum = json.Unmarshal(kve.Value(), &DJJSON)
						if errum != nil {
							log.Println("DJ ReceiveONAIRMP3", errum)
							config.Send("messages."+"DJ", "DJ Receive On Air mp3 ", errum.Error())

						}
						runtime.GC()
						runtime.ReadMemStats(&memoryStats)
						if w != nil {
							w.SetTitle("On Air MP3 " + DJJSON.Artist + " - " + DJJSON.Song + " - " + DJJSON.Album + " " + strconv.FormatUint(memoryStats.Alloc/1024/1024, 10) + " Mib")
						}
						drawgGui(DJJSON)

					}
				} */

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

var tl float64
var ttp float64
var green = tc.RGBA{0, 255, 0, 255}

var progress = widget.NewProgressBar()
var progresgrid = container.New(layout.NewGridLayout(1), progress)
var artisth = canvas.NewText("Artist", tc.White)
var songh = canvas.NewText("Song", tc.White)
var albumh = canvas.NewText("Album", tc.White)
var lengthh = canvas.NewText("Length", tc.White)

var artist = canvas.NewText("", green)
var song = canvas.NewText("", green)
var album = canvas.NewText("", green)
var length = canvas.NewText("", green)

var asalgridhead = container.New(layout.NewGridLayout(4), artisth, songh, albumh, lengthh)
var asalgrid = container.New(layout.NewGridLayout(4), artist, song, album, length)

var categoryh = canvas.NewText("Category", tc.White)
var dayh = canvas.NewText("Day", tc.White)
var hourh = canvas.NewText("Hour", tc.White)
var positionh = canvas.NewText("Position", tc.White)

var category = canvas.NewText("", green)
var day = canvas.NewText("", green)
var hour = canvas.NewText("", green)
var position = canvas.NewText("", green)
var cdhpgridhead = container.New(layout.NewGridLayout(4), categoryh, dayh, hourh, positionh)
var cdhpgrid = container.New(layout.NewGridLayout(4), category, day, hour, position)

var stph = canvas.NewText("Spins To Play", tc.White)
var sltph = canvas.NewText("Spins Left To Play", tc.White)
var stp = canvas.NewText("", green)
var sltp = canvas.NewText("", green)
var ssgridhead = container.New(layout.NewGridLayout(2), stph, sltph)
var ssgrid = container.New(layout.NewGridLayout(2), stp, sltp)

var cucath = canvas.NewText("Next Category", tc.White)
var cudayh = canvas.NewText("Day", tc.White)
var cuhourh = canvas.NewText("Hour", tc.White)
var cuposh = canvas.NewText("Next Position", tc.White)
var custph = canvas.NewText("Next Spins", tc.White)
var cugridhead = container.New(layout.NewGridLayout(5), cucath, cudayh, cuhourh, cuposh, custph)

var cucat1 = canvas.NewText("", green)
var cuday1 = canvas.NewText("", green)
var cuhour1 = canvas.NewText("", green)
var cupos1 = canvas.NewText("", green)
var nextspins1 = ""
var custp1 = canvas.NewText(nextspins1, green)
var cugrid1 = container.New(layout.NewGridLayout(5), cucat1, cuday1, cuhour1, cupos1, custp1)

var cucat2 = canvas.NewText("", green)
var cuday2 = canvas.NewText("", green)
var cuhour2 = canvas.NewText("", green)
var cupos2 = canvas.NewText("", green)
var nextspins2 = ""
var custp2 = canvas.NewText(nextspins2, green)
var cugrid2 = container.New(layout.NewGridLayout(5), cucat2, cuday2, cuhour2, cupos2, custp2)

var cucat3 = canvas.NewText("", green)
var cuday3 = canvas.NewText("", green)
var cuhour3 = canvas.NewText("", green)
var cupos3 = canvas.NewText("", green)
var nextspins3 = ""
var custp3 = canvas.NewText(nextspins3, green)
var cugrid3 = container.New(layout.NewGridLayout(5), cucat3, cuday3, cuhour3, cupos3, custp3)

func drawgGui(oa DJ) {
	//progress = widget.NewProgressBar()
	progress.Min = 0
	ttp, _ = strconv.ParseFloat(oa.Length, 64)
	progress.Max = ttp

	if strings.HasPrefix("DJ", oa.SchedCategory) && (oa.Length == "0" || oa.Length == "00") {
		oa.Length = "300"
	}
	if strings.HasPrefix("NWS", oa.SchedCategory) && (oa.Length == "0" || oa.Length == "00") {
		oa.Length = "60"
	}

	go func() {
		for tl = 1.0; tl <= ttp; tl++ {
			time.Sleep(time.Second)
			progress.SetValue(tl)
		}
	}()

	artist.Text = oa.Artist
	song.Text = oa.Song
	album.Text = oa.Album
	length.Text = oa.Length

	category.Text = oa.SchedCategory
	day.Text = oa.SchedDay
	hour.Text = oa.SchedHour
	position.Text = oa.SchedPosition

	stp.Text = oa.SchedSpinsToPlay
	sltp.Text = oa.SchedSpinsLefToPlay

	if oa.SchedPosition == "0" || oa.SchedPosition == "00" {
		oa.SchedPosition = "99"
	}
	config.ScheduleGetPlan(oa.SchedDay, oa.SchedHour, oa.SchedPosition)

	cucat1.Text = config.SchedulePlan[0].Category
	cuday1.Text = config.SchedulePlan[0].Days
	cuhour1.Text = config.SchedulePlan[0].Hours
	cupos1.Text = config.SchedulePlan[0].Position
	nextspins1 = strconv.Itoa(config.SchedulePlan[0].Spinstoplay)
	custp1.Text = nextspins1

	cucat2.Text = config.SchedulePlan[1].Category
	cuday2.Text = config.SchedulePlan[1].Days
	cuhour2.Text = config.SchedulePlan[1].Hours
	cupos2.Text = config.SchedulePlan[1].Position
	nextspins2 = strconv.Itoa(config.SchedulePlan[1].Spinstoplay)
	custp2.Text = nextspins2

	cucat3.Text = config.SchedulePlan[2].Category
	cuday3.Text = config.SchedulePlan[2].Days
	cuhour3.Text = config.SchedulePlan[2].Hours
	cupos3.Text = config.SchedulePlan[2].Position
	nextspins3 = strconv.Itoa(config.SchedulePlan[2].Spinstoplay)
	custp3.Text = nextspins3

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
		progresgrid,
		widget.NewLabel(" "),
		cugridhead,
		cugrid1,
		cugrid2,
		cugrid3,
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
	config.NatsBucketType = config.Decrypt(fmt.Sprintf("%v", cfg["NatsBucketType"]), MySecret)
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
