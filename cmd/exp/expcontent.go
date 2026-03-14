package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"strconv"
	"strings"
	"time"

	"github.com/nh3000-org/broadcast/config"
)

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
	log.Println(config.DBaddress)

	config.DBuser = config.Decrypt(fmt.Sprintf("%v", cfg["DBUSER"]), MySecret)
	config.NatsBucketType = config.Decrypt(fmt.Sprintf("%v", cfg["NatsBucketType"]), MySecret)
	config.NatsCaroot = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCaroot"]), MySecret)
	config.NatsClientkey = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCakey"]), MySecret)
	config.NatsClientcert = config.Decrypt(fmt.Sprintf("%v", cfg["NatsCaclient"]), MySecret)
	config.NatsQueuePassword = config.Decrypt(fmt.Sprintf("%v", cfg["NatsQueuePassword"]), MySecret)

	//log.Println("CONFIG NatsBucketType", config.NatsBucketType)
	//log.Println("NATS AUTH user", config.NatsServer, config.NatsUser, config.NatsUserPassword)
	config.NewNatsJS()
	config.NewPGSQL()
}

var userHome string

func main() {

	exportPath := flag.String("exportpath", "/", "-exportpath ")

	flag.Parse()
	userHome = *exportPath
	log.Println(userHome)
	readPreferences()
	CategoriesWriteStub(true)
}

var count = 0

func CategoriesWriteStub(withinventory bool) string {

	if withinventory {
		userHome = userHome + "contentstub/stub"
	}
	if !withinventory {
		userHome = userHome + "blankstub/stub"
	}
	log.Println("CategoriesWriteStub User Home", userHome)
	os.Mkdir(userHome, 0755)
	/* 	db, dberr := NewPGSQL()
	   	if dberr != nil {
	   		log.Println("WriteCategories", dberr)
	   	} */
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := config.SQL.Pool.Acquire(ctxsql)
	log.Println("CategoriesWriteStub Writing Categories to Stub backup:", withinventory)

	err4 := os.RemoveAll(userHome)
	if err4 != nil {
		log.Println("CategoriesWriteStub Remove Stub", err4)
	}

	err3 := os.MkdirAll(userHome, os.ModePerm)
	if err3 != nil {
		log.Println("CategoriesWriteStub Get Categories row for Stub", err3)
	}

	//os.WriteFile(userHome+"/README.txt", []byte(instructions), os.ModePerm)
	rows, rowserr := conn.Query(ctxsql, "select * from categories order by id")
	var rowid int
	var id string
	var desc string
	for rows.Next() {
		err := rows.Scan(&rowid, &id, &desc)
		if err != nil {
			log.Println("CategoriesWriteStub Get Categories row for Stub", err)
		}
		log.Println(count, "CategoriesWriteStub Writing Stub", userHome+"/"+id)
		err2 := os.Mkdir(userHome+"/"+id, os.ModePerm)
		if err2 != nil {
			log.Println("CategoriesWriteStub Create Stub", err2)
		}
		if err2 == nil {
			//get all inv items or category read and write
			if withinventory {
				ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 8000*time.Minute)
				conn, _ := config.SQL.Pool.Acquire(ctxsql)

				//ScheduleStore = make(map[int]ScheduleStruct)
				rows, rowserr := conn.Query(ctxsql, "select rowid,category,artist,song,album from inventory  where category = $1", id)

				for rows.Next() {
					var rowid int       // rowid
					var category string // category
					var artist string   // artist
					var song string     // song
					var album string    // Album

					err := rows.Scan(&rowid, &category, &artist, &song, &album)
					if err != nil {
						log.Println("CategoriesWriteStub Get Schedule row", err)
					}
					var invitem = artist + "-" + song + "-" + album
					if err != nil {
						log.Println(count, err, "CategoriesWriteStub Write Stub", userHome+"/"+id+"/"+invitem+".wav")
					}
					if err == nil {
						count++
						data := config.GetBucket(config.NatsBucketType, strconv.Itoa(rowid), "COPY")
						os.WriteFile(userHome+"/"+id+"/"+invitem+".wav", data, os.ModePerm)
						log.Println(count, "CategoriesWriteStub Write Stub", userHome+"/"+id+"/"+invitem+".wav")
					}

				}
				if rowserr != nil {
					log.Println("CategoriesWriteStub Get Schedule row error", rowserr)
				}
				conn.Release()
				ctxsqlcan()

			}
		}
	}
	if rowserr != nil {
		log.Println("CategoriesWriteStub Create Categories row error", rowserr)
	}
	if rowserr == nil {
		if strings.Contains(userHome, "blank") {
			os.RemoveAll(userHome + ".zip")
			cmd := exec.Command("zip", "-r", userHome+".zip", userHome)
			out, err := cmd.Output()
			if err != nil {
				log.Println("could not run command: ", err)
			} else {
				log.Println("Output: ", string(out))
			}
			// create new zip
		}
	}
	conn.Release()
	ctxsqlcan()
	return "Stub Written to: " + userHome
}
