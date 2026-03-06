package main

// scott findex directory index
type IndexRecord struct {
	Song string "fixed:1,44"
	Junk string "fixed:45,57"
	Artist string "fixed:58,90"
	Junk2 string  "fixed:91,103"
	File string  "fixed:104,182"
}
func main() {
	hourtimingstart = time.Now()
	schedDay := flag.String("schedday", "MON", "-schedday MON || TUE || WED || THU || FRI || SAT || SUN")
	stationId := flag.String("stationid", "WRRW", "-station WRRW")
	StationId = *stationId
	schedHour := flag.String("schedhour", "00", "-schedhour 00..23")
	schedhour = *schedHour
	Logging := flag.String("logging", "false", "-logging true || false")
	flag.Parse()



}