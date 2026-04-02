package csv

import record "github.com/Blue-Onion/MahilAi/handler/Record"


func DownloadCsv(){

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
