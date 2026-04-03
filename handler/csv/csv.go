package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	record "github.com/Blue-Onion/MahilAi/handler/Record"
)


func ExportToCsv(cam string,date string){
	records,err:=getData(cam,date)
	if err!=nil{
		log.Fatal(err.Error())
	}
	err=createCsv(*records)
	if err!=nil{
		log.Fatal(err.Error())
	}
	fmt.Println("File is downloaded")
}
func getData(cam string,date string)(*[]record.Records,error){
	record,err:=record.ReadEvent(cam,date)
	if err!=nil{
		return nil,err
	}
	return &record,nil
}
func createCsv(records []record.Records)error{
	file, err := os.Create("events.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// header
	err = writer.Write([]string{"Camera", "Time", "Event", "Confidence"})
	if err != nil {
		return err
	}

	// rows
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
