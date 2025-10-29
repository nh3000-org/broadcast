package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nh3000-org/broadcast/config"
)

// INVID,INVTYPE,INVSCHEDROTATION,INVLOC,INVARTIST,INVDROPDATE,INVCHART,INVLENGTH,INVSTARTDATE,INVENDDATE,
// INVSPINS,INVSPINSCOUNT,INVORDER,INVLASTPLAYDATE,INVSPINYTD,INVTITLE,LASTUSERNAME,LASTCHANGEDATE,
type InvRecord struct {
	invid            string
	invtype          string
	invschedrotation string
	invloc           string
	invartist        string
	invdropdate      string
	invchart         string
	invlength        string
	invstartdate     string
	invenddate       string
	invspins         string
	invspinscount    string
	invorder         string
	invlastplaydate  string
	invspinsytd      string
	invtitle         string
	invlastuser      string
	invlastchange    string
}

func createInventory(data [][]string, importdir string) {
	log.Println("Starting Conversion")
	var hp = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}
	var dp = []string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}
	for i, line := range data {
		if i > 0 { // omit header line
			var rec InvRecord
			for j, field := range line {
				if j == 0 {
					rec.invid = field
				} else if j == 1 {
					rec.invtype = field
				} else if j == 2 {
					rec.invschedrotation = field
				} else if j == 3 {
					rec.invloc = field
				} else if j == 4 {
					rec.invartist = field
				} else if j == 5 {
					rec.invdropdate = field
				} else if j == 6 {
					rec.invchart = field
				} else if j == 7 {
					rec.invlength = field
				} else if j == 8 {
					rec.invstartdate = field
				} else if j == 9 {
					rec.invenddate = field
				} else if j == 10 {
					rec.invspins = field
				} else if j == 11 {
					rec.invspinscount = field
				} else if j == 12 {
					rec.invorder = field
				} else if j == 13 {
					rec.invlastplaydate = field
				} else if j == 14 {
					rec.invspinsytd = field
				} else if j == 15 {
					rec.invtitle = field
				} else if j == 16 {
					rec.invlastuser = field
				} else if j == 17 {
					rec.invlastchange = field
				}
			}
			var art = ""
			var song = ""
			if rec.invtype == "001" || rec.invtype == "002" || rec.invtype == "907" {
				artstring := strings.Split(rec.invartist, "-")
				if len(artstring) >= 2 {
					art = artstring[0]
					song = artstring[1]
				}
				if len(artstring) == 1 {
					art = artstring[0]
					song = artstring[0]
				}
				log.Println(rec.invid, "type:", rec.invtype, "artist:", art, "song:", song, "chart:", rec.invchart, "length:", rec.invlength)
				cat := ""
				if rec.invtype == "001" {
					cat = "RECURRENTS"
				}
				if rec.invtype == "002" {
					cat = "CURRENTS"
				}
				if rec.invtype == "907" {
					cat = "IMAGINGID"
				}
				length, _ := strconv.Atoi(rec.invlength)
				added := config.GetDateTime("0h")
				rowreturned := config.InventoryAdd(cat, art, song, "", length, "000000", "1999-01-01 00:00:00", "9999-01-01 00:00:00", hp, dp, 0, 0, "1999-01-01 00:00:00", added[0:19], 0, 0, 0, "CVT")
				row := strconv.Itoa(rowreturned)
				if row != "0" {
					songbytes, songerr := os.ReadFile(importdir + rec.invid + ".mp3")
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
					if cat == "CURRENTS" {
						rowreturned := config.InventoryGetRow(cat, art, song, "")
						if len(rowreturned) > 0 {
							//log.Println("importing intro", rowreturned)
							songbytes, songerr = os.ReadFile(importdir + rec.invid + "INTRO.mp3")
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
						}
					}
					if cat == "CURRENTS" {
						rowreturned := config.InventoryGetRow(cat, art, song, "")
						if len(rowreturned) > 0 {
							//log.Println("importing outro", rowreturned)
							songbytes, songerr = os.ReadFile(importdir + rec.invid + "OUTRO.mp3")
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
						}
					}

				}

			}

		}
	}
}
func main() {
	readPreferences()
	// read inventory.csv
	//process inventory types
	// 001 - recurrents
	// 002 - currents
	// 907 - imaging
	var filename = "INVENTORY.csv"
	var datadir = "/opt/wrrw/"
	csvdata, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to open input file "+filename, err)

	}
	csvReader := csv.NewReader(csvdata)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Println("CVT Read err", err)
	}
	createInventory(data, datadir)

	csvdata.Close()
}
var PreferencesLocation = "/home/oem/.config/fyne/org.nh3000.nh3000/preferences.json"

const MySecret string = "abd&1*~#^2^#s0^=)^^7%c34"

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

	config.NatsCaroot = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCaroot"]), MySecret)
	config.NatsClientkey = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCakey"]), MySecret)
	config.NatsClientcert = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCaclient"]), MySecret)
	config.NatsQueuePassword = config.Decrypt(fmt.Sprintf("%v", cfg["NatsQueuePassword"]), MySecret)
	//amm := strconv.Itoa(cfg["AdsMaxMinutes"])

	//log.Println("CONFIG AdsMaxMinutes", config.AdsMaxMinutes)
	//log.Println("NATS AUTH user", config.NatsServer, config.NatsUser, config.NatsUserPassword)
	config.NewNatsJS()
	config.NewPGSQL()
}
