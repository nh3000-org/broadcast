package main

import (
	"runtime"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

//var frame *tview.Frame
//var app *tview.Application

var memoryStats runtime.MemStats

func drawgui(app *tview.Application, flex *tview.Flex) {
	runtime.GC()
	v, _ := mem.VirtualMemory()
	vtotal := strconv.FormatUint(v.Total/1024/1024, 10)
	vfree := strconv.FormatUint(v.Free/1024/1024, 10)
	vusedpercent := strconv.FormatFloat(v.UsedPercent, 'f', 2, 64)
	runtime.ReadMemStats(&memoryStats)
	cpupercent, _ := cpu.Percent(time.Second, true)
	usage, _ := disk.Usage("/")
	usagehome, _ := disk.Usage("/home")
	mem := "MEMORY: " + "\nTotal: " + vtotal + "mb\nFree: " + vfree + "mb\nUsed:" + vusedpercent + "%" + "\n\nPgm:\n" + strconv.FormatUint(memoryStats.Alloc/1024, 10) + "k" +
		"\n\nCPU:\n" + strconv.FormatFloat(cpupercent[0], 'f', 2, 64) +
		"\n\nDisk Root \nTotal " + strconv.FormatUint(usage.Total/1024/1024, 10) + "gb" + "\nUsed " + strconv.FormatUint(usage.Used/1024/1024, 10) + "gb" + "\nFree " + strconv.FormatUint(usage.Free/1024/1024, 10) + "gb" +
		"\n\nDisk /home\nTotal " + strconv.FormatUint(usagehome.Total/1024/1024, 10) + "gb" + "\nUsed " + strconv.FormatUint(usagehome.Used/1024/1024, 10) + "gb" + "\nFree " + strconv.FormatUint(usagehome.Free/1024/1024, 10) + "gb"

	memcpu := tview.NewTextView()
	memcpu.SetTitle("MEM/CPU/DISK")
	memcpu.SetText(mem)
	memcpu.SetBorder(true)
	if v.UsedPercent < 80.0 {
		memcpu.SetTextColor(tcell.ColorGreen)
	} else {
		memcpu.SetTextColor(tcell.ColorGreen)
	}

	flex.AddItem(memcpu, 0, 1, false)

	nats := tview.NewTextView()
	nats.SetTitle("On Air")

	du := "\n\nDisk \n/ Total " + strconv.FormatUint(usage.Total/1024/1024, 10) + "gb\n" + "\n/ Used " + strconv.FormatUint(usage.Used/1024/1024, 10) + "gb\n" + "\n/ Free " + strconv.FormatUint(usage.Free/1024/1024, 10) + "gb\n"

	nats.SetText(du)
	nats.SetBorder(true)
	flex.AddItem(nats, 0, 2, false)
	//memcpu.AddItem(tview.NewTextArea().SetText(mem, false), 30, 0, false)

	//frame.AddText("CPU Percentage", true, tview.AlignLeft, tcell.ColorWhite)
	/* 	frame.AddText(mem, true, tview.AlignLeft, tcell.ColorWhite)
	   	frame.AddText("Cache Usage", false, tview.AlignRight, tcell.ColorWhite)
	   	frame.AddText("Onair", false, tview.AlignCenter, tcell.ColorRed)
	   	frame.AddText("Footer middle", false, tview.AlignCenter, tcell.ColorGreen)
	   	frame.AddText("Footer second middle", false, tview.AlignCenter, tcell.ColorGreen) */
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
func main() {
	app := tview.NewApplication()
	flex := tview.NewFlex()

	drawgui(app, flex)

}
