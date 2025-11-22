package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/nh3000-org/broadcast/config"
	"github.com/rivo/tview"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

var PreferencesLocation = "/home/oem/.config/fyne/org.nh3000.nh3000/preferences.json"

const MySecret string = "abd&1*~#^2^#s0^=)^^7%c34"

var app *tview.Application

var flex *tview.Flex
var memcpu *tview.TextView
var vtotal string
var vfree string
var vusedpercent string
var ttp float64
var tl float64
var timeleft float64
var s string
var cpupercent []float64
var vcpupercent string
var usage *disk.UsageStat
var usagehome *disk.UsageStat
var usageopt *disk.UsageStat
var v *mem.VirtualMemoryStat
var nextspins1 string
var nextspins2 string
var nextspins3 string
var jsondata []byte
var readerr error
var errunmarshal error

func readPreferences() {
	// read config preferences.json
	jsondata, readerr = os.ReadFile(PreferencesLocation)
	if readerr != nil {
		log.Println("ERROR Preferences readerr ", readerr)
	}
	// parse json
	var cfg map[string]any
	errunmarshal = json.Unmarshal(jsondata, &cfg)
	if errunmarshal != nil {
		log.Println("ERROR Preferences errunmarshal ", errunmarshal)
	}

	config.DBpassword = config.Decrypt(fmt.Sprintf("%v", cfg["DBPASSWORD"]), MySecret)

	config.DBaddress = config.Decrypt(fmt.Sprintf("%v", cfg["DBADDRESS"]), MySecret)
	//log.Println(config.DBaddress)

	config.DBuser = config.Decrypt(fmt.Sprintf("%v", cfg["DBUSER"]), MySecret)
	config.NatsBucketType = config.Decrypt(fmt.Sprintf("%v", cfg["NatsBucketType"]), MySecret)
	config.NatsCaroot = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCaroot"]), MySecret)
	config.NatsClientkey = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCakey"]), MySecret)
	config.NatsClientcert = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCaclient"]), MySecret)
	config.NatsQueuePassword = config.Decrypt(fmt.Sprintf("%v", cfg["NatsQueuePassword"]), MySecret)

	//log.Println("CONFIG NatsBucketType", config.NatsBucketType)
	//log.Println("NATS AUTH user", config.NatsServer, config.NatsUser, config.NatsUserPassword)
	config.NewNatsJS()
	config.NewPGSQL()
}

var memoryStats runtime.MemStats
var mcp tview.Table
var msgv tview.TextView

func domemory() {
	v, _ = mem.VirtualMemory()
	vtotal = strconv.FormatUint(v.Total/1024/1024, 10)
	vfree = strconv.FormatUint(v.Free/1024/1024, 10)
	vusedpercent = strconv.FormatFloat(v.UsedPercent, 'f', 2, 64)
	runtime.ReadMemStats(&memoryStats)
	cpupercent, _ = cpu.Percent(time.Second, true)
	vcpupercent = strconv.FormatFloat(cpupercent[0], 'f', 2, 64)
	usage, _ = disk.Usage("/")
	usagehome, _ = disk.Usage("/home")
	usageopt, _ = disk.Usage("/opt")

	mcp.SetBorder(true)
	// header r,c
	mcp.SetCell(0, 0, tview.NewTableCell("Type").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	mcp.SetCell(1, 0, tview.NewTableCell("MEM Total").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	mcp.SetCell(2, 0, tview.NewTableCell("MEM Free").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	mcp.SetCell(3, 0, tview.NewTableCell("MEM Used").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))

	mcp.SetCell(5, 0, tview.NewTableCell("CPU Used").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))

	mcp.SetCell(7, 0, tview.NewTableCell("PGM Used").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))

	mcp.SetCell(9, 0, tview.NewTableCell("/").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	mcp.SetCell(10, 0, tview.NewTableCell("Used").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	mcp.SetCell(11, 0, tview.NewTableCell("Free").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))

	mcp.SetCell(13, 0, tview.NewTableCell("/home").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	mcp.SetCell(14, 0, tview.NewTableCell("Used").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	mcp.SetCell(15, 0, tview.NewTableCell("Free").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))

	mcp.SetCell(17, 0, tview.NewTableCell("/opt").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	mcp.SetCell(18, 0, tview.NewTableCell("Used").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	mcp.SetCell(19, 0, tview.NewTableCell("Free").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	// values
	mcp.SetCell(0, 1, tview.NewTableCell("Value").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	mcp.SetCell(1, 1, tview.NewTableCell(vtotal).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	mcp.SetCell(2, 1, tview.NewTableCell(vfree).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	mcp.SetCell(3, 1, tview.NewTableCell(vcpupercent).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

	mcp.SetCell(5, 1, tview.NewTableCell(vusedpercent).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

	mcp.SetCell(7, 1, tview.NewTableCell(strconv.FormatUint(memoryStats.Alloc/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

	mcp.SetCell(9, 1, tview.NewTableCell(strconv.FormatUint(usage.Total/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	mcp.SetCell(10, 1, tview.NewTableCell(strconv.FormatUint(usage.Used/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	mcp.SetCell(11, 1, tview.NewTableCell(strconv.FormatUint(usage.Free/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

	mcp.SetCell(13, 1, tview.NewTableCell(strconv.FormatUint(usagehome.Total/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	mcp.SetCell(14, 1, tview.NewTableCell(strconv.FormatUint(usagehome.Used/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	mcp.SetCell(15, 1, tview.NewTableCell(strconv.FormatUint(usagehome.Free/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

	mcp.SetCell(17, 1, tview.NewTableCell(strconv.FormatUint(usageopt.Total/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	mcp.SetCell(18, 1, tview.NewTableCell(strconv.FormatUint(usageopt.Used/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	mcp.SetCell(19, 1, tview.NewTableCell(strconv.FormatUint(usageopt.Free/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

	// uom
	mcp.SetCell(0, 2, tview.NewTableCell("UOM").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	mcp.SetCell(1, 2, tview.NewTableCell("mb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))
	mcp.SetCell(2, 2, tview.NewTableCell("mb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))
	mcp.SetCell(3, 2, tview.NewTableCell("%").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))

	mcp.SetCell(5, 2, tview.NewTableCell("%").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))

	mcp.SetCell(7, 2, tview.NewTableCell("k").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))

	mcp.SetCell(9, 2, tview.NewTableCell("gb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))
	mcp.SetCell(10, 2, tview.NewTableCell("gb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))
	mcp.SetCell(11, 2, tview.NewTableCell("gb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))

	mcp.SetCell(13, 2, tview.NewTableCell("gb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))
	mcp.SetCell(14, 2, tview.NewTableCell("gb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))
	mcp.SetCell(15, 2, tview.NewTableCell("gb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))

	mcp.SetCell(17, 2, tview.NewTableCell("gb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))
	mcp.SetCell(18, 2, tview.NewTableCell("gb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))
	mcp.SetCell(19, 2, tview.NewTableCell("gb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))

}

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
var onair tview.Table

func drawonair() {

	onair.SetBorder(true)

	// header r,c
	onair.SetCell(0, 0, tview.NewTableCell("On Air").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	onair.SetCell(0, 1, tview.NewTableCell("Artist").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	onair.SetCell(0, 2, tview.NewTableCell("Song").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	onair.SetCell(0, 3, tview.NewTableCell("Album").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	onair.SetCell(0, 4, tview.NewTableCell("Length").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	onair.SetCell(0, 5, tview.NewTableCell("Left").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	// onair r,c
	onair.SetCell(1, 0, tview.NewTableCell("Playing").SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(1, 1, tview.NewTableCell(DJJSON.Artist).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(1, 2, tview.NewTableCell(DJJSON.Song).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(1, 3, tview.NewTableCell(DJJSON.Album).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	if strings.HasPrefix("DJ", DJJSON.SchedCategory) && (DJJSON.Length == "0" || DJJSON.Length == "00") {
		DJJSON.Length = "300"
	}
	if strings.HasPrefix("NWS", DJJSON.SchedCategory) && (DJJSON.Length == "0" || DJJSON.Length == "00") {
		DJJSON.Length = "60"
	}
	onair.SetCell(1, 4, tview.NewTableCell(DJJSON.Length).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))

	ttp, _ = strconv.ParseFloat(DJJSON.Length, 64)
	go func() {
		for tl = 1.0; tl <= ttp; tl++ {
			time.Sleep(time.Second)
			timeleft = ttp - tl
			s = strconv.FormatFloat(timeleft, 'f', -1, 64)
			if timeleft > 10 {
				onair.SetCell(1, 5, tview.NewTableCell(s).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
			} else {
				onair.SetCell(1, 5, tview.NewTableCell(s).SetTextColor(tcell.ColorRed).SetAlign(tview.AlignLeft))
			}
		}
	}()

	// header r,c
	onair.SetCell(3, 0, tview.NewTableCell("Category").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	onair.SetCell(3, 1, tview.NewTableCell("Day").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	onair.SetCell(3, 2, tview.NewTableCell("Hour").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	onair.SetCell(3, 3, tview.NewTableCell("Position").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	onair.SetCell(3, 4, tview.NewTableCell("Spins").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	onair.SetCell(3, 5, tview.NewTableCell("Left").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	// onair r,c
	onair.SetCell(4, 0, tview.NewTableCell(DJJSON.SchedCategory).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(4, 1, tview.NewTableCell(DJJSON.SchedDay).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(4, 2, tview.NewTableCell(DJJSON.SchedHour).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	if DJJSON.SchedPosition == "0" || DJJSON.SchedPosition == "00" {
		DJJSON.SchedPosition = "99"
	}
	onair.SetCell(4, 3, tview.NewTableCell(DJJSON.SchedPosition).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(4, 4, tview.NewTableCell(DJJSON.SchedSpinsToPlay).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(4, 5, tview.NewTableCell(DJJSON.SchedSpinsLefToPlay).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

	// header r,c
	onair.SetCell(6, 0, tview.NewTableCell("Category").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	onair.SetCell(6, 1, tview.NewTableCell("Day").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	onair.SetCell(6, 2, tview.NewTableCell("Hour").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	onair.SetCell(6, 3, tview.NewTableCell("Position").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	onair.SetCell(6, 4, tview.NewTableCell("Spins").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))

	config.ScheduleGetPlan(DJJSON.SchedDay, DJJSON.SchedHour, DJJSON.SchedPosition)

	onair.SetCell(7, 0, tview.NewTableCell(config.SchedulePlan[0].Category).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(7, 1, tview.NewTableCell(config.SchedulePlan[0].Days).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(7, 2, tview.NewTableCell(config.SchedulePlan[0].Hours).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(7, 3, tview.NewTableCell(config.SchedulePlan[0].Position).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	nextspins1 = strconv.Itoa(config.SchedulePlan[0].Spinstoplay)
	onair.SetCell(7, 4, tview.NewTableCell(nextspins1).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

	onair.SetCell(8, 0, tview.NewTableCell(config.SchedulePlan[1].Category).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(8, 1, tview.NewTableCell(config.SchedulePlan[1].Days).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(8, 2, tview.NewTableCell(config.SchedulePlan[1].Hours).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(8, 3, tview.NewTableCell(config.SchedulePlan[1].Position).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	nextspins2 = strconv.Itoa(config.SchedulePlan[1].Spinstoplay)
	onair.SetCell(8, 4, tview.NewTableCell(nextspins2).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

	onair.SetCell(9, 0, tview.NewTableCell(config.SchedulePlan[2].Category).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(9, 1, tview.NewTableCell(config.SchedulePlan[2].Days).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(9, 2, tview.NewTableCell(config.SchedulePlan[2].Hours).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	onair.SetCell(9, 3, tview.NewTableCell(config.SchedulePlan[2].Position).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	nextspins3 = strconv.Itoa(config.SchedulePlan[2].Spinstoplay)
	onair.SetCell(9, 4, tview.NewTableCell(nextspins3).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

}

var msgi int
var msgs string

func drawmessages() {

	for {

		msgs = "\n"
		for msgi = 0; msgi < len(config.NatsMessages); msgi++ {
			msgs = msgs + config.NatsMessages[msgi].MSmessage + "\n"
		}
		app.QueueUpdateDraw(func() {
			msgv.SetTitle(strconv.Itoa(len(config.NatsMessages)))
			msgv.SetText(msgs)
		})
		time.Sleep(5 * time.Minute)
	}
}
func countdown(tottime float64) {
	for tl = 1.0; tl <= tottime; tl++ {
		time.Sleep(time.Second)
		timeleft = tottime - tl
		s = strconv.FormatFloat(timeleft, 'f', -1, 64)

		if timeleft > 10 {
			app.QueueUpdateDraw(func() {
				onair.SetCell(1, 5, tview.NewTableCell(s).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
			})
		} else {
			app.QueueUpdateDraw(func() {
				onair.SetCell(1, 5, tview.NewTableCell(s).SetTextColor(tcell.ColorRed).SetAlign(tview.AlignLeft))
			})
		}

	}
}
func doonair() {

	ctxmain, ctxmaincan = context.WithCancel(context.Background())
	if config.NatsBucketType == "mp3" {
		mp3msg, mp3err = config.NATS.OnAirmp3.Watch(ctxmain, "OnAirmp3")
		if mp3err != nil {
			log.Println("ReceiveONAIRMP3", mp3err)
		}
		for {

			kve = <-mp3msg.Updates()
			//log.Println("ReceiveONAIRMP3", kve)
			if kve != nil {
				errum = json.Unmarshal(kve.Value(), &DJJSON)
				if errum != nil {
					log.Println("DJ ReceiveONAIRMP3", errum)
				}
				runtime.GC()
				runtime.ReadMemStats(&memoryStats)

				ttp, _ = strconv.ParseFloat(DJJSON.Length, 64)
				go countdown(ttp)
				app.QueueUpdateDraw(func() {
					runtime.GC()
					runtime.ReadMemStats(&memoryStats)
					v, _ = mem.VirtualMemory()
					vtotal = strconv.FormatUint(v.Total/1024/1024, 10)
					vfree = strconv.FormatUint(v.Free/1024/1024, 10)
					vusedpercent = strconv.FormatFloat(v.UsedPercent, 'f', 2, 64)
					runtime.ReadMemStats(&memoryStats)
					cpupercent, _ = cpu.Percent(time.Second, true)
					vcpupercent = strconv.FormatFloat(cpupercent[0], 'f', 2, 64)
					usage, _ = disk.Usage("/")
					usagehome, _ = disk.Usage("/home")
					usageopt, _ = disk.Usage("/opt")
					mcp.SetCell(1, 1, tview.NewTableCell(vtotal).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(2, 1, tview.NewTableCell(vfree).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(3, 1, tview.NewTableCell(vcpupercent).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					mcp.SetCell(5, 1, tview.NewTableCell(vusedpercent).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					mcp.SetCell(7, 1, tview.NewTableCell(strconv.FormatUint(memoryStats.Alloc/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					mcp.SetCell(9, 1, tview.NewTableCell(strconv.FormatUint(usage.Total/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(10, 1, tview.NewTableCell(strconv.FormatUint(usage.Used/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(11, 1, tview.NewTableCell(strconv.FormatUint(usage.Free/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					mcp.SetCell(13, 1, tview.NewTableCell(strconv.FormatUint(usagehome.Total/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(14, 1, tview.NewTableCell(strconv.FormatUint(usagehome.Used/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(15, 1, tview.NewTableCell(strconv.FormatUint(usagehome.Free/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					mcp.SetCell(17, 1, tview.NewTableCell(strconv.FormatUint(usageopt.Total/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(18, 1, tview.NewTableCell(strconv.FormatUint(usageopt.Used/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(19, 1, tview.NewTableCell(strconv.FormatUint(usageopt.Free/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					onair.SetCell(1, 1, tview.NewTableCell(DJJSON.Artist).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(1, 2, tview.NewTableCell(DJJSON.Song).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(1, 3, tview.NewTableCell(DJJSON.Album).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					if strings.HasPrefix("DJ", DJJSON.SchedCategory) && (DJJSON.Length == "0" || DJJSON.Length == "00") {
						DJJSON.Length = "300"
					}
					if strings.HasPrefix("NWS", DJJSON.SchedCategory) && (DJJSON.Length == "0" || DJJSON.Length == "00") {
						DJJSON.Length = "60"
					}
					onair.SetCell(1, 4, tview.NewTableCell(DJJSON.Length).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(4, 0, tview.NewTableCell(DJJSON.SchedCategory).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(4, 1, tview.NewTableCell(DJJSON.SchedDay).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(4, 2, tview.NewTableCell(DJJSON.SchedHour).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					if DJJSON.SchedPosition == "0" || DJJSON.SchedPosition == "00" {
						DJJSON.SchedPosition = "99"
					}
					onair.SetCell(4, 3, tview.NewTableCell(DJJSON.SchedPosition).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(4, 4, tview.NewTableCell(DJJSON.SchedSpinsToPlay).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(4, 5, tview.NewTableCell(DJJSON.SchedSpinsLefToPlay).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					// header r,c
					onair.SetCell(6, 0, tview.NewTableCell("Category").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
					onair.SetCell(6, 1, tview.NewTableCell("Day").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
					onair.SetCell(6, 2, tview.NewTableCell("Hour").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
					onair.SetCell(6, 3, tview.NewTableCell("Position").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
					onair.SetCell(6, 4, tview.NewTableCell("Spins").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))

					config.ScheduleGetPlan(DJJSON.SchedDay, DJJSON.SchedHour, DJJSON.SchedPosition)

					onair.SetCell(7, 0, tview.NewTableCell(config.SchedulePlan[0].Category).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(7, 1, tview.NewTableCell(config.SchedulePlan[0].Days).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(7, 2, tview.NewTableCell(config.SchedulePlan[0].Hours).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(7, 3, tview.NewTableCell(config.SchedulePlan[0].Position).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					nextspins1 = strconv.Itoa(config.SchedulePlan[0].Spinstoplay)
					onair.SetCell(7, 4, tview.NewTableCell(nextspins1).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					onair.SetCell(8, 0, tview.NewTableCell(config.SchedulePlan[1].Category).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(8, 1, tview.NewTableCell(config.SchedulePlan[1].Days).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(8, 2, tview.NewTableCell(config.SchedulePlan[1].Hours).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(8, 3, tview.NewTableCell(config.SchedulePlan[1].Position).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					nextspins2 = strconv.Itoa(config.SchedulePlan[1].Spinstoplay)
					onair.SetCell(8, 4, tview.NewTableCell(nextspins2).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					runtime.GC()
					runtime.ReadMemStats(&memoryStats)
					onair.SetCell(9, 0, tview.NewTableCell(config.SchedulePlan[2].Category).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(9, 1, tview.NewTableCell(config.SchedulePlan[2].Days).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(9, 2, tview.NewTableCell(config.SchedulePlan[2].Hours).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(9, 3, tview.NewTableCell(config.SchedulePlan[2].Position).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					nextspins3 = strconv.Itoa(config.SchedulePlan[2].Spinstoplay)
					onair.SetCell(9, 4, tview.NewTableCell(nextspins3).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

				})
			}
		}
	}
	if config.NatsBucketType == "wav" {
		wavmsg, waverr = config.NATS.OnAirwav.Watch(ctxmain, "OnAirwav")
		if waverr != nil {
			log.Println("ReceiveONAIRWAV", waverr)
		}
		for {

			kve = <-wavmsg.Updates()
			//log.Println("ReceiveONAIRMP3", kve)
			if kve != nil {
				errum = json.Unmarshal(kve.Value(), &DJJSON)
				if errum != nil {
					log.Println("DJ ReceiveONAIRWAV", errum)
				}
				runtime.GC()
				runtime.ReadMemStats(&memoryStats)
				ttp, _ = strconv.ParseFloat(DJJSON.Length, 64)
				go countdown(ttp)
				app.QueueUpdateDraw(func() {
					runtime.GC()
					runtime.ReadMemStats(&memoryStats)
					v, _ = mem.VirtualMemory()
					vtotal = strconv.FormatUint(v.Total/1024/1024, 10)
					vfree = strconv.FormatUint(v.Free/1024/1024, 10)
					vusedpercent = strconv.FormatFloat(v.UsedPercent, 'f', 2, 64)
					runtime.ReadMemStats(&memoryStats)
					cpupercent, _ = cpu.Percent(time.Second, true)
					vcpupercent = strconv.FormatFloat(cpupercent[0], 'f', 2, 64)
					usage, _ = disk.Usage("/")
					usagehome, _ = disk.Usage("/home")
					usageopt, _ = disk.Usage("/opt")
					mcp.SetCell(1, 1, tview.NewTableCell(vtotal).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(2, 1, tview.NewTableCell(vfree).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(3, 1, tview.NewTableCell(vcpupercent).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					mcp.SetCell(5, 1, tview.NewTableCell(vusedpercent).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					mcp.SetCell(7, 1, tview.NewTableCell(strconv.FormatUint(memoryStats.Alloc/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					mcp.SetCell(9, 1, tview.NewTableCell(strconv.FormatUint(usage.Total/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(10, 1, tview.NewTableCell(strconv.FormatUint(usage.Used/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(11, 1, tview.NewTableCell(strconv.FormatUint(usage.Free/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					mcp.SetCell(13, 1, tview.NewTableCell(strconv.FormatUint(usagehome.Total/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(14, 1, tview.NewTableCell(strconv.FormatUint(usagehome.Used/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(15, 1, tview.NewTableCell(strconv.FormatUint(usagehome.Free/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					mcp.SetCell(17, 1, tview.NewTableCell(strconv.FormatUint(usageopt.Total/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(18, 1, tview.NewTableCell(strconv.FormatUint(usageopt.Used/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					mcp.SetCell(19, 1, tview.NewTableCell(strconv.FormatUint(usageopt.Free/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					onair.SetCell(1, 1, tview.NewTableCell(DJJSON.Artist).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(1, 2, tview.NewTableCell(DJJSON.Song).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(1, 3, tview.NewTableCell(DJJSON.Album).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					if strings.HasPrefix("DJ", DJJSON.SchedCategory) && (DJJSON.Length == "0" || DJJSON.Length == "00") {
						DJJSON.Length = "300"
					}
					if strings.HasPrefix("NWS", DJJSON.SchedCategory) && (DJJSON.Length == "0" || DJJSON.Length == "00") {
						DJJSON.Length = "60"
					}
					onair.SetCell(1, 4, tview.NewTableCell(DJJSON.Length).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(4, 0, tview.NewTableCell(DJJSON.SchedCategory).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(4, 1, tview.NewTableCell(DJJSON.SchedDay).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(4, 2, tview.NewTableCell(DJJSON.SchedHour).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					if DJJSON.SchedPosition == "0" || DJJSON.SchedPosition == "00" {
						DJJSON.SchedPosition = "99"
					}
					onair.SetCell(4, 3, tview.NewTableCell(DJJSON.SchedPosition).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(4, 4, tview.NewTableCell(DJJSON.SchedSpinsToPlay).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(4, 5, tview.NewTableCell(DJJSON.SchedSpinsLefToPlay).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					// header r,c
					onair.SetCell(6, 0, tview.NewTableCell("Category").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
					onair.SetCell(6, 1, tview.NewTableCell("Day").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
					onair.SetCell(6, 2, tview.NewTableCell("Hour").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
					onair.SetCell(6, 3, tview.NewTableCell("Position").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
					onair.SetCell(6, 4, tview.NewTableCell("Spins").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))

					config.ScheduleGetPlan(DJJSON.SchedDay, DJJSON.SchedHour, DJJSON.SchedPosition)

					onair.SetCell(7, 0, tview.NewTableCell(config.SchedulePlan[0].Category).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(7, 1, tview.NewTableCell(config.SchedulePlan[0].Days).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(7, 2, tview.NewTableCell(config.SchedulePlan[0].Hours).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(7, 3, tview.NewTableCell(config.SchedulePlan[0].Position).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					nextspins1 = strconv.Itoa(config.SchedulePlan[0].Spinstoplay)
					onair.SetCell(7, 4, tview.NewTableCell(nextspins1).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					onair.SetCell(8, 0, tview.NewTableCell(config.SchedulePlan[1].Category).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(8, 1, tview.NewTableCell(config.SchedulePlan[1].Days).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(8, 2, tview.NewTableCell(config.SchedulePlan[1].Hours).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(8, 3, tview.NewTableCell(config.SchedulePlan[1].Position).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					nextspins2 = strconv.Itoa(config.SchedulePlan[1].Spinstoplay)
					onair.SetCell(8, 4, tview.NewTableCell(nextspins2).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

					onair.SetCell(9, 0, tview.NewTableCell(config.SchedulePlan[2].Category).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(9, 1, tview.NewTableCell(config.SchedulePlan[2].Days).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(9, 2, tview.NewTableCell(config.SchedulePlan[2].Hours).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					onair.SetCell(9, 3, tview.NewTableCell(config.SchedulePlan[2].Position).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
					nextspins3 = strconv.Itoa(config.SchedulePlan[2].Spinstoplay)
					onair.SetCell(9, 4, tview.NewTableCell(nextspins3).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

				})
				runtime.GC()
				runtime.ReadMemStats(&memoryStats)

			}
		}
	}

}

func main() {
	readPreferences()
	app = tview.NewApplication()
	flex = tview.NewFlex()
	mcp = *tview.NewTable()
	domemory()
	memcpu = tview.NewTextView()
	memcpu.SetTitle("MEM/CPU/DISK")
	memcpu.SetLabel("")
	memcpu.SetText("")
	memcpu.SetBorder(true)
	flex.AddItem(&mcp, 0, 1, false)
	onair = *tview.NewTable()
	//doonair()
	drawonair()
	flex.AddItem(&onair, 0, 3, false)
	//drawonair()
	go doonair()
	msgv = *tview.NewTextView()
	msgv.SetWordWrap(true)
	msgv.SetBorder(true)
	msgv.SetTitle("Messages")
	msgv.SetText("Error Messages")
	msgv.SetWordWrap(true)
	msgv.SetScrollable(true)

	go config.ReceiveMESSAGE()
	go drawmessages()
	flex.AddItem(&msgv, 0, 1, false)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
