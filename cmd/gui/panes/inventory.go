package panes

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/nh3000-org/broadcast/config"
)

var songbytes []byte
var songerr error
var pberr error
var shadowCategory string

func InventoryScreen(win fyne.Window) fyne.CanvasObject {

	config.FyneInventoryWin = win
	//Details := widget.NewLabel("")

	larow := widget.NewLabel("Row: ")
	edrow := widget.NewEntry()
	edrow.SetPlaceHolder("Automatically Assigned")
	edrow.Disable()
	gridrow := container.New(layout.NewGridLayoutWithRows(2), larow, edrow)

	lacategory := widget.NewLabel("Category: ")
	edcategory := widget.NewSelect(config.CategoriesToArray(), func(string) {})
	gridcategory := container.New(layout.NewGridLayoutWithRows(2), lacategory, edcategory)

	laartist := widget.NewLabel("Artist: ")
	edartist := widget.NewEntry()
	edartist.SetPlaceHolder("Enter Artist Name")
	gridartist := container.New(layout.NewGridLayoutWithRows(2), laartist, edartist)

	lasong := widget.NewLabel("Song: ")
	edsong := widget.NewEntry()
	edsong.SetPlaceHolder("Enter Song Name")
	gridsong := container.New(layout.NewGridLayoutWithRows(2), lasong, edsong)

	laalbum := widget.NewLabel("Album: ")
	edalbum := widget.NewEntry()
	edalbum.SetPlaceHolder("Enter Album Name")
	gridalbum := container.New(layout.NewGridLayoutWithRows(2), laalbum, edalbum)

	lalength := widget.NewLabel("Length: ")
	edlength := widget.NewEntry()
	edlength.SetPlaceHolder("Enter Song Length")
	edlength.SetText("0")
	edlength.Disable()
	gridlength := container.New(layout.NewGridLayoutWithRows(2), lalength, edlength)

	laorder := widget.NewLabel("Play Order: ")
	edorder := widget.NewEntry()
	edorder.SetText("000000")
	edorder.Disable()
	gridorder := container.New(layout.NewGridLayoutWithRows(2), laorder, edorder)

	laexpires := widget.NewLabel("Expires On: ")
	edexpires := widget.NewEntry()
	edexpires.SetText("9999-01-01T00:00:00Z")
	gridexpires := container.New(layout.NewGridLayoutWithRows(2), laexpires, edexpires)

	lastartson := widget.NewLabel("Starts On: ")
	edstartson := widget.NewEntry()
	so := config.GetDateTime("1")
	edstartson.SetText(so[0:19])
	gridstartson := container.New(layout.NewGridLayoutWithRows(2), lastartson, edstartson)

	laadstimeslot := widget.NewLabel("ADS Time Slot: ")
	edadstimeslot := widget.NewCheckGroup([]string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}, func([]string) {})
	edadstimeslot.Horizontal = true
	gridadstimeslot := container.New(layout.NewGridLayoutWithRows(2), laadstimeslot, edadstimeslot)

	laadsdayslot := widget.NewLabel("ADS Day Slot: ")
	edadsdayslot := widget.NewCheckGroup([]string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}, func([]string) {})
	edadsdayslot.Horizontal = true
	gridadsdayslot := container.New(layout.NewGridLayoutWithRows(2), laadsdayslot, edadsdayslot)

	laadsmaxspins := widget.NewLabel("ADS Max Spins Per Day: ")
	edadsmaxspins := widget.NewSelect([]string{"12", "24", "36", "48", "9999"}, func(string) {})
	gridadsmaxspins := container.New(layout.NewGridLayoutWithRows(2), laadsmaxspins, edadsmaxspins)

	laadsmaxspinsperhour := widget.NewLabel("ADS Max Spins Per Hour: ")
	edadsmaxspinsperhour := widget.NewSelect([]string{"1", "2", "3", "4"}, func(string) {})
	gridadsmaxspinsperhour := container.New(layout.NewGridLayoutWithRows(2), laadsmaxspinsperhour, edadsmaxspinsperhour)

	lalastplayed := widget.NewLabel("Last Played: ")
	edlastplayed := widget.NewEntry()
	edlastplayed.SetText("2000-01-01 00:00:00")
	gridlastplayed := container.New(layout.NewGridLayoutWithRows(2), lalastplayed, edlastplayed)

	ladateadded := widget.NewLabel("Date Added: ")
	eddateadded := widget.NewEntry()
	eddateadded.Disable()

	da := config.GetDateTime("1h")

	eddateadded.SetText(da[0:19])
	gridedateadded := container.New(layout.NewGridLayoutWithRows(2), ladateadded, eddateadded)

	laspinstoday := widget.NewLabel("Spins Today: ")
	edspinstoday := widget.NewEntry()
	edspinstoday.Disable()
	edspinstoday.SetText("0")
	gridspinstoday := container.New(layout.NewGridLayoutWithRows(2), laspinstoday, edspinstoday)

	laspinsweek := widget.NewLabel("Spins Weekly: ")
	edspinsweek := widget.NewEntry()
	edspinsweek.Disable()
	edspinsweek.SetText("0")
	gridspinsweek := container.New(layout.NewGridLayoutWithRows(2), laspinsweek, edspinsweek)

	lalinks := widget.NewLabel("Source Links: ")
	edlinks := widget.NewEntry()
	edlinks.SetPlaceHolder("Enter Website: ")
	gridlinks := container.New(layout.NewGridLayoutWithRows(2), lalinks, edlinks)

	laspinstotal := widget.NewLabel("Spins Total: ")
	edspinstotal := widget.NewEntry()
	edspinstotal.Disable()
	edspinstotal.SetText("0")

	edsongsz := widget.NewLabel("0")
	edintrosz := widget.NewLabel("0")
	edoutrosz := widget.NewLabel("0")

	gridspinstotal := container.New(layout.NewGridLayoutWithRows(2), laspinstotal, edspinstotal)

	importbutton := widget.NewButtonWithIcon("Import Stub", theme.UploadIcon(), func() {
		var imartist string
		var imsong string
		var imalbum string
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				return
			}
			if reader == nil {
				return
			}

			var imcategory string
			sp := reader.URI()

			sp1 := strings.Replace(sp.Path(), "file//", "", 1)
			startpath := strings.Replace(sp1, "/README.txt", "", 1)
			os.Chdir(startpath)

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
							songbytes, songerr = os.ReadFile(imimportdir)
							if songerr != nil {
								config.Send("messages."+config.NatsAlias, "Put Bucket Song Read Error", config.NatsAlias)
							}
							if songerr == nil {
								pberr = config.PutBucket("mp3", row, songbytes)
								if pberr == nil {
									songbytes = []byte("")
								}
								if pberr != nil {
									config.Send("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
								}
							}
						}
					}

					if strings.HasSuffix(cat, "INTRO.mp3") {
						rowreturned := config.InventoryGetRow(imcategory, imartist, imsong, imalbum)
						if len(rowreturned) > 0 {
							//log.Println("importing intro", rowreturned)
							songbytes, songerr = os.ReadFile(imimportdir)
							if songerr != nil {
								log.Println("messages."+config.NatsAlias, "Put Bucket Intro Read Error", config.NatsAlias)
								config.Send("messages."+config.NatsAlias, "Put Bucket Intro Read Error", config.NatsAlias)
							}
							if songerr == nil {
								//log.Println("putting intro", rowreturned+"INTRO")
								pberr = config.PutBucket("mp3", rowreturned+"INTRO", songbytes)
								if pberr == nil {
									songbytes = []byte("")
								}
								if pberr != nil {
									log.Println("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
									config.Send("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
								}
							}
						}
					}
					if strings.HasSuffix(cat, "OUTRO.mp3") {
						rowreturned := config.InventoryGetRow(imcategory, imartist, imsong, imalbum)
						if len(rowreturned) > 0 {
							//log.Println("importing outro", rowreturned)
							songbytes, songerr = os.ReadFile(imimportdir)
							if songerr != nil {
								log.Println("messages."+config.NatsAlias, "Put Bucket Outro Read Error", config.NatsAlias)
								config.Send("messages."+config.NatsAlias, "Put Bucket Outro Read Error", config.NatsAlias)
							}
							if songerr == nil {
								//log.Println("putting outro", rowreturned+"OUTRO")
								pberr = config.PutBucket("mp3", rowreturned+"OUTRO", songbytes)
								if pberr == nil {
									songbytes = []byte("")
								}
								if pberr != nil {
									log.Println("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
									config.Send("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
								}
							}
						}
					}

				}
				return nil
			})
			if walkstuberr != nil {
				log.Println("messages.IMPORT", "Inventory Walk Err FileInfo "+walkstuberr.Error(), "onair")
				config.Send("messages.IMPORT", "Inventory Walk Err FileInfo "+walkstuberr.Error(), "onair")
			}
			win.SetTitle("Importing Complete")
			config.InventoryGet()
			config.FyneInventoryList.Refresh()
		}, win)

		fd.Show()

	})
	openSong := widget.NewButtonWithIcon("Load Song ", theme.UploadIcon(), func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				return
			}
			if reader == nil {
				return
			}
			var song = reader

			if songerr == nil {
				if strings.HasSuffix(song.URI().String(), "mp3") {
					songbytes, songerr := os.ReadFile(strings.Replace(song.URI().String(), "file://", "", -1))
					if songerr != nil {
						config.Send("messages."+config.NatsAlias, "Put Bucket Song Read Error", config.NatsAlias)
					}
					pberr = config.PutBucket("mp3", edrow.Text, songbytes)
					if pberr == nil {
						edsongsz.SetText(strconv.Itoa(len(songbytes)))
						songbytes = []byte("")
					}
					if pberr != nil {
						config.Send("messages."+config.NatsAlias, "Put Bucket ", config.NatsAlias)
						log.Fatal("ERROR PutBucket mp3" + song.URI().String() + ": " + pberr.Error())
					}
				}
				if strings.HasSuffix(song.URI().String(), "mp4") {
					log.Println("Starting" + song.URI().String() + ".ts")
					os.Remove(song.URI().String() + ".ts")
					win.SetTitle("Converting mp4" + song.URI().String() + ".ts")
					cmd := exec.Command("ffmpeg", "-i", song.URI().String(), "-c", "copy", "-f", "mpegts", song.URI().String()+".ts")
					err := cmd.Run()
					if err != nil {
						log.Println("Converting" + song.URI().String() + ".ts" + " " + err.Error())
						log.Fatal("ERROR ffmpeg mp4" + song.URI().String() + ": " + pberr.Error())
					}
					time.Sleep(5 * time.Second) // Wait for the process to finish
					win.SetTitle("Sending " + song.URI().String() + ".ts")
					songbytes, songerr := os.ReadFile(strings.Replace(song.URI().String()+".ts", "file://", "", -1))
					if songerr != nil {
						log.Fatal("ERROR ReadFile mp4 " + song.URI().String() + ": " + err.Error())
						config.Send("messages."+config.NatsAlias, "Put Bucket Song Read Error mp4 ", config.NatsAlias)
					}
					pberr = config.PutBucket("mp4", edrow.Text, songbytes)
					if pberr != nil {
						config.Send("messages."+config.NatsAlias, "Put Bucket mp4 ", config.NatsAlias)
						log.Fatal("ERROR PutBucket mp4 " + song.URI().String() + ": " + err.Error())
					}
					if pberr == nil {
						edsongsz.SetText(strconv.Itoa(len(songbytes)))
						log.Println("Size" + song.URI().String() + ".ts " + strconv.Itoa(len(songbytes)))
						songbytes = []byte("")
					}
					os.Remove(song.URI().String() + ".ts")
					win.SetTitle("Sent " + song.URI().String() + ".ts")
				}
				if pberr != nil {
					config.Send("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
				}

			}

		}, win)

		fd.Show()

	})

	openSongIntro := widget.NewButtonWithIcon("Load Song Intro ", theme.UploadIcon(), func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				return
			}
			if reader == nil {
				return
			}

			song := reader

			songbytes, songerr = os.ReadFile(strings.Replace(song.URI().String(), "file://", "", -1))
			if songerr != nil {
				config.Send("messages."+config.NatsAlias, "Put Bucket Read Error "+songerr.Error(), config.NatsAlias)
			}
			if songerr == nil {
				pberr = config.PutBucket("mp3", edrow.Text+"INTRO", songbytes)

				if pberr == nil {
					edintrosz.SetText(strconv.Itoa(len(songbytes)))
					songbytes = []byte("")
				}
				if pberr != nil {
					config.Send("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
				}
			}

		}, win)

		fd.Show()
	})

	openSongOutro := widget.NewButtonWithIcon("Load Song Outro ", theme.UploadIcon(), func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				return
			}
			if reader == nil {
				return
			}

			song := reader
			songbytes, songerr = os.ReadFile(strings.Replace(song.URI().String(), "file://", "", -1))
			if songerr != nil {
				config.Send("messages."+config.NatsAlias, "Put Bucket Read Error "+songerr.Error(), config.NatsAlias)
			}
			if songerr == nil {
				pberr = config.PutBucket("mp3", edrow.Text+"OUTRO", songbytes)
				if pberr == nil {
					edoutrosz.SetText(strconv.Itoa(len(songbytes)))
					songbytes = []byte("")
				}
				if pberr != nil {
					config.Send("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
				}

			}
		}, win)

		fd.Show()
		runtime.GC()
	})
	gridfile := container.New(layout.NewGridLayoutWithColumns(3), openSong, openSongIntro, openSongOutro)
	bucket := "mp3"
	if strings.Contains(edcategory.Selected, "VIDEO") {
		bucket = "mp4"
	}
	openSongsz := strconv.Itoa(int(config.GetBucketSize(bucket, edrow.Text)))
	edsongsz.SetText("Song Size:" + openSongsz)
	openSongIntrosz := strconv.Itoa(int(config.GetBucketSize(bucket, edrow.Text+"INTRO")))
	edintrosz.SetText("Intro Size:" + openSongIntrosz)
	openSongOutrosz := strconv.Itoa(int(config.GetBucketSize(bucket, edrow.Text+"OUTRO")))
	edoutrosz.SetText("Intro Size:" + openSongOutrosz)
	gridfilesz := container.New(layout.NewGridLayoutWithColumns(3), edsongsz, edintrosz, edoutrosz)

	saveaddbutton := widget.NewButtonWithIcon("Add Inventory Item", theme.ContentCopyIcon(), func() {
		length, _ := strconv.Atoi(edlength.Text)
		today, _ := strconv.Atoi(edspinstoday.Text)
		week, _ := strconv.Atoi(edspinsweek.Text)
		total, _ := strconv.Atoi(edspinstotal.Text)
		addsmaxspins, _ := strconv.Atoi(edadsmaxspins.Selected)
		addsmaxspinsperhour, _ := strconv.Atoi(edadsmaxspinsperhour.Selected)
		rowreturned := config.InventoryAdd(edcategory.Selected, edartist.Text, edsong.Text, edalbum.Text, length, edorder.Text, edstartson.Text, edexpires.Text, edadstimeslot.Selected, edadsdayslot.Selected, addsmaxspins, addsmaxspinsperhour, edlastplayed.Text, eddateadded.Text, today, week, total, edlinks.Text)
		row := strconv.Itoa(rowreturned)
		edrow.SetText(row)
		openSong.Enable()
		openSongIntro.Enable()
		openSongOutro.Enable()
		config.InventoryGet()
		config.FyneInventoryList.Refresh()

	})

	List := widget.NewTable(func() (int, int) {
		return len(config.InventoryStore), 5
	}, func() fyne.CanvasObject {
		return container.NewMax(widget.NewLabel("template11"), widget.NewIcon(nil))
	}, func(id widget.TableCellID, o fyne.CanvasObject) {
		l := o.(*fyne.Container).Objects[0].(*widget.Label)
		l.Show()
		switch id.Col {

		case 0: // rowid
			l.SetText(strconv.Itoa(config.InventoryStore[id.Row].Row))
		case 1: // category
			l.SetText(config.InventoryStore[id.Row].Category)
		case 2: // artist
			l.SetText(config.InventoryStore[id.Row].Artist)
		case 3: // song
			l.SetText(config.InventoryStore[id.Row].Song)
		case 4: // album
			l.SetText(config.InventoryStore[id.Row].Album)
		}
	})
	List.SetColumnWidth(0, 32)
	List.SetColumnWidth(1, 170)
	List.SetColumnWidth(2, 128)
	List.SetColumnWidth(3, 96)
	List.SetColumnWidth(4, 96)
	List.SetColumnWidth(5, 32)

	config.FyneInventoryList = List
	List.OnSelected = func(id widget.TableCellID) {

		config.SelectedInventory = id.Row

		//Details.SetText("[" + config.InventoryStore[id.Row].Category + "] " + config.InventoryStore[id].Artist + " - " + config.InventoryStore[id].Song)

		edrow.SetText(strconv.Itoa(config.InventoryStore[id.Row].Row))
		edrow.Disable()
		edcategory.SetSelected(config.InventoryStore[id.Row].Category)
		shadowCategory = config.InventoryStore[id.Row].Category
		edartist.SetText(config.InventoryStore[id.Row].Artist)
		edsong.SetText(config.InventoryStore[id.Row].Song)
		edalbum.SetText(config.InventoryStore[id.Row].Album)
		edlength.SetText(strconv.Itoa(config.InventoryStore[id.Row].Songlength))
		edlength.Disable()
		edorder.SetText(config.InventoryStore[id.Row].Rndorder)
		edorder.Disable()

		edstartson.SetText(config.InventoryStore[id.Row].Startson)
		edexpires.SetText(config.InventoryStore[id.Row].Expireson)
		edadstimeslot.SetSelected(config.InventoryStore[id.Row].AdsTimeSlots)
		edadsdayslot.SetSelected(config.InventoryStore[id.Row].AdsDaySlots)
		edadsmaxspins.SetSelected(strconv.Itoa(config.InventoryStore[id.Row].AdsMaxSpins))
		edadsmaxspinsperhour.SetSelected(strconv.Itoa(config.InventoryStore[id.Row].AdsMaxSpinsPerHour))
		eddateadded.SetText(config.InventoryStore[id.Row].Dateadded)
		edlastplayed.SetText(config.InventoryStore[id.Row].Lastplayed)
		edlinks.SetText(config.InventoryStore[id.Row].Sourcelink)
		bucket := "mp3"
		if strings.Contains(edcategory.Selected, "VIDEO") {
			bucket = "mp4"
		}
		edsongsz.SetText("Song Size: " + strconv.Itoa(int(config.GetBucketSize(bucket, edrow.Text))))
		edintrosz.SetText("Intro Size: " + strconv.Itoa(int(config.GetBucketSize(bucket, edrow.Text+"INTRO"))))
		edoutrosz.SetText("Outro Size: " + strconv.Itoa(int(config.GetBucketSize(bucket, edrow.Text+"OUTRO"))))
		edspinstoday.SetText(strconv.Itoa(config.InventoryStore[id.Row].Spinstoday))
		edspinsweek.SetText(strconv.Itoa(config.InventoryStore[id.Row].Spinsweek))
		edspinstotal.SetText(strconv.Itoa(config.InventoryStore[id.Row].Spinstotal))

		deletebutton := widget.NewButtonWithIcon("Delete Inventory Item", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			config.InventoryDelete(myrow)
			config.InventoryGet()
			config.FyneInventoryList.Refresh()
		})
		savebutton := widget.NewButtonWithIcon("Save Inventory Item", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			var length, _ = strconv.Atoi(edlength.Text)
			var today, _ = strconv.Atoi(edspinstoday.Text)
			var week, _ = strconv.Atoi(edspinsweek.Text)
			var total, _ = strconv.Atoi(edspinstotal.Text)
			var maxspins, _ = strconv.Atoi(edadsmaxspins.Selected)
			var maxspinsperhour, _ = strconv.Atoi(edadsmaxspinsperhour.Selected)
			config.InventoryUpdate(myrow, edcategory.Selected, edartist.Text, edsong.Text, edalbum.Text, length, edorder.Text, edstartson.Text, edexpires.Text, edadstimeslot.Selected, edadsdayslot.Selected, maxspins, maxspinsperhour, edlastplayed.Text, eddateadded.Text, today, week, total, edlinks.Text)
			config.InventoryGet()
			config.FyneInventoryList.Refresh()
			if shadowCategory == "CURRENTS" {
				if edcategory.Selected != "CURRENTS" {
					// delete intro/outro
					config.DeleteBucket("MP3", strconv.Itoa(myrow)+"INTRO.mp3")
					config.DeleteBucket("MP3", strconv.Itoa(myrow)+"OUTRO.mp3")
				}
			}

		})
		databox := container.NewVBox(

			gridrow,
			gridcategory,
			gridartist,
			gridsong,
			gridalbum,
			gridfile,
			gridfilesz,
			gridlength,
			gridorder,
			gridstartson,
			gridexpires,
			gridadstimeslot,
			gridadsdayslot,
			gridadsmaxspins,
			gridadsmaxspinsperhour,
			gridlastplayed,
			gridedateadded,
			gridlinks,
			gridspinstoday,
			gridspinsweek,
			gridspinstotal,
			savebutton,
			deletebutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow("Manage Inventory Items")

		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()
		config.InventoryGet()
		config.FyneInventoryList.Refresh()
		List.Unselect(id)
	}
	//	importbutton := widget.NewButtonWithIcon("Import Inventory From Stub ", theme.ContentCopyIcon(), func() {
	//	})

	addbutton := widget.NewButtonWithIcon("Add", theme.ContentCopyIcon(), func() {

		edrow.SetText("")
		edartist.SetText("")
		edsong.SetText("")
		edalbum.SetText("")
		edalbum.SetText("")
		edsongsz.SetText("0")
		edintrosz.SetText("0")
		edoutrosz.SetText("0")
		edlength.SetText("0")
		edorder.SetText("000000")
		edstartson.SetText("2023-12-31 00:00:00")
		edexpires.SetText("9999-12-31 00:00:00")

		edlastplayed.SetText("1999-01-01 00:00:00")

		da := config.GetDateTime("1h")

		eddateadded.SetText(da[0:19])
		edspinstoday.SetText("0")
		edspinsweek.SetText("0")
		edspinstotal.SetText("0")
		databox := container.NewVBox(
			gridrow,
			gridcategory,
			gridartist,
			gridsong,
			gridalbum,
			gridfile,
			gridfilesz,
			gridlength,
			gridorder,
			gridstartson,
			gridexpires,
			gridadstimeslot,
			gridadsmaxspins,
			gridadsmaxspinsperhour,
			gridlastplayed,
			gridedateadded,
			gridlinks,
			gridspinstoday,
			gridspinsweek,
			gridspinstotal,
			saveaddbutton,
		)
		openSong.Disable()
		openSongIntro.Disable()
		openSongOutro.Disable()
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))

		dlg := fyne.CurrentApp().NewWindow("Manage Inventory Item")

		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()

	})
	exportstub := widget.NewButtonWithIcon("Export Stub", theme.DownloadIcon(), func() {
		where := config.CategoriesWriteStub(true)
		Errors.SetText(where)
		config.Send("messages.Export", where, config.NatsAlias)

	})
	syncdbtofs := widget.NewButtonWithIcon("Sync Db to FS", theme.ListIcon(), func() {
		config.InventoryGet()

		for i := range config.InventoryStore {
			if !config.TestBucket("mp3", strconv.Itoa(config.InventoryStore[i].Row)) {
				config.InventoryDelete(config.InventoryStore[i].Row)
			}

		}

		config.InventoryGet()
		config.FyneInventoryList.Refresh()

	})
	topbox := container.NewGridWithColumns((4), addbutton, importbutton, exportstub, syncdbtofs)
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
