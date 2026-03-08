package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	//"strconv"
	//"github.com/nh3000-org/broadcast/config"
)

// scott findex directory index
type IndexRecord struct {
	Song   string
	Junk   string
	Artist string
	Junk2  string
	File   string
}

var musicIncludes = []string{"401"}
var legalIncludes = []string{"ID4"}
var linersIncludes = []string{"LI4","JI4"}
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
		ir.Junk2 = sfb[91:102]
		ir.File = sfb[103:183]
		log.Println("c:", count, "s:", ir.Song, "a:", ir.Artist, "f:", ir.File)
		count++
		countforcurrents++
		currentsselected = false
		// write currents intro/outro 
		if countforcurrents == 13 {
			countforcurrents = 1
			currentsselected = true
		}
		//addInventory(ir, currentsselected,path)
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
			dircount++
		}
		return nil
	})
	if walkfileerr != nil {
		log.Println("walkfileerr", walkfileerr)
	}
}

/*
	 func addInventory(rec IndexRecord, currents bool,path string) {
		// read the file
		// add the meta data

		added := config.GetDateTime("0h")
		rowreturned := config.InventoryAdd(cat, art, song, "", int(lengthFloat), "000000", "1999-01-01 00:00:00", "9999-01-01 00:00:00", hp, dp, 0, 0, "1999-01-01 00:00:00", added[0:19], 0, 0, 0, rec.invchart)
		row := strconv.Itoa(rowreturned)
		if row != "0" {
			songbytes, songerr := os.ReadFile(importdir + rec.invid + ".mp3")
			if songerr != nil {
				log.Println("messages."+"cvtwrrw", "Put Bucket Song Read Error", "cvtwrrw", songerr)
				config.Send("messages."+"cvtwrrw", "Put Bucket Song Read Error", "cvtwrrw")
			}
			if songerr == nil {
				pberr := config.PutBucket("mp3", row, songbytes)
				if pberr == nil {
					songbytes = []byte("")
				}
				if pberr != nil {
					log.Println("messages."+"cvtwrrw", "Put Bucket Write Error", "cvtwrrw", songerr)
					config.Send("messages."+"cvtwrrw", "Put Bucket Write Error", "cvtwrrw")
				}
			}
		}
	}
*/
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
			processIndex(path, station)
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
	readPath(*rootImport, *stationid)
	// read the findex
	// get the data
	// add it to nats and db

}
