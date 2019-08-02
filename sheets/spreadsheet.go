package sheets

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"io/ioutil"
	"strconv"
)

var SpreadSheet = spreadsheet.Spreadsheet{}

func connect(client_secret string, sheetId string) spreadsheet.Spreadsheet {
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

func Initialize(client_secret string, sheetId string) {
	SpreadSheet = connect(client_secret, sheetId)
}

func AutoSumAllValuesOnColumn(sheetTitle string, column int) float64 {
	res := 0.0
	sheet, err := SpreadSheet.SheetByTitle(sheetTitle)
	checkError(err)
	for _, row := range sheet.Rows {
		for _, cell := range row {
			if cell.Column == uint(column) {
				f, err := strconv.ParseFloat(cell.Value, 64)
				checkError(err)
				res += f
			}
		}
	}
	return res
}

func GetAllValuesFromColumn(sheetTitle string, column int) []string {
	ret := []string{}
	sheet, err := SpreadSheet.SheetByTitle(sheetTitle)
	checkError(err)
	for _, row := range sheet.Rows {
		for _, cell := range row {
			if cell.Column == uint(column) {
				ret = append(ret, cell.Value)
			}
		}
	}
	return ret
}

func UpdateCell(sheetTitle string, row int, column int, value string) (err error) {
	sheet, err := SpreadSheet.SheetByTitle(sheetTitle)
	checkError(err)
	sheet.Update(row, column, value)
	err = sheet.Synchronize()
	checkError(err)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func UpdateContent(sheetTitle string, indexColumn string, index string, content map[string]interface{}) (err error) {
	sheet, err := SpreadSheet.SheetByTitle(sheetTitle)
	checkError(err)
	for _, columns := range sheet.Rows[0] {
		if columns.Value == indexColumn {
			for j, columnCell := range sheet.Rows {
				if columnCell[columns.Column].Value == index {
					for key, value := range content {
						for _, tuple := range sheet.Rows[j] {
							if sheet.Rows[0][tuple.Column].Value == key && value != "" {
								UpdateCell(sheetTitle, int(tuple.Row), int(tuple.Column), fmt.Sprintf("%v", value))
								break
							}
						}
					}
				}
			}
		}
	}
	err = sheet.Synchronize()
	checkError(err)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func AddContent(sheetTitle string, content map[string]interface{}) (err error) {
	sheet, err := SpreadSheet.SheetByTitle(sheetTitle)
	checkError(err)
	cap := len(sheet.Rows)
	for i, cell := range sheet.Rows[0] {
		for key, value := range content {
			if cell.Value == key {
				err = UpdateCell(sheetTitle, cap, i, fmt.Sprintf("%v", value))
				checkError(err)
			}
		}
	}
	err = sheet.Synchronize()
	checkError(err)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetRowByValueInColumn(sheetTitle string, value string, column string) (data map[string]interface{}, err error) {
	sheet, err := SpreadSheet.SheetByTitle(sheetTitle)
	checkError(err)
	if err != nil {
		return nil, err
	} else {
		data = map[string]interface{}{}
		for _, columns := range sheet.Rows[0] {
			if columns.Value == column {
				for j, columnCell := range sheet.Rows {
					if columnCell[columns.Column].Value == value {
						for _, tuple := range sheet.Rows[j] {
							data[sheet.Rows[0][tuple.Column].Value] = tuple.Value
						}
					}
				}
			}
		}
		return data, nil
	}
}

func GetRow(sheetTitle string, row int) (data map[string]interface{}, err error) {
	sheet, err := SpreadSheet.SheetByTitle(sheetTitle)
	checkError(err)
	if err != nil {
		return nil, err
	} else {
		data = map[string]interface{}{}
		for i := range sheet.Rows[0] {
			data[sheet.Rows[0][i].Value] = sheet.Rows[row][i].Value
		}
		return data, nil
	}
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
