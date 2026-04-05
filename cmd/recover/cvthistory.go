// Package recover is used to restore SQL historical
// data after importing a stub of all the content
// generated from expcontent.go
//
// This provides restoring values from a restore
// path to the database such as inventory ad and traffic
// data [exp/expcontent]

/*
cvthistory restore historic values after recovery.

Use this after importing a stub into a fresh database build.
The metrics for a given item are restored to the dump values.
This is required to match the new row ids to the restored content.

Usage:

	  cvthistory [flags] [path...]

	  The flags are:

	    -v output debug info
		-t test run does not update data base
*/
package main

import (
	"bufio"

	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nh3000-org/broadcast/config"
)

// processTraffic - writes all the historical data
// to the database
// does not process duplicates
func processTraffic(ts TrafficStruct, station, verbose, test string) {
	cd := config.TrafficCheckDuplicate(ts.Playedon, ts.Album)
	if cd != 0 {
		log.Println("processTraffic Found Duplicate ", ts.Album, ts.Playedon, ts.Album, ts.Artist)
		return
	}
	if test == "false" {
		config.TrafficAdd(ts.Category, ts.Artist, ts.Song, ts.Album, ts.Playedon)
	}
	if verbose == "true" {
		log.Println("Traffic Added", ts.Category, ts.Artist, ts.Song, ts.Album, ts.Playedon)
	}
}

// processInventory - updates the existing content with
// historical data
var Admaxspins int
var Admaxspinsperhour int
var Spinstoday int
var SpinsPerWeek int
var Spinstotal int

func processInventory(is InventoryStruct, station, verbose, test string) {
	// lookup up item by artist, song and album
	row := config.InventoryGetRowRecovery(is.Artist, is.Song, is.Album)
	// update with dump values
	if row != "0" {
		//convert string to int
		Admaxspins, _ = strconv.Atoi(is.Admaxspins)
		Admaxspinsperhour, _ = strconv.Atoi(is.Admaxspinsperhour)
		Spinstoday, _ = strconv.Atoi(is.Spinstoday)
		SpinsPerWeek, _ = strconv.Atoi(is.Spinsweek)
		Spinstotal, _ = strconv.Atoi(is.Spinstotal)
		log.Println("CheckingInventory", is.Artist, is.Song)
		if test == "false" {
			config.InventoryUpdateRecovery(row, is.Length, is.Startson, is.Expireson, is.Adtimeslots, is.Addayslots, Admaxspins, Admaxspinsperhour, is.Lastplayed, is.Dateadded, Spinstoday, SpinsPerWeek, Spinstotal)
		}
	}
	if verbose == "true" {
		if row == "0" {
			log.Println("processInventory Database record not found", is.Artist, is.Song, is.Album)
		}
	}
}

// readSQl process the sql dump from postgresql
// only use traffic and inventory tables
// traffic import entire table
// inventory only import
var importType = "" // INVENTORY or TRAFFIC

type InventoryStruct struct {
	Artist            string
	Song              string
	Album             string
	Length            string
	Startson          string
	Expireson         string
	Adtimeslots       string
	Addayslots        string
	Admaxspins        string
	Admaxspinsperhour string
	Lastplayed        string
	Dateadded         string
	Spinstoday        string
	Spinsweek         string
	Spinstotal        string
}
type TrafficStruct struct {
	Category string
	Artist   string
	Song     string
	Album    string
	Playedon string
}

// readSQL - process partial updates to production database
// from a pgsql dump
func readSQL(rootimport string, station string, verbose string, test string) {
	if test == "true" {
		log.Println("readSQL rootImport:", rootimport, "station:", station, "verbose:", verbose, "test", test)
	}
	// determine start of stream file
	// looking for traffic or imventory
	// read the file
	inputfile, err := os.Open(rootimport)
	if err != nil {
		fmt.Println("readSQL file error:", err)
		return
	}

	scanner := bufio.NewScanner(inputfile)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "\\.") {
			importType = ""
		}
		if importType == "INVENTORY" {
			fields := strings.Split(scanner.Text(), "\t")
			i := InventoryStruct{}
			for index, value := range fields {
				if index == 2 { // artist
					i.Artist = value
				}
				if index == 3 { // song
					i.Song = value
				}
				if index == 4 { // artist
					i.Album = value
				}
				if index == 5 { // song
					i.Length = value
				}

				if index == 7 { // song
					i.Startson = value
				}
				if index == 8 { // expireson
					i.Expireson = value
				}
				if index == 9 { // adtimeslots
					i.Adtimeslots = value
				}
				if index == 10 { // addayslots
					i.Addayslots = value
				}
				if index == 11 { // ad max spins
					i.Admaxspins = value
				}
				if index == 12 { // ad max spins per hpur
					i.Admaxspinsperhour = value
				}
				if index == 14 { // last played
					i.Lastplayed = value
				}
				if index == 15 { // spins today
					i.Spinstoday = value
				}
				if index == 16 { // spins per week
					i.Spinsweek = value
				}
				if index == 17 { // spins total
					i.Spinstotal = value
				}

			}
			processInventory(i, station, verbose, test)
			//if verbose == "true" {
			//	fmt.Printf("%#v\n", i.Artist)
			//}

		}
		if importType == "TRAFFIC" {

			fields := strings.Split(scanner.Text(), "\t")
			i := TrafficStruct{}
			for index, value := range fields {
				if index == 1 { // category
					i.Category = value
				}
				if index == 2 { // song
					i.Artist = value
				}
				if index == 3 { // artist
					i.Song = value
				}
				if index == 4 { // album
					i.Album = value
				}
				if index == 5 { // played on
					i.Playedon = value
				}
				
			}
			processTraffic(i, station, verbose, test)
			if verbose == "true" {
				fmt.Printf("%#v\n", i)
			}

		}
		if strings.Contains(scanner.Text(), "COPY public.inventory") {
			importType = "INVENTORY"
		}
		if strings.Contains(scanner.Text(), "COPY public.traffic") {
			importType = "TRAFFIC"
		}
	}

}
func main() {
	rootImport := flag.String("rootimport", "./", "-rootimport base directory of SQL export from postgresql files")
	stationid := flag.String("stationid", "WVOD", "-stationid call letters of station")
	verbose := flag.String("v", "true", "-v print execution plan")
	test := flag.String("t", "true", "-t test execution no poatgresql update")

	flag.Parse()
	log.Println("init path:", *rootImport, "station:", *stationid, "verbose:", *verbose, "test", *test)
	readPreferences()
	readSQL(*rootImport, *stationid, *verbose, *test)

}

var erramm error

func readPreferences() {
	// read config preferences.json
	jsondata, readerr := os.ReadFile(config.PreferencesLocation)
	if readerr != nil {
		log.Println("ERROR Preferences readerr ", readerr)
	}
	// parse json
	var cfg map[string]any
	errunmarshal := json.Unmarshal(jsondata, &cfg)
	if errunmarshal != nil {
		log.Println("ERROR Preferences errunmarshal ", errunmarshal)
	}

	config.DBpassword = config.Decrypt(fmt.Sprintf("%v", cfg["DBPASSWORD"]), config.MySecret)

	config.DBaddress = config.Decrypt(fmt.Sprintf("%v", cfg["DBADDRESS"]), config.MySecret)

	config.DBuser = config.Decrypt(fmt.Sprintf("%v", cfg["DBUSER"]), config.MySecret)
	config.NatsBucketType = config.Decrypt(fmt.Sprintf("%v", cfg["NatsBucketType"]), config.MySecret)
	config.NatsCaroot = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCaroot"]), config.MySecret)
	config.NatsClientkey = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCakey"]), config.MySecret)
	config.NatsClientcert = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCaclient"]), config.MySecret)
	config.NatsQueuePassword = config.Decrypt(fmt.Sprintf("%v", cfg["NatsQueuePassword"]), config.MySecret)

	amm := config.Decrypt(fmt.Sprintf("%v", cfg["AdsMaxMinutes"]), config.MySecret)
	config.AdsMaxMinutes, erramm = strconv.Atoi(amm)
	if erramm != nil {
		log.Println("CONFIG AdsMaxMinutes", amm, erramm)
	}
	//log.Println("CONFIG NatsBucketType", config.NatsBucketType)
	//log.Println("NATS AUTH user", config.NatsServer, config.NatsUser, config.NatsUserPassword)
	config.NewNatsJS()
	config.NewPGSQL()
}
