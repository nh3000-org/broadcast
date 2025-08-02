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

	"github.com/google/uuid"
)

var PreferencesLocation = "/home/oem/.config/fyne/org.nh3000.nh3000/preferences.json"

var authtoken = ""

var isbusy = false

var userdata = ""
var KeyAes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}  // must be 16 bytes
var KeyHmac = []byte{36, 45, 53, 21, 87, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05} // must be 16 bytes
const MySecret string = "abd&1*~#^2^#s0^=)^^7%c34"

func uploadFile(w http.ResponseWriter, r *http.Request) {
	isbusy = true
	importHome := "/opt/radio/stub.zip"

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
		log.Println("File Upload Authorization")
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

	// Create a destination file
	dst, _ := os.Create(importHome)
	defer dst.Close()

	// Upload the file to the destination path
	nb_bytes, _ := io.Copy(dst, file)

	fmt.Println("File uploaded successfully", nb_bytes)
	w.Write([]byte("File uploaded successfully"))
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern

	os.RemoveAll("/opt/radio/stub")
	cmd := exec.Command("unzip", "-d", "/opt/radio", importHome)
	out, err := cmd.Output()
	if err != nil {
		w.Write([]byte("UNZIP could not run command"))
		log.Println("UNZIP could not run command: ", err, importHome)
	} else {
		w.Write([]byte(string(out)))
		log.Println("Output: ", string(out))
	}

	var imartist string
	var imsong string
	var imalbum string

	var imcategory string
	sp := "/opt/radio/stub"
	os.Chdir(sp)
	startpath := strings.Replace(sp, "/README.txt", "", 1)
	walkstuberr := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {

		//removepath := "/opt/radio/stub/"
		//cat := strings.Replace(path, removepath, "", 1)
		//imimportdir := startpath + "/" + cat
		removepath := startpath + "/"
		cat := strings.Replace(path, removepath, "", 1)
		imimportdir := startpath + "/" + cat
		w.Write([]byte("Upload file " + imimportdir))
		log.Println("uploadfile ", imimportdir)
		//imimportdir := removepath + cat
		if info.IsDir() {
			imcategory = cat
		}

		if strings.HasSuffix(cat, "mp3") {
			rmcat := imcategory + "/"
			songfull := strings.ReplaceAll(path, rmcat, "")
			songunparsed := strings.ReplaceAll(songfull, ".mp3", "")
			result := strings.Split(songunparsed, "-")
			if len(result) == 0 {
				w.Write([]byte("Unparsed" + songunparsed))
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
			length, _ := strconv.Atoi("0")
			today, _ := strconv.Atoi("0")
			week, _ := strconv.Atoi("0")
			total, _ := strconv.Atoi("0")

			da := time.Now()
			added := "YYYY-MM-DD 00:00:00"
			added = strings.Replace(added, "YYYY", strconv.Itoa(da.Year()), 1)
			m := strconv.Itoa(int(da.Month()))
			if len(m) == 1 {
				m = "0" + m
			}
			added = strings.Replace(added, "MM", m, 1)
			d := strconv.Itoa(int(da.Day()))
			if len(d) == 1 {
				d = "0" + d
			}
			added = strings.Replace(added, "DD", d, 1)
			added = strings.Replace(added, "YYYY", strconv.Itoa(da.Year()), 1)
			m = strconv.Itoa(int(da.Month()))
			if len(m) == 1 {
				m = "0" + m
			}
			added = strings.Replace(added, "MM", m, 1)
			d = strconv.Itoa(int(da.Day()))
			if len(d) == 1 {
				d = "0" + d
			}
			added = strings.Replace(added, "DD", d, 1)
			adstimeslots := []string{}
			rowreturned := config.InventoryAdd(imcategory, imartist, imsong, imalbum, length, "000000", "2023-12-31 00:00:00", "9999-12-31 00:00:00", adstimeslots, 12, "1999-01-01 00:00:00", added, today, week, total, "Stub")

			row := strconv.Itoa(rowreturned)
			if row == "0" {
				w.Write([]byte("Inventory Not Added" + imcategory + "-" + imartist + "-" + imalbum))
			}
			if row != "0" {
				songbytes, songerr := os.ReadFile(imimportdir)
				if songerr != nil {
					w.Write([]byte("Read Error" + imimportdir))
					config.Send("messages."+config.NatsAlias, "Put Bucket Song Read Error", config.NatsAlias)
				}
				if songerr == nil {
					pberr := config.PutBucket("mp3", row, songbytes)
					if pberr == nil {

						w.Write([]byte("Imported" + imcategory + "-" + imartist + "-" + imalbum))
						fmt.Println("Imported", imcategory, imartist, imsong, imalbum)
					}
					if pberr != nil {
						w.Write([]byte("Not Imported" + imcategory + "-" + imartist + "-" + imalbum + " " + pberr.Error()))
						config.Send("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
					}
				}
				if strings.HasSuffix(cat, "INTRO.mp3") {
					songbytes, songerr = os.ReadFile(imimportdir)
					if songerr != nil {
						config.Send("messages."+config.NatsAlias, "Put Bucket Intro Read Error", config.NatsAlias)
					}
					if songerr == nil {
						pberr := config.PutBucket("mp3", row, songbytes)
						if pberr == nil {
							fmt.Println("Imported INTRO", imcategory, imartist, imsong, imalbum)
						}
						if pberr != nil {
							config.Send("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
						}
					}
				}
				if strings.HasSuffix(cat, "OUTRO.mp3") {
					songbytes, songerr = os.ReadFile(imimportdir)
					if songerr != nil {
						config.Send("messages."+config.NatsAlias, "Put Bucket Outro Read Error", config.NatsAlias)
					}
					if songerr == nil {
						pberr := config.PutBucket("mp3", row, songbytes)

						if pberr == nil {
							fmt.Println("Imported OUTRO", imcategory, imartist, imsong, imalbum)
						}
						if pberr != nil {
							config.Send("messages."+config.NatsAlias, "Put Bucket Write Error", config.NatsAlias)
						}
					}
				}
			}
		}
		log.Println("Unsuccessfully Processed stub File")
		log.Println(userdata)
		log.Println("Upload File")
		isbusy = false
		return nil
	})
	if walkstuberr != nil {
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
		log.Println("ZIP could not run command: ", err3, importHome)
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
	http.HandleFunc("/login", login)
	http.HandleFunc("/config", configFile)
	http.HandleFunc("/download", downloadFile)
	http.HandleFunc("/upload", uploadFile)
	err := http.ListenAndServeTLS(":9000", "server.pem", "cert.pem", nil)
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
	s.WriteString("  </form>\n")
	s.WriteString("  <form  action=\"" + config.WebAddress + "/download\" method=\"post\">\n")
	s.WriteString("    <input type=\"submit\" value=\"Download stub.zip\" />\n")
	s.WriteString("    <input type=\"hidden\" name=\"Authorization\" id=\"Authorization\" value=\"" + authtoken + "\" />\n")
	s.WriteString("  </form>\n")
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
