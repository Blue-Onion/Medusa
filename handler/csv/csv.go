package csv

import (
	"fmt"
	"log"

	record "github.com/Blue-Onion/MahilAi/handler/Record"
)


func DownloadCsv(cam string,date string){
	records,err:=getData(cam,date)
	if err!=nil{
		log.Fatal(err.Error())
	}
	fmt.Println(records)
}
func getData(cam string,date string)(*[]record.Records,error){
	record,err:=record.ReadEvent(cam,date)
	if err!=nil{
		return nil,err
	}
	return &record,nil
}
func createCsv([]record.Records){

}
