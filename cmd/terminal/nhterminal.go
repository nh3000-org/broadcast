package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/nh3000-org/broadcast/config"
	"github.com/rivo/tview"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

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

func domemory() tview.Table {
	v, _ := mem.VirtualMemory()
	vtotal := strconv.FormatUint(v.Total/1024/1024, 10)
	vfree := strconv.FormatUint(v.Free/1024/1024, 10)
	vusedpercent := strconv.FormatFloat(v.UsedPercent, 'f', 2, 64)
	runtime.ReadMemStats(&memoryStats)
	cpupercent, _ := cpu.Percent(time.Second, true)
	vcpupercent := strconv.FormatFloat(cpupercent[0], 'f', 2, 64)
	usage, _ := disk.Usage("/")
	usagehome, _ := disk.Usage("/home")
	usageopt, _ := disk.Usage("/opt")

	table := tview.NewTable()
	table.SetBorder(true)
	// header r,c
	table.SetCell(0, 0, tview.NewTableCell("Type").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	table.SetCell(1, 0, tview.NewTableCell("MEM Total").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	table.SetCell(2, 0, tview.NewTableCell("MEM Free").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	table.SetCell(3, 0, tview.NewTableCell("MEM Used").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))

	table.SetCell(5, 0, tview.NewTableCell("CPU Used").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))

	table.SetCell(7, 0, tview.NewTableCell("PGM Used").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))

	table.SetCell(9, 0, tview.NewTableCell("Disk /").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))

	table.SetCell(11, 0, tview.NewTableCell("Disk /home").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))

	table.SetCell(13, 0, tview.NewTableCell("Disk /opt").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	// values
	table.SetCell(0, 1, tview.NewTableCell("Value").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	table.SetCell(1, 1, tview.NewTableCell(vtotal).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	table.SetCell(2, 1, tview.NewTableCell(vfree).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))
	table.SetCell(3, 1, tview.NewTableCell(vcpupercent).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

	table.SetCell(5, 1, tview.NewTableCell(vusedpercent).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

	table.SetCell(7, 1, tview.NewTableCell(strconv.FormatUint(memoryStats.Alloc/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

	table.SetCell(9, 1, tview.NewTableCell(strconv.FormatUint(usage.Total/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

	table.SetCell(11, 1, tview.NewTableCell(strconv.FormatUint(usagehome.Total/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

	table.SetCell(13, 1, tview.NewTableCell(strconv.FormatUint(usageopt.Total/1024/1024, 10)).SetTextColor(tcell.ColorGreen).SetAlign(tview.AlignLeft))

	// uom
	table.SetCell(0, 2, tview.NewTableCell("UOM").SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignLeft))
	table.SetCell(1, 2, tview.NewTableCell("mb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))
	table.SetCell(2, 2, tview.NewTableCell("mb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))
	table.SetCell(3, 2, tview.NewTableCell("%").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))

	table.SetCell(5, 2, tview.NewTableCell("%").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))

	table.SetCell(7, 2, tview.NewTableCell("k").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))

	table.SetCell(9, 2, tview.NewTableCell("gb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))

	table.SetCell(11, 2, tview.NewTableCell("gb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))

	table.SetCell(13, 2, tview.NewTableCell("gb").SetTextColor(tcell.ColorBlue).SetAlign(tview.AlignLeft))

	return *table
}

func drawgui(app *tview.Application, flex *tview.Flex) {
	runtime.GC()
	readPreferences()
	memcpudisk := domemory()
	memcpu := tview.NewTextView()
	memcpu.SetTitle("MEM/CPU/DISK")
	memcpu.SetLabel("")
	memcpu.SetText("")
	memcpu.SetBorder(true)

	flex.AddItem(&memcpudisk, 0, 1, false)

	nats := tview.NewTextView()
	nats.SetTitle("On Air")

	du := "\n\nONair \n/ Total "

	nats.SetText(du)
	nats.SetBorder(true)
	flex.AddItem(nats, 0, 2, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
func main() {
	app := tview.NewApplication()
	flex := tview.NewFlex()

	drawgui(app, flex)

}
