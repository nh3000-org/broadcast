package main

import (
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

	"github.com/nh3000-org/radio/config"
)

var PreferencesLocation = "/home/oem/.config/fyne/org.nh3000.nh3000/preferences.json"

/*
	var NatsServer = "nats://nats.newhorizons3000.org:4222"

var NatsCaroot = "-----BEGIN CERTIFICATE-----\nMIID7zCCAtegAwIBAgIUaXAPxJvZRRdTq5RWlwxs1XYo+5kwDQYJKoZIhvcNAQEL\nBQAwgYAxCzAJBgNVBAYTAlVTMRAwDgYDVQQIEwdGbG9yaWRhMRIwEAYDVQQHEwlD\ncmVzdHZpZXcxGjAYBgNVBAoTEU5ldyBIb3Jpem9ucyAzMDAwMQwwCgYDVQQLEwNX\nV1cxITAfBgNVBAMTGG5hdHMubmV3aG9yaXpvbnMzMDAwLm9yZzAeFw0yMzEyMTkw\nMzA4MDBaFw0yODEyMTcwMzA4MDBaMIGAMQswCQYDVQQGEwJVUzEQMA4GA1UECBMH\nRmxvcmlkYTESMBAGA1UEBxMJQ3Jlc3R2aWV3MRowGAYDVQQKExFOZXcgSG9yaXpv\nbnMgMzAwMDEMMAoGA1UECxMDV1dXMSEwHwYDVQQDExhuYXRzLm5ld2hvcml6b25z\nMzAwMC5vcmcwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCrVIXA/SxU\n7GeW92UNyiPnQEZgbJIHHQ31AQE2C/vFdpEtv32uoX1SsDl5drWvBrMnd5zrw1tL\nOEPA26tk/ACfQYL0n0HfeutLLu8H9jUWNp8ziX6Qbgd01M+/BixobHQjyDMxulo4\nJU2VK6QBLs9VI6TIihEU2BZhc/XCD9QbWcikAif1JySpz93MjFv3pcQU8ci4vQ0T\nImaGnHesr1qDbX1NuFVuBOPavZ64sQ1RsZtH5CdD+RU772wQWUgkPkwyUn8QBwTS\ne9XV5DNQD5nGEXjKTgjrd9KRf9pmRDnf6gBLi2r6C/l6q2w3ItOOHARdK0mc9CYh\ngY1Nzl59vrWdAgMBAAGjXzBdMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8EBTAD\nAQH/MB0GA1UdDgQWBBR0qq9ueABC5RDsg/02FZFpBOR1hDAbBgNVHREEFDAShwTA\nqAAFhwTAqFjohwR/AAABMA0GCSqGSIb3DQEBCwUAA4IBAQBfdX0IMya9Dh9dHLJj\nnJZyb96htMWD5nuQQVBAu3ay+8O2GWj5mlsLJXAP2y7p/+3gyvHKTRDdJLux7N79\nHn6AYjmp3PCyZzuL1M/kHhSQxhxqJHGwjGXILt5pLovVkvkl4iukdxWJ5HAPsUGY\nO3QSDDFdoLflsG5VcrtdODm8uyxAjhMPAR2PXKfX8ABI79N7VKcbb98338fifrN8\n9H1r3BXcdsyhpH0gB0ZKJFSpMGWXlfudFEe9mXI9898xbEI2znqlYGhboVsuv5LM\nRESH2zXrkhmZyHqw0RtDROzyZOy5g1LcxbtVMn4w1LI4h3MDuE9B+Vud77A48qtA\ny+5x\n-----END CERTIFICATE-----\n"
var NatsClientcert = "-----BEGIN CERTIFICATE-----\nMIIEMTCCAxmgAwIBAgIUB7+OFX1LQrWtYMl5XIOXsOaLac0wDQYJKoZIhvcNAQEL\nBQAwgYAxCzAJBgNVBAYTAlVTMRAwDgYDVQQIEwdGbG9yaWRhMRIwEAYDVQQHEwlD\ncmVzdHZpZXcxGjAYBgNVBAoTEU5ldyBIb3Jpem9ucyAzMDAwMQwwCgYDVQQLEwNX\nV1cxITAfBgNVBAMTGG5hdHMubmV3aG9yaXpvbnMzMDAwLm9yZzAgFw0yMzEyMTkw\nMzA4MDBaGA8yMDUzMTIxMTAzMDgwMFowcjELMAkGA1UEBhMCVVMxEDAOBgNVBAgT\nB0Zsb3JpZGExEjAQBgNVBAcTCUNyZXN0dmlldzEaMBgGA1UEChMRTmV3IEhvcml6\nb25zIDMwMDAxITAfBgNVBAsTGG5hdHMubmV3aG9yaXpvbnMzMDAwLm9yZzCCASIw\nDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMWARyniHy8r342e3aKSsLDPwVMC\n2mRwuILP2JkXp5FllaFKnu/Z+0mF+iQlSchcC6DOcMQk00Cp/I8cCP865zyxPhqN\n2F2/qVItCU4+PTwe6ZnrfpJgXWwyk1hjS3vVNTT+idI5+pJgFH9YL0lbJ7q1UyPB\n+KP0x/c5T3K2Ec6U4uXhbVt/ePxFmsl1sHw6FE//XrA4EzbqCMEPCTcOfInvFrCJ\ny4/pAqjCxegT/1YDMNEdzmG8vg2tc3jPV+3GIAV3YL5nDE5mprHPEEDJtNQi+E4o\nXXXMobNhrJh9KJ59VbxTF8m5yM3b8fvof97OYhK0KYggplnTH+bhnYU9V5ECAwEA\nAaOBrTCBqjAOBgNVHQ8BAf8EBAMCBaAwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDAYD\nVR0TAQH/BAIwADAdBgNVHQ4EFgQUpffi3LSreerO756B/VnZkyyEVBIwHwYDVR0j\nBBgwFoAUdKqvbngAQuUQ7IP9NhWRaQTkdYQwNQYDVR0RBC4wLIIYbmF0cy5uZXdo\nb3Jpem9uczMwMDAub3JnhwR/AAABhwTAqAAFhwTAqFjoMA0GCSqGSIb3DQEBCwUA\nA4IBAQALlRqqW2HH4flFIgR/nh51gc/Hxv5xivhkzWUHHXRdltECSXknI4yBPchQ\n6Zsy0HZ7XQRlhQSIYd4Bp6eyHbny5t3JA978dHzpGJFCUVQDMY4yHLaCQgFJ+ESn\nwyyDWTRGA3cpEikL0B0ekDfqjWUEMTzmT/gnoSl0vM69nZDLZm1xMx1+EH+bpfFB\nRaVM6gKSAuFJmNYEL2e7JSags+3IHyVHkdo8GDlY//71Z4lxsFxFCF6xF9GDdAr2\niCA4OfydjiBSOz0eLJVgqkk1KGXtMqZXAojX62NrIWnFTW1Vzd46ekOHhq93B3tA\nkjWmHY/KdCZUjQSWss+YXgG4mI8c\n-----END CERTIFICATE-----\n"
var NatsClientkey = "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAxYBHKeIfLyvfjZ7dopKwsM/BUwLaZHC4gs/YmRenkWWVoUqe\n79n7SYX6JCVJyFwLoM5wxCTTQKn8jxwI/zrnPLE+Go3YXb+pUi0JTj49PB7pmet+\nkmBdbDKTWGNLe9U1NP6J0jn6kmAUf1gvSVsnurVTI8H4o/TH9zlPcrYRzpTi5eFt\nW394/EWayXWwfDoUT/9esDgTNuoIwQ8JNw58ie8WsInLj+kCqMLF6BP/VgMw0R3O\nYby+Da1zeM9X7cYgBXdgvmcMTmamsc8QQMm01CL4Tihddcyhs2GsmH0onn1VvFMX\nybnIzdvx++h/3s5iErQpiCCmWdMf5uGdhT1XkQIDAQABAoIBAB+Iu9QUJqaBetBB\n7WFnyo5wnY2DhxtCZDN+vDa1cCvm7F00bOwfAeBbY/UhfwZeq/yg+aBXwOMyQQEY\nmNcnsIQgSKo0u7c8Quy8BCBaD6zpwqKw1yTH/iKocJ5MPGEpSbWMbrUCTN/SN3Od\nwO8VfuJw0TWEYw7KpqLyo5zNNUqmczEO438CPGotbkFfzUqkumeUOsGWJFongyZY\na9EwpcTH2TkxuXum9SQVyLy+hSG/AEBp0cQPaRcoNh8sWYk43y5HrkIAqFo7dkMa\n9usAVMz9JCqIH2UNV04cDASFaiDMpYoD2hV2YHlL7/CQ7v5nb6OHT2A9aoSBOAfm\ns+dBzYECgYEA1l8+T9Xux73TCbFO2p7F094xSx4hhBZhaYpvzZoNN7iQdbdUVt2l\n1yHSoRgJUJMZlnKpMoNMLCxo34Lr3ww/TkIE/rrg10pqbqvojIDLCbi103EEB2v9\nWix8MSeOgFCa72T4lg9fDm5T493n4C5dade3LzZczUBF6dgmth3D+nMCgYEA69pa\nlob9n7eNXqDPk9kZUJV1jfLATC8eN4jupEiKfjnxEz9mUewvL/RF8kFhiS1ISC50\nKgM0v+isYBwwX00c7P02L6xCoGT35qOeoutEWVy/tYIHIHsD0jUBBsdnpQVNf58l\n9DDy2hZrpUwrsVHylVHpufBgKOfxgP2Jr3qD0OsCgYEAn4vzTGfkdzSIRMZ58awJ\ngE32Ufny5+PgTDSEUXk+LSJoIbR4SM5eB2dc5BiHljhk6twboUSnBJlo1DEUa8Up\nuIzaOtvLS3BPFl9LjIaulmWqrduHLB7rSJmjNNJD9KwJI/L6MHTwQkVKmmUllmvr\nikLKS5EiMICNiCUfaptsqJECgYEApYaSqzBEUdK1oeMErAPis16hqSTkdtNexqUQ\nrzXGFP6/Rb3qJra3C1XJvVLLjEW+hAIuPsoPPFyklbNS85+gHGc9n0mrXPxfy3ur\nuzWYu4rPdSizrcUIEoBmnwZVpEhLcrUUIwQzfIHdvJ3v0DvuH4PkoD2mjy7xnJDU\nD9bRKk8CgYAqK1lY5waFR0u3eFIPnrV4ATHXYuxcup2DCF+KJ6qwc4nNI6OB/ovU\nttiVZGr1rca42+XdWUQL5ufPFuKymeLbsuVzabbGKi+4RMvL+TIuorYtJRUPF+C7\nA9jlMeckpTZvl0yn5s3lC817N27B+U0M/jGow8sO0NtjBiImuTC5dg==\n-----END RSA PRIVATE KEY-----\n"

var NatsAdmin = "natsadmin"
var NatsUser = "natsoperator"
var NatsUserPassword = "hjscr44iod"
var NatsUserCommands = "natscommands"
var NatsUserCommandsPassword = "PASSWORD"
var NatsUserEvents = "natsevents"
var NatsUserEventsPassword = "PASSWORD"
var NatsUserDevices = "natsdevices"
var NatsUserDevicesPassword = "PASSWORD"
var NatsUserAuthorizations = "natsauthorizations"
var NatsUserAuthorizationsPassword = "PASSWORD"
var NatsQueuePassword = "987654321098765432109876"
*/
var KeyAes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}  // must be 16 bytes
var KeyHmac = []byte{36, 45, 53, 21, 87, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05} // must be 16 bytes
const MySecret string = "abd&1*~#^2^#s0^=)^^7%c34"

//var DBaddress = "db.newhorizons3000.org:5432/radio?sslmode=verify-ca"
//var DBuser = "postgres"
//var DBpassword = "postgres"

func uploadFile(w http.ResponseWriter, r *http.Request) {

	importHome := "/opt/radio/stub.zip"

	log.Println("File Upload Endpoint Hit for User", importHome)

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	pmuerr := r.ParseMultipartForm(10 << 20)
	if pmuerr != nil {
		log.Println("File Upload r.FormFile", pmuerr)
		w.Write([]byte("File Upload Parse Error r.FormFile"))
	}

	file, handler, reqerr := r.FormFile("stub")
	if reqerr != nil {
		w.Write([]byte("File Upload Error r.FormFile"))
		log.Println("File Upload r.FormFile", reqerr)
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

	//	cat := importHome
	//cat := strings.Replace(path, removepath, "", 1)
	//imimportdir := userHome + "/stub"

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
			rowreturned := config.InventoryAdd(imcategory, imartist, imsong, imalbum, length, "000000", "2023-12-31 00:00:00", "9999-12-31 00:00:00", "1999-01-01 00:00:00", added, today, week, total, "Stub")
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
		return nil
	})
	if walkstuberr != nil {
		config.Send("messages.IMPORT", "Inventory Walk Err FileInfo "+walkstuberr.Error(), "onair")
	}

	// read all of the contents of our uploaded file into a

	log.Println("Successfully Processed stub File")
}
func downloadFile(w http.ResponseWriter, r *http.Request) {

	importHome := "/opt/radio/blankstub"
	config.CategoriesWriteStub(false)
	os.Remove("/opt/radio/stub.zip")
	os.Chdir(importHome)
	cmd := exec.Command("zip", "-r", "/opt/radio/stub.zip", "stub")
	out, err := cmd.Output()
	if err != nil {
		w.Write([]byte("ZIP could not run command"))
		log.Println("ZIP could not run command: ", err, importHome)
	} else {
		//w.Write([]byte(string(out)))
		log.Println("ZIP Output: ", string(out))
	}
	//df, errdf := os.Open("/opt/radio/blankstub/stub.zip")
	//if errdf != nil {
	//	log.Println("Download Could Not Open file stub.zip ")
	//}
	hl, _ := os.ReadFile("/opt/radio/stub.zip")

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
	http.HandleFunc("/config", configFile)
	http.HandleFunc("/download", downloadFile)
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Waiting for Input")
	setupRoutes()
}
