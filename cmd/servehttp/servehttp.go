package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/nh3000-org/broadcast/config"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type UserSessionJSON struct {
	UserToken          string   `json:"token"`        // expiry token
	UserPasswordHash   string   `json:"passwordhash"` // password hash
	UserIPA            string   `json:"ips"`          // session remote addresses
	UserAuthCategories []string `json:"categoriess"`  // authorized categories
	UserAuthAction     []string `json:"actions"`      // actions allowed
}

var SessionCategories []string
var SessionAction []string

var PreferencesLocation = "/home/oem/.config/fyne/org.nh3000.nh3000/preferences.json"
var HashLocation = "/home/oem/.config/fyne/org.nh3000.nh3000/config.hash"
var authtoken = ""

var isbusy = false

var userdata = ""
var KeyAes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}  // must be 16 bytes
var KeyHmac = []byte{36, 45, 53, 21, 87, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05} // must be 16 bytes
const MySecret string = "abd&1*~#^2^#s0^=)^^7%c34"

func ADS(w http.ResponseWriter, r *http.Request) {
	if !checkauthorization(r.FormValue("Authorization")) {
		w.Write([]byte(ilogon()))
		return
	}
	line := charts.NewLine()
	items := make([]opts.LineData, 0)

	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Ads",
			Subtitle: "Series",
		}))

	var xdates []string // yyyy-mm-dd
	var xdesc []string  // mm-dd

	rd := r.FormValue("Days")
	rangedays, err := strconv.Atoi(rd)
	if err != nil {
		rangedays = 7
		log.Println("ADS Days Error", err)
	}
	for d := rangedays; d > 0; d-- {
		hours := 24 * d
		parm := "-" + strconv.Itoa(hours) + "h"
		dt := config.GetDateTime(parm)
		xdates = append(xdates, dt[0:10])
		xdesc = append(xdesc, dt[5:10])
		//log.Println("Setting Range", dt[0:10], dt[5:10])
	}
	dt := config.GetDateTime("0h")
	xdates = append(xdates, dt[0:10])
	xdesc = append(xdesc, dt[5:10])
	//log.Println("Setting Range", dt[0:10], dt[5:10])
	c := r.FormValue("Categories1")
	//log.Println("ADS Categories", c)

	//ycats := make(map[string]int) // ADS-xdates Count .....

	line.SetXAxis(xdesc)

	//for x := 0; x < len(c); x++ {
	for d := 0; d < len(xdates); d++ {
		//if strings.HasPrefix(c[x], "ADS") {
		data := config.TrafficGetCountByAlbum(xdates[d], c)
		items = append(items, opts.LineData{Value: data})
		line.AddSeries(c, items)
		//log.Println("add series", c, items)
		//}
	}
	//}
	line.PageTitle = "Advertising Line Chart"
	line.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(true)}))
	line.Render(w)
}

var stend = "23:59:59"
var ststart = "00:00:00"

func TRAFICREPORT(w http.ResponseWriter, r *http.Request) {
	if !checkauthorization(r.FormValue("Authorization")) {
		w.Write([]byte(ilogon()))
		return
	}
	rd := r.FormValue("Days1")
	sd, e := strconv.Atoi(rd)
	if e != nil {
		log.Println("ERROR ", e)
	}
	sd = sd * 24
	log.Println("SD", "-"+strconv.Itoa(sd))
	config.TrafficStart = config.GetDateTime("-" + strconv.Itoa(sd) + "h")[0:10] + " " + ststart
	config.TrafficEnd = config.GetDateTime("-0h")[0:10] + " " + stend
	//selalbum := widget.NewSelect(config.AlbumToArray(), func(string) {})
	log.Println("TRAFFICREPORT", sd, config.TrafficStart, config.TrafficEnd)
	config.TrafficAlbum = r.FormValue("Categories1")
	config.ToPDF("TrafficReport", "ADMIN")
	cmd := exec.Command("xdg-open", "TrafficReport.pdf")
	cmderr := cmd.Start()
	log.Println("Traffic", cmderr, "rpt", "TrafficReport.pdf")
	if cmderr != nil {
		log.Println("Traffic", cmderr, "rpt", "TrafficReport.pdf")
	}
}
func chart(w http.ResponseWriter, r *http.Request) {
	if !checkauthorization(r.FormValue("Authorization")) {
		w.Write([]byte(ilogon()))
		return
	}
	line := charts.NewLine()
	items := make([]opts.LineData, 0)

	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Charts",
			Subtitle: "Series",
		}))

	var xdates []string // yyyy-mm-dd
	var xdesc []string  // mm-dd

	rd := r.FormValue("Days")
	rangedays, err := strconv.Atoi(rd)
	if err != nil {
		rangedays = 7
		log.Println("Chart Days Error", err)
	}
	for d := rangedays; d > 0; d-- {
		hours := 24 * d
		parm := "-" + strconv.Itoa(hours) + "h"
		dt := config.GetDateTime(parm)
		xdates = append(xdates, dt[0:10])
		xdesc = append(xdesc, dt[5:10])
		//log.Println("Setting Range", dt[0:10], dt[5:10])
	}
	dt := config.GetDateTime("0h")
	xdates = append(xdates, dt[0:10])
	xdesc = append(xdesc, dt[5:10])
	//log.Println("Setting Range", dt[0:10], dt[5:10])
	c := r.FormValue("Categories")
	//log.Println("Categories", c)

	//ycats := make(map[string]int) // ADS-xdates Count .....

	line.SetXAxis(xdesc)

	//for x := 0; x < len(c); x++ {
	for d := 0; d < len(xdates); d++ {
		//if strings.HasPrefix(c[x], "ADS") {
		data := config.TrafficGetCountByDate(c, xdates[d])
		items = append(items, opts.LineData{Value: data})
		line.AddSeries(c, items)
		//}
	}
	//}
	line.PageTitle = "History Chart"
	line.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(true)}))
	line.Render(w)
}

func counts(w http.ResponseWriter, r *http.Request) {
	if !checkauthorization(r.FormValue("Authorization")) {
		w.Write([]byte(ilogon()))
		return
	}
	pie := charts.NewPie()

	// set some global options like Title/Legend/ToolTip or anything else
	pie.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Music",
			Subtitle: "Distribution",
		}))

	destinations := []opts.PieData{{Name: "CURRENTS", Value: config.InventoryGetCount("CURRENTS")},
		{Name: "RECURRENTS", Value: config.InventoryGetCount("RECURRENTS")},
		{Name: "PROMOS", Value: config.InventoryGetCount("PROMOS")},
		{Name: "FILLTOTOH", Value: config.InventoryGetCount("FILLTOTOH")},
		{Name: "ADS", Value: config.InventoryGetCount("ADS")},
		{Name: "IMAGINGID", Value: config.InventoryGetCount("IMAGINGID")}}

	pie.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeChalk}),
		charts.WithTitleOpts(opts.Title{Title: "Inventory Distributions"}),
	)
	pie.AddSeries("distributions", destinations)

	pie.PageTitle = "Distribution Chart"

	pie.Render(w)
}
func schedcounts(w http.ResponseWriter, r *http.Request) {
	if !checkauthorization(r.FormValue("Authorization")) {
		w.Write([]byte(ilogon()))
		return
	}
	pie := charts.NewPie()

	// set some global options like Title/Legend/ToolTip or anything else
	pie.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Music",
			Subtitle: "Distribution",
		}))

	destinations := []opts.PieData{{Name: "CURRENTS", Value: config.ScheduleGetCount("CURRENTS")},
		{Name: "RECURRENTS", Value: config.ScheduleGetCount("RECURRENTS")},
		{Name: "PROMOS", Value: config.ScheduleGetCount("PROMOS")},
		{Name: "ADS", Value: config.ScheduleGetCount("ADS")},
		{Name: "FILLTOTOH", Value: config.InventoryGetCount("FILLTOTOH")},
		{Name: "IMAGINGID", Value: config.ScheduleGetCount("IMAGINGID")}}

	pie.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeChalk}),
		charts.WithTitleOpts(opts.Title{Title: "Schedule Distributions"}),
	)
	pie.AddSeries("distributions", destinations)

	pie.PageTitle = "Schedule Distribution Chart"

	pie.Render(w)
}
func cleartraffic(w http.ResponseWriter, r *http.Request) {
	if !checkauthorization(r.FormValue("Authorization")) {
		w.Write([]byte(ilogon()))
		return
	}

	w.Write([]byte(config.TrafficClear() + " Traffic Records Deleted"))
}
func uploadFile(w http.ResponseWriter, r *http.Request) {
	if !checkauthorization(r.FormValue("Authorization")) {
		w.Write([]byte(ilogon()))
		return
	}
	isbusy = true
	importHome := "/opt/radio/upload/stub"

	log.Println("File Upload Endpoint Hit for User", importHome)

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	pmuerr := r.ParseMultipartForm(10 << 20)
	if pmuerr != nil {
		log.Println("File Upload r.FormFile", pmuerr)
		w.Write([]byte("File Upload Parse Error r.FormFile"))
		isbusy = false
		return
	}

	file, handler, reqerr := r.FormFile("stub")
	if reqerr != nil {
		w.Write([]byte("File Upload Error r.FormFile"))
		log.Println("File Upload r.FormFile", reqerr)
		isbusy = false
		return
	}
	defer file.Close()

	// Create a destination file to copy upload into
	thezipped, err0 := os.Create(importHome + ".zip")
	if err0 != nil {
		w.Write([]byte("File create error" + ": " + importHome + ".zip" + " error " + err0.Error()))
		log.Println("File create error"+": "+importHome+".zip", err0)
		isbusy = false
		return
	}
	defer thezipped.Close()

	// Upload the file to the destination path
	nb_bytes, _ := io.Copy(thezipped, file)

	fmt.Println("File uploaded successfully", nb_bytes)
	w.Write([]byte("File uploaded successfully"))
	fmt.Printf("\nUploaded File: %+v\n", handler.Filename)
	fmt.Printf("\nFile Size: %+v\n", handler.Size)
	fmt.Printf("\nMIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern

	os.RemoveAll("/opt/radio/upload/" + strings.Replace(handler.Filename, ".zip", "", 1))
	//os.Remove("/opt/radio/upload/stub.zip")

	log.Println("UNZIP input: ", "/opt/radio/upload/stub", "output", importHome)
	os.Chdir("/opt/radio/upload")
	cmd := exec.Command("unzip", "/opt/radio/upload/stub.zip")
	//cmd := exec.Command("unzip", handler.Filename, "-d", importHome)
	out, err := cmd.Output()
	if err != nil {
		w.Write([]byte("\nUNZIP could not run command\n"))
		log.Println("UNZIP could not run command: ", err, "importhome", importHome)
	} else {
		w.Write([]byte(string(out)))
		log.Println("Output: ", string(out))
	}

	var imartist string
	var imsong string
	var imalbum string

	var imcategory string
	sp := "/opt/radio/upload/" + strings.Replace(handler.Filename, ".zip", "", 1)
	os.Chdir(sp)
	startpath := strings.Replace(sp, "/README.txt", "", 1)
	log.Println("WalkPath: ", startpath)
	walkstuberr := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {

		removepath := startpath + "/"
		cat := strings.Replace(path, removepath, "", 1)
		imimportdir := startpath + "/" + cat
		if info.IsDir() {
			imcategory = cat
		}

		if strings.HasSuffix(cat, "mp3") {
			rmcat := imcategory + "/"
			songfull := strings.ReplaceAll(path, rmcat, "")
			songunparsed := strings.ReplaceAll(songfull, ".mp3", "")
			result := strings.Split(songunparsed, "-")
			if len(result) == 0 {
				log.Println("messages."+config.NatsAlias, "Unparsed"+songunparsed, config.NatsAlias)
				config.Send("messages."+config.NatsAlias, "Unparsed"+songunparsed, config.NatsAlias)
			}
			if len(result) == 3 {
				imartist = result[0]
				imsong = result[1]
				imalbum = result[2]
			}
			if len(result) == 2 {
				imartist = result[0]
				imsong = result[1]
				imalbum = "Digital"
			}
			if len(result) == 1 {
				imartist = result[0]
				imsong = result[0]
				imalbum = "Digital"
			}
			if strings.HasSuffix(cat, "OUTRO.mp3") {
				imalbum = strings.ReplaceAll(imalbum, "OUTRO", "")
			}
			if strings.HasSuffix(cat, "INTRO.mp3") {
				imalbum = strings.ReplaceAll(imalbum, "INTRO", "")
			}

			maxspins := 0
			maxspinsperhour := 0
			length := 0
			today := 0
			week := 0
			total := 0

			added := config.GetDateTime("1h")
			sd := config.GetDateTime("1h")

			ed := "9999-01-01 00:00:00"
			//log.Println("init", "sd", sd[0:19], "ed", ed[0:19], "added", added[0:19])
			//ft := "9999-01-01 00:00:00"
			var hp = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}
			var dp = []string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}

			rowexists := config.InventoryGetRow(imcategory, imartist, imsong, imalbum)
			if rowexists != "0" {
				log.Println("Row Exists Skipping : ", rowexists, "cat", imcategory, "artist", imartist, "song", imsong, "album", imalbum)
				w.Write([]byte("\nRow Exists Skipping " + "cat: " + imcategory + " artist: " + imartist + " song: " + imsong + " album: " + imalbum + "\n"))
			}
			if rowexists == "0" {

				maxspins = 0
				maxspinsperhour = 0
				if imcategory == "ADS" {
					log.Println("before sd", sd, "ed", ed[0:19], "added", added[0:19])

					ed = config.GetDateTime("720h")

					//log.Println("in ads sd", sd, "ed", ed[0:19], "added", added[0:19])
					// default hour parts 07 - 18
					hp = []string{"06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}
					// default day slots MON-FRI
					dp = []string{"MON", "TUE", "WED", "THU", "FRI"}
					// mas spins per day 24
					maxspins = 12
					maxspinsperhour = 1

				}
				rowreturned := config.InventoryAdd(imcategory, imartist, imsong, imalbum, length, "000000", sd[0:19], ed[0:19], hp, dp, maxspins, maxspinsperhour, "1999-01-01 00:00:00", added[0:19], today, week, total, "Stub")
				row := strconv.Itoa(rowreturned)
				if row != "0" {
					songbytes, songerr := os.ReadFile(imimportdir)
					if songerr != nil {
						config.Send("messages."+config.NatsAlias, "Put Bucket Song Read Error", config.NatsAlias)
					}
					if songerr == nil {
						pberr := config.PutBucket("mp3", row, songbytes)
						if pberr == nil {
							songbytes = []byte("")
						}
						if pberr != nil {
							config.Send("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
						}
					}
					log.Println("Inventory Added", imcategory, imartist, imalbum)
					w.Write([]byte("\nAdded " + "cat: " + imcategory + " artist: " + imartist + " song: " + imsong + " album: " + imalbum + "\n"))
				}
			}

			if strings.HasSuffix(cat, "INTRO.mp3") {
				rowreturned := config.InventoryGetRow(imcategory, imartist, imsong, imalbum)
				if len(rowreturned) > 0 {
					//log.Println("importing intro", rowreturned)
					songbytes, songerr := os.ReadFile(imimportdir)
					if songerr != nil {
						log.Println("messages."+config.NatsAlias, "Put Bucket Intro Read Error", config.NatsAlias)
						config.Send("messages."+config.NatsAlias, "Put Bucket Intro Read Error", config.NatsAlias)
					}
					if songerr == nil {
						//log.Println("putting intro", rowreturned+"INTRO")
						pberr := config.PutBucket("mp3", rowreturned+"INTRO", songbytes)
						if pberr == nil {
							songbytes = []byte("")
						}
						if pberr != nil {
							log.Println("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
							config.Send("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
						}
					}
					log.Println("Inventory Added INTRO", imcategory, imartist, imalbum)
					w.Write([]byte("\nAdded INTRO " + "cat: " + imcategory + " artist: " + imartist + " song: " + imsong + " album: " + imalbum + "\n"))
				}
			}
			if strings.HasSuffix(cat, "OUTRO.mp3") {
				rowreturned := config.InventoryGetRow(imcategory, imartist, imsong, imalbum)
				if len(rowreturned) > 0 {
					//log.Println("importing outro", rowreturned)
					songbytes, songerr := os.ReadFile(imimportdir)
					if songerr != nil {
						log.Println("messages."+config.NatsAlias, "Put Bucket Outro Read Error", config.NatsAlias)
						config.Send("messages."+config.NatsAlias, "Put Bucket Outro Read Error", config.NatsAlias)
					}
					if songerr == nil {
						//log.Println("putting outro", rowreturned+"OUTRO")
						pberr := config.PutBucket("mp3", rowreturned+"OUTRO", songbytes)
						if pberr == nil {
							songbytes = []byte("")
						}
						if pberr != nil {
							log.Println("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
							config.Send("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
						}
					}
					log.Println("Inventory Added OUTRO", imcategory, imartist, imalbum)
					w.Write([]byte("\nAdded OUTRO " + "cat: " + imcategory + " artist: " + imartist + " song: " + imsong + " album: " + imalbum + "\n"))
				}
			}

		}

		if strings.HasSuffix(cat, "wav") {
			rmcat := imcategory + "/"
			songfull := strings.ReplaceAll(path, rmcat, "")
			songunparsed := strings.ReplaceAll(songfull, ".wav", "")
			result := strings.Split(songunparsed, "-")
			if len(result) == 0 {
				log.Println("messages."+config.NatsAlias, "Unparsed"+songunparsed, config.NatsAlias)
				config.Send("messages."+config.NatsAlias, "Unparsed"+songunparsed, config.NatsAlias)
			}
			if len(result) == 3 {
				imartist = result[0]
				imsong = result[1]
				imalbum = result[2]
			}
			if len(result) == 2 {
				imartist = result[0]
				imsong = result[1]
				imalbum = "Digital"
			}
			if len(result) == 1 {
				imartist = result[0]
				imsong = result[0]
				imalbum = "Digital"
			}
			if strings.HasSuffix(cat, "OUTRO.wav") {
				imalbum = strings.ReplaceAll(imalbum, "OUTRO", "")
			}
			if strings.HasSuffix(cat, "INTRO.wav") {
				imalbum = strings.ReplaceAll(imalbum, "INTRO", "")
			}

			maxspins := 0
			maxspinsperhour := 0
			length := 0
			today := 0
			week := 0
			total := 0

			added := config.GetDateTime("1h")
			sd := config.GetDateTime("1h")

			ed := "9999-01-01 00:00:00"
			//log.Println("init", "sd", sd[0:19], "ed", ed[0:19], "added", added[0:19])
			//ft := "9999-01-01 00:00:00"
			var hp = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}
			var dp = []string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}

			rowexists := config.InventoryGetRow(imcategory, imartist, imsong, imalbum)
			if rowexists != "0" {
				log.Println("Row Exists Skipping : ", rowexists, "cat", imcategory, "artist", imartist, "song", imsong, "album", imalbum)
				w.Write([]byte("\nRow Exists Skipping " + "cat: " + imcategory + " artist: " + imartist + " song: " + imsong + " album: " + imalbum + "\n"))
			}
			if rowexists == "0" {

				maxspins = 0
				maxspinsperhour = 0
				if imcategory == "ADS" {
					log.Println("before sd", sd, "ed", ed[0:19], "added", added[0:19])

					ed = config.GetDateTime("720h")

					//log.Println("in ads sd", sd, "ed", ed[0:19], "added", added[0:19])
					// default hour parts 07 - 18
					hp = []string{"06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}
					// default day slots MON-FRI
					dp = []string{"MON", "TUE", "WED", "THU", "FRI"}
					// mas spins per day 24
					maxspins = 12
					maxspinsperhour = 1

				}
				rowreturned := config.InventoryAdd(imcategory, imartist, imsong, imalbum, length, "000000", sd[0:19], ed[0:19], hp, dp, maxspins, maxspinsperhour, "1999-01-01 00:00:00", added[0:19], today, week, total, "Stub")
				row := strconv.Itoa(rowreturned)
				if row != "0" {
					songbytes, songerr := os.ReadFile(imimportdir)
					if songerr != nil {
						config.Send("messages."+config.NatsAlias, "Put Bucket Song Read Error", config.NatsAlias)
					}
					if songerr == nil {
						pberr := config.PutBucket("mp3", row, songbytes)
						if pberr == nil {
							songbytes = []byte("")
						}
						if pberr != nil {
							config.Send("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
						}
					}
					log.Println("Inventory Added", imcategory, imartist, imalbum)
					w.Write([]byte("\nAdded " + "cat: " + imcategory + " artist: " + imartist + " song: " + imsong + " album: " + imalbum + "\n"))
				}
			}

			if strings.HasSuffix(cat, "INTRO.wav") {
				rowreturned := config.InventoryGetRow(imcategory, imartist, imsong, imalbum)
				if len(rowreturned) > 0 {
					//log.Println("importing intro", rowreturned)
					songbytes, songerr := os.ReadFile(imimportdir)
					if songerr != nil {
						log.Println("messages."+config.NatsAlias, "Put Bucket Intro Read Error", config.NatsAlias)
						config.Send("messages."+config.NatsAlias, "Put Bucket Intro Read Error", config.NatsAlias)
					}
					if songerr == nil {
						//log.Println("putting intro", rowreturned+"INTRO")
						pberr := config.PutBucket("wav", rowreturned+"INTRO", songbytes)
						if pberr == nil {
							songbytes = []byte("")
						}
						if pberr != nil {
							log.Println("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
							config.Send("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
						}
					}
					log.Println("Inventory Added INTRO", imcategory, imartist, imalbum)
					w.Write([]byte("\nAdded INTRO " + "cat: " + imcategory + " artist: " + imartist + " song: " + imsong + " album: " + imalbum + "\n"))
				}
			}
			if strings.HasSuffix(cat, "OUTRO.wav") {
				rowreturned := config.InventoryGetRow(imcategory, imartist, imsong, imalbum)
				if len(rowreturned) > 0 {
					//log.Println("importing outro", rowreturned)
					songbytes, songerr := os.ReadFile(imimportdir)
					if songerr != nil {
						log.Println("messages."+config.NatsAlias, "Put Bucket Outro Read Error", config.NatsAlias)
						config.Send("messages."+config.NatsAlias, "Put Bucket Outro Read Error", config.NatsAlias)
					}
					if songerr == nil {
						//log.Println("putting outro", rowreturned+"OUTRO")
						pberr := config.PutBucket("wav", rowreturned+"OUTRO", songbytes)
						if pberr == nil {
							songbytes = []byte("")
						}
						if pberr != nil {
							log.Println("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
							config.Send("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
						}
					}
					log.Println("Inventory Added OUTRO", imcategory, imartist, imalbum)
					w.Write([]byte("\nAdded OUTRO " + "cat: " + imcategory + " artist: " + imartist + " song: " + imsong + " album: " + imalbum + "\n"))
				}
			}

		}

		return nil
	})
	if walkstuberr != nil {
		log.Println("messages.IMPORT", "Inventory Walk Err FileInfo "+walkstuberr.Error(), "onair")
		config.Send("messages.IMPORT", "Inventory Walk Err FileInfo "+walkstuberr.Error(), "onair")
	}

	// read all of the contents of our uploaded file into a
	isbusy = false
	log.Println("Successfully Processed stub File")
	log.Println(userdata)
	log.Println("Upload File")
}
func downloadFile(w http.ResponseWriter, r *http.Request) {
	if !checkauthorization(r.FormValue("Authorization")) {
		w.Write([]byte(ilogon()))
		return
	}
	isbusy = true
	log.Println("Download Stub")
	pmuerr := r.ParseForm()
	if pmuerr != nil {
		log.Println("File Download", pmuerr)
		w.Write([]byte("File Download Parse Error r.FormFile"))

	}

	importHome := "/opt/radio/blankstub"
	log.Println("Download File: ", importHome)
	config.CategoriesWriteStub(false)
	err := os.Remove("/opt/radio/blankstub/stub.zip")
	if err != nil {
		w.Write([]byte("Could not remove previous entry"))
		log.Println("Could not remove previous entry: ", err, importHome)
		isbusy = false
		return
	}
	err1 := os.Chdir(importHome)
	if err1 != nil {
		w.Write([]byte("Could not change to directory"))
		log.Println("Could not change to directory: ", err1, importHome)
		isbusy = false
		return
	}
	cmd := exec.Command("zip", "-r", "/opt/radio/blankstub/stub.zip", "stub")
	out, err3 := cmd.Output()
	if err3 != nil {
		w.Write([]byte("ZIP could not run command"))
		log.Println("ZIP could not run command: ", err3, " importHome", importHome)
		isbusy = false
		return
	} else {
		log.Println("ZIP Output: ", string(out))
	}

	hl, err4 := os.ReadFile("/opt/radio/blankstub/stub.zip")
	if err4 != nil {
		w.Write([]byte("Could not read /opt/radio/blankstub/stub.zip"))
		log.Println("Could not read: /opt/radio/blankstub/stub.zip ", err4, importHome)
		isbusy = false
		return
	}
	log.Println(userdata)
	log.Println("Download File Created")
	isbusy = false
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=stub.zip")
	w.Header().Add("Content-Length", fmt.Sprint(len(hl)))
	w.Write(hl)

}
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

	//config.WebPassword = config.Decrypt(fmt.Sprintf("%v", cfg["WEBPASSWORD"]), MySecret)

	config.DBaddress = config.Decrypt(fmt.Sprintf("%v", cfg["DBADDRESS"]), MySecret)
	//log.Println(config.DBaddress)

	config.DBuser = config.Decrypt(fmt.Sprintf("%v", cfg["DBUSER"]), MySecret)
	config.NatsAlias = config.Decrypt(fmt.Sprintf("%v", cfg["NatsAlias"]), MySecret)
	config.NatsCaroot = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCaroot"]), MySecret)
	config.NatsClientkey = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCakey"]), MySecret)
	config.NatsClientcert = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCaclient"]), MySecret)
	config.NatsQueuePassword = config.Decrypt(fmt.Sprintf("%v", cfg["NatsQueuePassword"]), MySecret)

	//log.Println("NATS AUTH user", config.NatsServer, config.NatsUser, config.NatsUserPassword)
	config.NewNatsJS()
	config.NewPGSQL()
}
func configFile(w http.ResponseWriter, r *http.Request) {
	log.Println("configFile", PreferencesLocation)
	jsondata, readerr := os.ReadFile(PreferencesLocation)
	if readerr != nil {
		log.Println("configFile ", readerr)
	}
	w.Header().Set("Content-Type", "application/text")
	w.Header().Set("Content-Disposition", "attachment; filename=preferences.json")
	w.Header().Add("Content-Length", fmt.Sprint(len(jsondata)))
	w.Write(jsondata)
}
func setupRoutes() {
	readPreferences()
	fileServer := http.FileServer(http.Dir("/opt/radio/publichtml"))
	http.Handle("/", fileServer)
	http.HandleFunc("/continue", login)
	http.HandleFunc("/login", login)
	http.HandleFunc("/config", configFile)
	http.HandleFunc("/download", downloadFile)
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/chart", chart)
	http.HandleFunc("/ADS", ADS)
	http.HandleFunc("/TRAFICREPORT", TRAFICREPORT)
	http.HandleFunc("/counts", counts)
	http.HandleFunc("/schedcounts", schedcounts)
	http.HandleFunc("/cleartraffic", cleartraffic)

	err := http.ListenAndServeTLS(":9000", "server.crt", "server.key", nil)
	if err != nil {
		log.Println("SSL ERROR ", err)
	}
}

func main() {
	fmt.Println("Writing Startup index.html")
	err := os.WriteFile("/opt/radio/publichtml/index.html", []byte(istartup()), 0644)
	if err != nil {
		log.Println("FAILED to Write Startup File", err)
	}
	fmt.Println("Waiting for Input")
	setupRoutes()
}
func checkauthorization(authtoken string) bool {
	decrypt := config.Decrypt(authtoken, MySecret)
	var jsondat = UserSessionJSON{}

	if err := json.Unmarshal([]byte(decrypt), &jsondat); err != nil {
		log.Println(err)

		return false
	}
	fmt.Println(jsondat)

	frombrowser, sterr := time.Parse(time.DateTime, jsondat.UserToken)
	if sterr != nil {
		log.Println("checkauthorization st ", sterr)
	}
	now, err2 := time.Parse(time.DateTime, config.GetDateTime("0h")[0:19])
	if err2 != nil {
		log.Println("checkauthorization oldest ", err2)
	}
	dif := now.Sub(frombrowser)
	//log.Println("checkauthorization token  ", "frombrowser"add , frombrowser, "now", now)

	if dif.Minutes() > 15 {
		log.Println("checkauthorization token expired ")
		return false
	}
	SessionCategories = jsondat.UserAuthCategories
	SessionAction = jsondat.UserAuthAction
	return true
}
func login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	pw := r.FormValue("pw")
	userid := r.FormValue("userid")
	userpassword := r.FormValue("userpassword")
	userdata = "\n"
	userdata = userdata + "=====================================\n"
	userdata = userdata + "Forwarded: " + r.Header.Get("X-Forwarded-For") + "\n"
	userdata = userdata + "RemoteAddr: " + r.RemoteAddr + "\n"
	userdata = userdata + "Agent: " + r.UserAgent() + "\n"
	userdata = userdata + "PW: " + pw + "\n"
	userdata = userdata + "User: " + userid + "\n"
	userdata = userdata + "UserPassword: " + userpassword + "\n"
	userdata = userdata + "=====================================\n"
	log.Println(userdata)
	//log.Println("Login")
	if isbusy {
		w.Write([]byte(ibusy()))
		return
	}

	hashdata, readerr := os.ReadFile(HashLocation)
	if readerr != nil {
		log.Println("ERROR Preferences readerr ", readerr)
	}
	var iserrors = false

	//log.Println("pw ", MyPrefs.Password)
	pwh, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("pword")), bcrypt.DefaultCost)
	config.PasswordHash = string(pwh)
	if err != nil {
		iserrors = true
		log.Println("Login Cant Open Password Hash", err)
	}

	// Comparing the password with the hash
	errpw := bcrypt.CompareHashAndPassword([]byte(hashdata), []byte(r.FormValue("pword")))
	if errpw != nil {
		iserrors = true
		log.Println("Login Bad Hash ", errpw)
	}
	if iserrors {
		w.Write([]byte(ilogon()))
		return
	}
	myuser := config.UserGetbyID(userid)
	var uajson = UserSessionJSON{}
	uajson.UserToken = config.GetDateTime("0h")[0:19]
	uajson.UserIPA = r.RemoteAddr
	uajson.UserPasswordHash = myuser.Userpasswordhash
	uajson.UserAuthCategories = myuser.Userauthcategories
	SessionCategories = myuser.Userauthcategories
	uajson.UserAuthAction = myuser.Userauthaction
	//log.Println("myuser caction", myuser.Userauthaction)
	SessionAction = myuser.Userauthaction
	js, err := json.Marshal(uajson)
	if err != nil {
		log.Println("UserJSON ERR ", err)
		w.Write([]byte(ilogon()))
		return
	}
	// check aut token expiry

	tobrowser := config.Encrypt(string(js), MySecret)
	//r.Header.Set("Authorization", authtoken)

	w.Write([]byte(ibuilder(tobrowser)))
}
func arrayhas(a []string, v string) bool {
	//log.Println("arrayhas: a", a, "v", v)
	for index, value := range a {
		log.Println("Index: ", index, value)
		if v == value {
			return true
		}
	}

	return false
}
func ibuilder(authtoken string) string {

	var s bytes.Buffer
	s.WriteString("<!DOCTYPE html>\n")
	s.WriteString("<html lang=\"en\">\n")
	s.WriteString("<head>\n")
	s.WriteString(" <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />\n")
	s.WriteString(" <meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\" />\n")
	s.WriteString(" <title>Broadcast Radio Web Interface</title>\n")
	s.WriteString(" <link rel=\"icon\" type=\"image/x-icon\" href=\"logo.png\">\n")
	s.WriteString("</head>\n")
	s.WriteString("<body>\n")
	s.WriteString("	   <table>\n")
	s.WriteString("	     <tr style=\"background-color: #ddd;\">\n")
	s.WriteString("	       <th colspan=\"1\"><img src=\"logo.png\" alt=\"Broadcast Radio\" style=\"width:128px;height:128px;\"></th>\n")
	s.WriteString("	       <th colspan=\"7\">Broadcat Web Interface</th>\n")
	s.WriteString("	     </tr>\n")

	if arrayhas(SessionAction, "ALL") || arrayhas(SessionAction, "Upload/Download") {
		s.WriteString("	     <tr>\n")
		s.WriteString("           <form enctype=\"multipart/form-data\" action=\"" + config.WebAddress + "/upload\" method=\"post\">\n")
		s.WriteString("              <td colspan=\"1\"><input type=\"file\" name=\"stub\" />\n</td>")
		s.WriteString("              <td colspan=\"1\"><input type=\"submit\" value=\"Upload stub.zip\" style=\"color: #4c14e477;\"/>\n</td>")
		s.WriteString("              <td colspan=\"1\"><input type=\"hidden\" name=\"Authorization\" id=\"Authorization\" value=\"" + authtoken + "\" />\n</td>")
		s.WriteString("	             <td colspan=\"5\">Upload and process a stub file. The file must be in zip format. Manually upload content with restrictions</td>\n")
		s.WriteString("          </form>\n")
		s.WriteString("	     </tr>\n")

		s.WriteString("	     <tr style=\"background-color: #11d4e277;\">\n")
		s.WriteString("           <form  action=\"" + config.WebAddress + "/download\" method=\"post\">\n")
		s.WriteString("              <td colspan=\"1\"><input type=\"submit\" value=\"Download stub.zip\" style=\"color: #4c14e477;\" /></td>\n")
		s.WriteString("              <td colspan=\"2\"><input type=\"hidden\" name=\"Authorization\" id=\"Authorization\" value=\"" + authtoken + "\" /></td>\n")
		s.WriteString("	             <td colspan=\"6\">Download a stub file.<br>The file is in zip format. Use this to build content</td>\n")
		s.WriteString("           </form>\n")
		s.WriteString("	     </tr>\n")
	}
	if arrayhas(SessionAction, "ALL") || arrayhas(SessionAction, "Category History") {
		s.WriteString("	     <tr>\n")
		s.WriteString("           <form  action=\"" + config.WebAddress + "/chart\" method=\"post\" target=\"_blank\">\n")
		s.WriteString("              <td colspan=\"1\"><select name=\"Days\" id=\"days\">")
		s.WriteString("                  <option value=\"7\">7 Days</option>")
		s.WriteString("                  <option value=\"14\">14 Days</option>")
		s.WriteString("                  <option value=\"28\">28 Days</option>")
		s.WriteString("                  </select></td>")
		s.WriteString("              <td colspan=\"1\"><select name=\"Categories\" id=\"categories\">")
		s.WriteString("                  <option value=\"ADS\">Advertising</option>")
		s.WriteString("                  <option value=\"PROMOS\">Promotions</option>")
		s.WriteString("                  <option value=\"NWS\">News Weather Sports</option>")
		s.WriteString("                  <option value=\"FILLTOTOH\">Top Of Hour</option>")
		s.WriteString("                  <option value=\"IMAGINGID\">Imaging Spots</option>")
		s.WriteString("                  <option value=\"DJ\">DJ Spots</option>")
		s.WriteString("                  </select></td>")
		s.WriteString("               <td colspan=\"1\"><input type=\"submit\" value=\"Category History\" style=\"color: #4c14e477;\" /></td>\n")
		s.WriteString("               <td colspan=\"1\"><input type=\"hidden\" name=\"Authorization\" id=\"Authorization\" value=\"" + authtoken + "\" /></td>\n")
		s.WriteString("           </form>\n")
		s.WriteString("	          <td colspan=\"1\">Produce a line chart using parameters</td>\n")
		s.WriteString("	     </tr>\n")
	}
	if arrayhas(SessionAction, "ALL") || arrayhas(SessionAction, "Traffic") {
		s.WriteString("	      <tr style=\"background-color: #11d4e277;\">\n")
		s.WriteString("         <form  action=\"" + config.WebAddress + "/TRAFICREPORT\" method=\"post\" target=\"_blank\">\n")
		s.WriteString("         <td colspan=\"1\"><select name=\"Days1\" id=\"days1\">")
		s.WriteString("             <option value=\"7\">7 Days</option>")
		s.WriteString("             <option value=\"14\">14 Days</option>")
		s.WriteString("             <option value=\"28\">28 Days</option>")
		s.WriteString("             <option value=\"90\">90 Days</option>")
		s.WriteString("             <option value=\"365\">365 Days</option>")
		s.WriteString("            </select></td>")
		s.WriteString("          <td colspan=\"1\"><select name=\"Categories1\" id=\"categories1\">")
		s.WriteString(config.TrafficGetAlbum("-2160"))
		s.WriteString("          </select></td>")
		s.WriteString("          <td colspan=\"1\"><input type=\"submit\" value=\"Ad Report\" style=\"color: #4c14e477;\" /></td>\n")
		s.WriteString("          <td colspan=\"1\"><input type=\"hidden\" name=\"Authorization\" id=\"Authorization\" value=\"" + authtoken + "\" /></td>\n")
		s.WriteString("         </form>\n")
		s.WriteString("	        <td colspan=\"1\">Produce a billing report using parameters</td>\n")
		s.WriteString("	     </tr>\n")
	}
	if arrayhas(SessionAction, "ALL") || arrayhas(SessionAction, "Chart") {
		s.WriteString("	      <tr>\n")
		s.WriteString("         <form  action=\"" + config.WebAddress + "/ADS\" method=\"post\" target=\"_blank\">\n")
		s.WriteString("         <td colspan=\"1\"><select name=\"Days\" id=\"days\">")
		s.WriteString("             <option value=\"7\">7 Days</option>")
		s.WriteString("             <option value=\"14\">14 Days</option>")
		s.WriteString("             <option value=\"28\">28 Days</option>")
		s.WriteString("             <option value=\"90\">90 Days</option>")
		s.WriteString("            </select></td>")
		s.WriteString("          <td colspan=\"1\"><select name=\"Categories1\" id=\"categories1\">")
		s.WriteString(config.TrafficGetAlbum("-2160"))
		s.WriteString("          </select></td>")
		s.WriteString("          <td colspan=\"1\"><input type=\"submit\" value=\"Ad History\" style=\"color: #4c14e477;\" /></td>\n")
		s.WriteString("          <td colspan=\"1\"><input type=\"hidden\" name=\"Authorization\" id=\"Authorization\" value=\"" + authtoken + "\" /></td>\n")
		s.WriteString("         </form>\n")
		s.WriteString("	        <td colspan=\"1\">Produce a line chart using parameters</td>\n")
		s.WriteString("	     </tr>\n")

		s.WriteString("	      <tr style=\"background-color: #11d4e277;\">\n")
		s.WriteString("        <form  action=\"" + config.WebAddress + "/counts\" method=\"post\" target=\"_blank\">\n")
		s.WriteString("          <td colspan=\"1\"><input type=\"submit\" value=\"Inventory counts\" style=\"color: #4c14e477;\" /></td>\n")
		s.WriteString("          <td colspan=\"2\"><input type=\"hidden\" name=\"Authorization\" id=\"Authorization\" value=\"" + authtoken + "\" /></td>\n")
		s.WriteString("        </form>\n")
		s.WriteString("	        <td colspan=\"5\">Produce a pie chart using parameters</td>\n")
		s.WriteString("	     </tr>\n")

		s.WriteString("	     <tr>\n")
		s.WriteString("  		<form  action=\"" + config.WebAddress + "/schedcounts\" method=\"post\" target=\"_blank\">\n")
		s.WriteString(" 	      <td colspan=\"1\"><input type=\"submit\" value=\"Schedule counts\" style=\"color: #4c14e477;\" /></td>\n")
		s.WriteString("  	      <td colspan=\"2\"><input type=\"hidden\" name=\"Authorization\" id=\"Authorization\" value=\"" + authtoken + "\" /></td>\n")
		s.WriteString("  		</form>\n")
		s.WriteString("	        <td colspan=\"5\">Produce a pie chart using parameters</td>\n")
		s.WriteString("	     </tr>\n")
	}
	if arrayhas(SessionAction, "ALL") || arrayhas(SessionAction, "Clear") {
		s.WriteString("	     <tr style=\"background-color: #ec292977;\">\n")
		s.WriteString("          <form  action=\"" + config.WebAddress + "/cleartraffic\" method=\"post\" target=\"_blank\">\n")
		s.WriteString("            <td colspan=\"1\"><input type=\"submit\" value=\"Clear Traffic Over 1 Year Old\" style=\"color: #04c14e477;\" /></td>\n")
		s.WriteString("           <td colspan=\"2\"><input type=\"hidden\" name=\"Authorization\" id=\"Authorization\" value=\"" + authtoken + "\" /></td>\n")
		s.WriteString("          </form>\n")
		s.WriteString("	        <td colspan=\"5\">Clear traffic counts </td>\n")
		s.WriteString("	     </tr>\n")
	}
	s.WriteString("	   </table>\n")
	s.WriteString("</body>\n")
	s.WriteString("</html>\n")

	return s.String()
}
func ilogon() string {
	log.Println("ilogon")
	var s bytes.Buffer
	s.WriteString("<!DOCTYPE html>\n")
	s.WriteString("<html lang=\"en\">\n")
	s.WriteString("<head>\n")
	s.WriteString(" <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />\n")
	s.WriteString(" <meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\" />\n")
	s.WriteString(" <title>Broadcast Web Interface</title>\n")
	s.WriteString("</head>\n")
	s.WriteString("<body>\n")
	s.WriteString("  <form action=\"" + config.WebAddress + "/login\" method=\"post\">\n")
	s.WriteString("	   <table>\n")
	s.WriteString("	     <tr>\n")
	s.WriteString("	       <th colspan=\"1\"><img src=\"logo.png\" alt=\"Broadcast Radio\"></th>\n")
	s.WriteString("	       <th colspan=\"3\">Broadcat Web Interface</th>\n")
	s.WriteString("	     </tr>\n")
	s.WriteString("	     <tr>\n")
	s.WriteString("	       <td colspan=\"1\"><label>System ID</label></td>\n")
	s.WriteString("	       <td colspan=\"1\"><input type=\"password\" id=\"pw\" name=\"pword\"></td>\n")
	s.WriteString("	     </tr>\n")
	s.WriteString("	     <tr>\n")
	s.WriteString("	       <td colspan=\"1\"><label>User ID</label></td>\n")
	s.WriteString("	       <td colspan=\"1\"><input type=\"password\" id=\"userid\" name=\"userid\"></td>\n")
	s.WriteString("	     </tr>\n")
	s.WriteString("	     <tr>\n")
	s.WriteString("	       <td colspan=\"1\"><label>User Password</label></td>\n")
	s.WriteString("	       <td colspan=\"1\"><input type=\"password\" id=\"userpassword\" name=\"userpassword\"></td>\n")

	s.WriteString("	       <td colspan=\"2\"><input type=\"submit\" value=\"Try Authentication\"></td>\n")
	s.WriteString("	     </tr>\n")

	s.WriteString("	   </table>\n")
	s.WriteString("  </form>\n")

	s.WriteString("</body>\n")
	s.WriteString("</html>\n")

	return s.String()
}
func ibusy() string {
	authtoken = ""
	var s bytes.Buffer
	s.WriteString("<!DOCTYPE html>\n")
	s.WriteString("<html lang=\"en\">\n")
	s.WriteString("<head>\n")
	s.WriteString(" <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />\n")
	s.WriteString(" <meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\" />\n")
	s.WriteString(" <title>Content Provider Logon</title>\n")
	s.WriteString("</head>\n")
	s.WriteString("<body>\n")
	s.WriteString("  <label> System in use try again later/label>\n")

	s.WriteString("</body>\n")
	s.WriteString("</html>\n")

	return s.String()
}
func istartup() string {
	authtoken = ""
	var s bytes.Buffer
	s.WriteString("<!DOCTYPE html>\n")
	s.WriteString("<html lang=\"en\">\n")
	s.WriteString("<head>\n")
	s.WriteString(" <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />\n")
	s.WriteString(" <meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\" />\n")
	s.WriteString(" <link rel=\"icon\" type=\"image/x-icon\" href=\"logo.png\">\n")
	s.WriteString(" <title>Authenticate</title>\n")
	s.WriteString("</head>\n")
	s.WriteString("<body>\n")
	s.WriteString("  \n")
	s.WriteString("  <form action=\"" + config.WebAddress + "/continue\" method=\"post\">\n")
	s.WriteString("  <table>\n")
	s.WriteString("  <tr>\n")
	s.WriteString("  <th colspan=\"1\"><img src=\"logo.png\" alt=\"Broadcast Radio\"></th>\n")
	s.WriteString("  <th colspan=\"2\">Broadcat Web Interface</th>\n")
	s.WriteString("  </tr>\n")
	s.WriteString("  <tr>\n")
	s.WriteString("  <td colspan=\"1\"><label>Authorized Users only</label></td>\n")
	s.WriteString("  <td colspan=\"3\"><input type=\"submit\" value=\"Continue\"></td>\n")
	s.WriteString("  </tr>\n")
	s.WriteString("  </table>\n")
	s.WriteString("  </form>\n")
	s.WriteString("</body>\n")
	s.WriteString("</html>\n")

	return s.String()
}
