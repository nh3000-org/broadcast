package cvt

import (
	"encoding/csv"
	"log"
	"os"
)

// INVID,INVTYPE,INVSCHEDROTATION,INVLOC,INVARTIST,INVDROPDATE,INVCHART,INVLENGTH,INVSTARTDATE,INVENDDATE,
// INVSPINS,INVSPINSCOUNT,INVORDER,INVLASTPLAYDATE,INVSPINYTD,INVTITLE,LASTUSERNAME,LASTCHANGEDATE,
type InvRecord struct {

}
func main() {
	// read inventory.csv
	//process inventory types
	// 001 - recurrents
	// 002 - currents
	// 907 - imaging
	var filename = "inventory.csv"
	csvdata, err := os.Open(filename)
    if err != nil {
        log.Fatal("Unable to read input file " + filename, err)
		    csvReader := csv.NewReader(csvdata)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filename, err)
    }
	records.

    return records
    }
}