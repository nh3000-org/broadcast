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

	//"time"

	"github.com/nh3000-org/broadcast/config"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/google/uuid"
)

var PreferencesLocation = "/home/oem/.config/fyne/org.nh3000.nh3000/preferences.json"

var authtoken = ""

var isbusy = false

var userdata = ""
var KeyAes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}  // must be 16 bytes
var KeyHmac = []byte{36, 45, 53, 21, 87, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05} // must be 16 bytes
const MySecret string = "abd&1*~#^2^#s0^=)^^7%c34"

func chart(w http.ResponseWriter, r *http.Request) {

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
		log.Println("Setting Range", dt[0:10], dt[5:10])
	}
	dt := config.GetDateTime("0h")
	xdates = append(xdates, dt[0:10])
	xdesc = append(xdesc, dt[5:10])
	log.Println("Setting Range", dt[0:10], dt[5:10])
	c := r.FormValue("Categories")
	log.Println("Categories", c)

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
	line.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(true)}))
	line.Render(w)
}
func uploadFile(w http.ResponseWriter, r *http.Request) {
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

	if authtoken != r.FormValue("Authorization") {
		log.Println("File Upload Authorization", authtoken, "form", r.FormValue("Authorization"))
		w.Write([]byte(ilogon()))
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
	isbusy = true
	log.Println("Download Stub")
	pmuerr := r.ParseForm()
	if pmuerr != nil {
		log.Println("File Download", pmuerr)
		w.Write([]byte("File Download Parse Error r.FormFile"))

	}

	if authtoken != r.FormValue("Authorization") {
		log.Println("File Download Authorization")
		isbusy = false
		w.Write([]byte(ilogon()))
		return
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

	config.WebPassword = config.Decrypt(fmt.Sprintf("%v", cfg["WEBPASSWORD"]), MySecret)

	config.DBaddress = config.Decrypt(fmt.Sprintf("%v", cfg["DBADDRESS"]), MySecret)
	log.Println(config.DBaddress)

	config.DBuser = config.Decrypt(fmt.Sprintf("%v", cfg["DBUSER"]), MySecret)
	config.NatsAlias = config.Decrypt(fmt.Sprintf("%v", cfg["NatsAlias"]), MySecret)
	config.NatsCaroot = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCaroot"]), MySecret)
	config.NatsClientkey = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCakey"]), MySecret)
	config.NatsClientcert = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCaclient"]), MySecret)
	config.NatsQueuePassword = config.Decrypt(fmt.Sprintf("%v", cfg["NatsQueuePassword"]), MySecret)

	log.Println("NATS AUTH user", config.NatsServer, config.NatsUser, config.NatsUserPassword)
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
	http.HandleFunc("/login", login)
	http.HandleFunc("/config", configFile)
	http.HandleFunc("/download", downloadFile)
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/chart", chart)
	err := http.ListenAndServeTLS(":9000", "server.crt", "server.key", nil)
	if err != nil {
		log.Println("SSL ERROR ", err)
	}
}

func main() {
	fmt.Println("Waiting for Input")
	setupRoutes()
}

func login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	userdata = "\n"
	userdata = userdata + "=====================================\n"
	userdata = userdata + "Forwarded: " + r.Header.Get("X-Forwarded-For") + "\n"
	userdata = userdata + "RemoteAddr: " + r.RemoteAddr + "\n"
	userdata = userdata + "Password: " + r.FormValue("pword") + "\n"
	userdata = userdata + "Agent: " + r.UserAgent() + "\n"
	userdata = userdata + "=====================================\n"
	log.Println(userdata)
	log.Println("Login")
	if isbusy {
		w.Write([]byte(ibusy()))
		return
	}
	if config.WebPassword != r.FormValue("pword") {
		w.Write([]byte(ilogon()))
		return
	}
	authtoken = r.RemoteAddr + "-" + uuid.New().String()

	w.Write([]byte(ibuilder()))
}

func ibuilder() string {

	var s bytes.Buffer
	s.WriteString("<!DOCTYPE html>\n")
	s.WriteString("<html lang=\"en\">\n")
	s.WriteString("<head>\n")
	s.WriteString(" <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />\n")
	s.WriteString(" <meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\" />\n")
	s.WriteString(" <title>Content Provider</title>\n")
	s.WriteString("</head>\n")
	s.WriteString("  <form enctype=\"multipart/form-data\" action=\"" + config.WebAddress + "/upload\" method=\"post\">\n")
	s.WriteString("    <input type=\"file\" name=\"stub\" />\n")
	s.WriteString("    <input type=\"submit\" value=\"Upload stub.zip\" />\n")
	s.WriteString("    <input type=\"hidden\" name=\"Authorization\" id=\"Authorization\" value=\"" + authtoken + "\" />\n")
	s.WriteString("  </form>\n")
	s.WriteString("  <hr>\n")
	s.WriteString("  <form  action=\"" + config.WebAddress + "/download\" method=\"post\">\n")
	s.WriteString("    <input type=\"submit\" value=\"Download stub.zip\" />\n")
	s.WriteString("    <input type=\"hidden\" name=\"Authorization\" id=\"Authorization\" value=\"" + authtoken + "\" />\n")
	s.WriteString("  </form>\n")
	s.WriteString("  <form  action=\"" + config.WebAddress + "/chart\" method=\"post\">\n")
	s.WriteString("  <hr>\n")
	s.WriteString("  <label for=\"days\">History:</label>")
	s.WriteString("  <select name=\"Days\" id=\"days\">")
	s.WriteString("    <option value=\"7\">7 Days</option>")
	s.WriteString("    <option value=\"14\">14 Days</option>")
	s.WriteString("   <option value=\"28\">28 Days</option>")
	s.WriteString("  </select>")
	s.WriteString("  <label for=\"catgories\">Choose a Category:</label>")
	s.WriteString("  <select name=\"Categories\" id=\"categories\">")
	s.WriteString("    <option value=\"ADS\">Advertising</option>")
	s.WriteString("    <option value=\"PROMOS\">Promotions</option>")
	s.WriteString("   <option value=\"NWS\">News Weather Sports</option>")
	s.WriteString("   <option value=\"IMAGINGID\">Imaging Spots</option>")
	s.WriteString("   <option value=\"DJ\">DJ Spots</option>")
	s.WriteString("  </select>")

	s.WriteString("    <input type=\"submit\" value=\"Line Chart\" />\n")
	s.WriteString("    <input type=\"hidden\" name=\"Authorization\" id=\"Authorization\" value=\"" + authtoken + "\" />\n")
	s.WriteString("  </form>\n")
	s.WriteString("  <hr>\n")
	s.WriteString("</body>\n")
	s.WriteString("</html>\n")

	return s.String()
}
func ilogon() string {
	var s bytes.Buffer
	s.WriteString("<!DOCTYPE html>\n")
	s.WriteString("<html lang=\"en\">\n")
	s.WriteString("<head>\n")
	s.WriteString(" <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />\n")
	s.WriteString(" <meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\" />\n")
	s.WriteString(" <title>Content Provider Logon</title>\n")
	s.WriteString("</head>\n")
	s.WriteString("  <form action=\"" + config.WebAddress + "/login\" method=\"post\">\n")
	s.WriteString("    <label for=\"pword\"> Password:</label>\n")
	s.WriteString("    <input type=\"text\" id=\"pw\" name=\"pword\"><br><br>\n")
	s.WriteString("    <input type=\"submit\" value=\"Try Password\">\n")
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
