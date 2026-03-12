package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	//"strconv"
	"github.com/go-audio/wav"
	"github.com/nh3000-org/broadcast/config"
)

// scott findex directory index
type IndexRecord struct {
	Song   string
	Junk   string
	Artist string
	Junk2  string
	File   string
	Length string
}

var musicIncludes = []string{"401"}
var legalIncludes = []string{"ID4"}
var linersIncludes = []string{"LI4", "JI4"}
var promosIncludes = []string{"PR4", "SW4"}
var category string
var findexfile *os.File
var findexfilerror error
var fb = make([]byte, 184)
var fbindex int64 = 1
var continuereading = true
var count = 1
var countforcurrents = 1
var currentsselected = false

//var goodfile string

func processIndex(path, station string) {
	// read the FINDEX01.DAT file block size 179
	log.Println("processIndex", path, station)
	os.Chdir(path)
	//if !info.IsDir() && category != "" {
	findexfile, findexfilerror = os.Open("FINDEX01.DAT")
	if findexfilerror != nil {
		log.Println("findexfile error reading", findexfilerror)
		return
	}
	continuereading = true
	for continuereading {
		var seekstart int64 = fbindex * 183
		_, fbseelerror := findexfile.Seek(seekstart, 0)
		if fbseelerror != nil {
			log.Println("findexfile seek error ", fbseelerror)
		}

		_, readerr := findexfile.Read(fb)
		//log.Println("read", rb)
		if readerr != nil {
			log.Println("findexfile read error ", readerr)
			continuereading = false
			return
		}

		//continuereading = errors.Is(readerr, io.EOF)
		//}
		//if continuereading {
		var ir = IndexRecord{}
		sfb := string(fb)
		/* 	Song   [45]byte
		Junk   [12]byte
		Artist [32]byte
		Junk2  [12]byte
		File   [78]byte */
		ir.Song = sfb[0:43]
		ir.Junk = sfb[44:55]
		ir.Artist = sfb[56:90]
		ir.Length = sfb[91:96]
		ir.File = sfb[103:183]
		ir.Song = strings.ReplaceAll(ir.Song, "\x00", "")
		ir.Artist = strings.ReplaceAll(ir.Artist, "\x00", "")
		ir.File = strings.ReplaceAll(ir.File, "\x00", "")
		log.Println("c:", count, "s:", ir.Song, "a:", ir.Artist, "f:", ir.File, "j1:", "l:", ir.Length)
		count++
		countforcurrents++
		currentsselected = false
		// write currents intro/outro
		if countforcurrents > 13 {
			countforcurrents = 1
			currentsselected = true
		}
		addInventory(ir, currentsselected, path, ir.File)

		fbindex++
		//if count > 3 {
		//	os.Exit(0)
		//}
	}
}

var dircount = 1
var suf = ""
var nm = ""

func processDirectory(path, station, category string) {
	// read the FINDEX01.DAT file block size 179
	log.Println("processDirectory", path, station, category)
	os.Chdir(path)
	dircount = 1
	walkfileerr := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		suf = strconv.Itoa(dircount)
		nm = category + suf
		if strings.HasPrefix(info.Name(), "SP") {
			log.Println("c:", dircount, "s:", nm, "a:", nm, "f:", nm)
			ir := IndexRecord{}
			ir.Artist = nm
			ir.Song = nm
			ir.File = path + "/" + info.Name()
			ir.Length = " 0:30"
			addInventory(ir, currentsselected, path, ir.File)
			dircount++
		}
		return nil
	})
	if walkfileerr != nil {
		log.Println("walkfileerr", walkfileerr)
	}
}

var hp = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}
var dp = []string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}
var tmpcategory string

func addInventory(rec IndexRecord, currentsselected bool, path string, file string) {
	var m = rec.Length[1:2]
	var s = rec.Length[3:5]
	min, _ := strconv.Atoi(m)
	sec, _ := strconv.Atoi(s)
	var l = min*60 + sec
	log.Println("AddInventory ", path, file, currentsselected, path, "Length", rec.Length, min, sec, l)
	cddirer := os.Chdir(path)
	if cddirer != nil {
		log.Println("AddInventory cddirer", cddirer)
	}
	f, ferr := os.Open(file)
	if ferr != nil {
		log.Println("AddInventory ferr", ferr)
	}

	d := wav.NewDecoder(f)

	buf, err := d.FullPCMBuffer()
	if err != nil {
		panic(err)
	}
	os.Remove("/opt/radio/wvod.wav")
	out, _ := os.Create("/opt/radio/wvod.wav")

	e := wav.NewEncoder(out, 44100, 16, 2, 1)
	// Add metadata
	e.Metadata = &wav.Metadata{
		Title:    rec.Song,
		Artist:   rec.Artist,
		Comments: "WVOD",
		Genre:    "AAA",
		TrackNbr: "1",
	}

	if err := e.Write(buf); err != nil {
		panic(err)
	}
	if err := e.Close(); err != nil {
		panic(err)
	}

	added := config.GetDateTime("0h")
	tmpcategory = category
	if currentsselected {
		tmpcategory = "CURRENTS"
	}
	rowreturned := config.InventoryAdd(tmpcategory, rec.Artist, rec.Song, "WVOD", l, "000000", "1999-01-01 00:00:00", "9999-01-01 00:00:00", hp, dp, 0, 0, "1999-01-01 00:00:00", added[0:19], 0, 0, 0, "DIGITAL")
	row := strconv.Itoa(rowreturned)
	if row != "0" {
		songbytes, songerr := os.ReadFile("/opt/radio/wvod.wav")
		if songerr != nil {
			log.Println("messages."+"cvtwvod", "Put Bucket Song Read Error", "cvtwvod", songerr)
			config.Send("messages."+"cvtwvod", "Put Bucket Song Read Error", "cvtwvod")
		}
		if songerr == nil {
			pberr := config.PutBucket("wav", row, songbytes)
			if pberr == nil {
				songbytes = []byte("")
			}
			if pberr != nil {
				log.Println("messages."+"cvtwvod", "Put Bucket Write Error", "cvtwvod", songerr)
				config.Send("messages."+"cvtwvod", "Put Bucket Write Error", "cvtwvod")
			}

			if currentsselected {
				pberr := config.PutBucket("wav", row+"INTRO", songbytes)
				if pberr == nil {
					songbytes = []byte("")
				}
				if pberr != nil {
					log.Println("messages."+"cvtwvod", "Put Bucket Write Error", "cvtwvod", songerr)
					config.Send("messages."+"cvtwvod", "Put Bucket Write Error", "cvtwvod")
				}
			}
			if currentsselected {
				pberr := config.PutBucket("wav", row+"OUTRO", songbytes)
				if pberr == nil {
					songbytes = []byte("")
				}
				if pberr != nil {
					log.Println("messages."+"cvtwvod", "Put Bucket Write Error", "cvtwvod", songerr)
					config.Send("messages."+"cvtwvod", "Put Bucket Write Error", "cvtwvod")
				}
			}

		}
	}
	f.Close()
	e.Close()
}

func readPath(startpath, station string) {

	os.Chdir(startpath)
	fbindex = 1
	walkfileerr := filepath.Walk(startpath, func(path string, info os.FileInfo, err error) error {
		category = ""
		continuereading = false
		if info.IsDir() {

			// determine the category
			if slices.Contains(musicIncludes, info.Name()) {
				category = "RECURRENTS"
			}
			if slices.Contains(legalIncludes, info.Name()) {
				category = "STATIONID"
			}
			if slices.Contains(linersIncludes, info.Name()) {
				category = "IMAGINGID"
			}
			if slices.Contains(promosIncludes, info.Name()) {
				category = "PROMOS"
			}
			//log.Println("read", info.Name(), category)
		}
		if category == "RECURRENTS" {
			//processIndex(path, station)
		}
		if category == "IMAGINGID" {
			processDirectory(path, station, category)
		}
		if category == "STATIONID" {
			processDirectory(path, station, category)
		}
		if category == "PROMOS" {
			processDirectory(path, station, category)
		}
		return nil
	})
	if walkfileerr != nil {
		log.Println("walkfileerr", walkfileerr)
	}
}
func main() {
	rootImport := flag.String("rootimport", "./", "-rootimport base directory of scott files")
	stationid := flag.String("stationid", "WVOD", "-stationid call letters of station")

	flag.Parse()
	log.Println("init", *rootImport, *stationid)
	readPreferences()
	readPath(*rootImport, *stationid)

}

var PreferencesLocation = "/home/oem/.config/fyne/org.nh3000.nh3000/preferences.json"

const MySecret string = "abd&1*~#^2^#s0^=)^^7%c34"

var erramm error

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

	amm := config.Decrypt(fmt.Sprintf("%v", cfg["AdsMaxMinutes"]), MySecret)
	config.AdsMaxMinutes, erramm = strconv.Atoi(amm)
	if erramm != nil {
		log.Println("CONFIG AdsMaxMinutes", amm, erramm)
	}
	//log.Println("CONFIG NatsBucketType", config.NatsBucketType)
	//log.Println("NATS AUTH user", config.NatsServer, config.NatsUser, config.NatsUserPassword)
	config.NewNatsJS()
	config.NewPGSQL()
}
