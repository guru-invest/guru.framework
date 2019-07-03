package sheets

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"io/ioutil"
	"strconv"
)

var SpreadSheet = spreadsheet.Spreadsheet{}

func connect(client_secret string, sheetId string)spreadsheet.Spreadsheet{
	data, err := ioutil.ReadFile(client_secret)
	checkError(err)
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	checkError(err)
	client := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet(sheetId)
	checkError(err)
	return spreadsheet
}

func Initialize(client_secret string, sheetId string){
	SpreadSheet = connect(client_secret, sheetId)
}

func AutoSumAllValuesOnColumn(sheetIndex int, column int) float64{
	res := 0.0
	sheet, err := SpreadSheet.SheetByIndex(uint(sheetIndex))
	checkError(err)
	for _, row := range sheet.Rows {
		for _, cell := range row {
			if cell.Column == uint(column){
				f, err := strconv.ParseFloat(cell.Value, 64)
				checkError(err)
				res += f
			}
		}
	}
	return res
}

func GetAllValuesFromColumn(sheetIndex int, column int) []string{
	ret := []string{}
	sheet, err := SpreadSheet.SheetByIndex(uint(sheetIndex))
	checkError(err)
	for _, row := range sheet.Rows {
		for _, cell := range row {
			if cell.Column == uint(column){
				ret = append(ret, cell.Value)
			}
		}
	}
	return ret
}

func UpdateCell(sheetIndex int, row int, column int, value string) (err error){
	sheet, err := SpreadSheet.SheetByIndex(uint(sheetIndex))
	checkError(err)
	sheet.Update(row, column, value)
	err = sheet.Synchronize()
	checkError(err)
	if err != nil {
		return err
	}else{
		return nil
	}
}

func AddContent(sheetIndex int, content map[string]string )(err error){
	sheet, err := SpreadSheet.SheetByIndex(uint(sheetIndex))
	checkError(err)
	cap := len(sheet.Rows)
	for i, cell := range sheet.Rows[0] {
		for key, value := range content {
			if cell.Value == key {
				err = UpdateCell(sheetIndex, cap, i, value)
				checkError(err)
			}
		}
	}
	err = sheet.Synchronize()
	checkError(err)
	if err != nil {
		return err
	}else{
		return nil
	}
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}