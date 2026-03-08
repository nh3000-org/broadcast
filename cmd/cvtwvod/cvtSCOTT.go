package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"slices"
	//"github.com/nh3000-org/radio/config"
)

// scott findex directory index
type IndexRecord struct {
	Song   string
	Junk   string
	Artist string
	Junk2  string
	File   string
}

var musicIncludes = []string{"401", "402", "403", "404", "405", "406", "407", "408", "409", "410"}
var legalIncludes = []string{"ID4"}
var linersIncludes = []string{"LI"}
var promosIncludes = []string{"PR4", "SW4"}
var category string
var findexfile *os.File
var finxexfileerr error
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
	findexfile, findexfilerror := os.Open("FINDEX01.DAT")
	if findexfilerror != nil {
		log.Println("findexfile error reading", findexfilerror)
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
		if countforcurrents == 13 {
			countforcurrents = 1
			currentsselected = true
		}
		addInventory(ir,currentsselected)
		fbindex++
		//if count > 3 {
		//	os.Exit(0)
		//}
	}
}
func addInventory(rec IndexRecord, currents bool) {
	
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
			log.Println("read", info.Name(), category)
		}
		if category != "" {
			processIndex(path, station)
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
