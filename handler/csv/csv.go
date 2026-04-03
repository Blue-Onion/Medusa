package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	record "github.com/Blue-Onion/MahilAi/handler/Record"
)


func ExportToCsv(cam string, date string) {
	records, err := getData(cam, date)
	if err != nil {
		log.Fatal(err.Error())
	}

	now := time.Now().Format("2006-01-02_15-04-05")

	var name string

	if cam == "" && date == "" {
		name = fmt.Sprintf("all_events_%s.csv", now)

	} else if cam == "" && date != "" {
		name = fmt.Sprintf("all_events_%s.csv", date)

	} else if cam != "" && date == "" {
		name = fmt.Sprintf("%s_events_%s.csv", cam, now)

	} else {
		name = fmt.Sprintf("%s_events_%s.csv", cam, date)
	}

	err = createCsv(name, *records)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("CSV file exported:", name)
}
func getData(cam string,date string)(*[]record.Records,error){
	record,err:=record.ReadEvent(cam,date)
	if err!=nil{
		return nil,err
	}
	return &record,nil
}
func createCsv(fileName string,records []record.Records)error{
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write([]string{"Camera", "Time", "Event", "Confidence"})
	if err != nil {
		return err
	}
	for _, r := range records {
		row := []string{
			r.Camera,
			r.Time,
			r.Event,
			strconv.FormatFloat(r.Confidence, 'f', 2, 64),
		}
		err = writer.Write(row)
		if err != nil {
			return err
		}
	}
	return nil
}
