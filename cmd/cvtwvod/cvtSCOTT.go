package main

import (
	"errors"
	"flag"
	"io"
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
var fb = make([]byte, 179)
var fbindex int64
var continuereading = true
var count = 1

func readPath(startpath, station string) {

	os.Chdir(startpath)
	fbindex = 1
	walkfileerr := filepath.Walk(startpath, func(path string, info os.FileInfo, err error) error {
		category = ""
		continuereading = true
		if info.IsDir() {
			log.Println("read", info.Name())
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
		}
		// read the FINDEX01.DAT file block size 179

		if !info.IsDir() && category != "" {
			log.Println("read2")
			findexfile, findexfilerror := os.Open("FINDEX01.DAT")
			if findexfilerror != nil {
				log.Println("findexfile error reading", findexfilerror)
			}
			var seekstart int64 = fbindex * 179
			_, fbseelerror := findexfile.Seek(seekstart, 1)
			if fbseelerror != nil {
				log.Println("findexfile seek error ", fbseelerror)
			}

			rb, readerr := findexfile.Read(fb)
			log.Println("read", rb)
			if readerr != nil {
				log.Println("findexfile read error ", readerr)
			}

			continuereading = errors.Is(readerr, io.EOF)
		}
		for continuereading {
			var ir = IndexRecord{}
			sfb := string(fb)
			/* 	Song   [45]byte
			Junk   [12]byte
			Artist [32]byte
			Junk2  [12]byte
			File   [78]byte */
			ir.Song = sfb[1:45]
			ir.Junk = sfb[46:58]
			ir.Artist = sfb[59:91]
			ir.Junk2 = sfb[92:104]
			ir.File = sfb[105:179]
			log.Println(fb, ir.Song, ir.Junk, ir.Artist, ir.Junk2, ir.File)
			count++
			fbindex++
			if count > 3 {
				os.Exit(0)
			}
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
