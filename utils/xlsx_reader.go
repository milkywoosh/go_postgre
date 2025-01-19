package utils

import (
	"fmt"
	"io"
	"strconv"

	excelize "github.com/xuri/excelize/v2"
)

// "xldata.csv"
func ReadExcel(filename string, sheet_number int) ([][]string, error) {

	var file *excelize.File = &excelize.File{}
	var err error
	opts := excelize.Options{
		Password: "",
	}

	if filename == "" {
		return [][]string{}, fmt.Errorf("filename is empty")
	}

	file, err = excelize.OpenFile(filename, opts)
	if err != nil {
		return [][]string{}, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheet_name := file.GetSheetName(sheet_number)

	datas, err := file.GetRows(sheet_name)
	// omit head
	length := len(datas)
	headless_data := datas[1:length]
	if err != nil {
		return [][]string{}, err
	}

	return headless_data, nil
}

// [MUST BE DONE] read stream from UI(sending xlsx file)
func OpenReader(file_content io.Reader, opts excelize.Options) ([][]string, error) {
	data, err := excelize.OpenReader(file_content, opts)
	if err != nil {
		return nil, err
	}
	sheet_name := data.GetSheetName(0)
	rows, err := data.GetRows(sheet_name)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func MultiplyStock(filename string, sheet_number int) ([][]interface{}, error) {
	data, err := ReadExcel(filename, sheet_number)
	if err != nil {
		return nil, err
	}
	// fmt.Println(data)

	var store [][]interface{} = [][]interface{}{}
	// for i, val := range data {
	// number_val, _ := strconv.Atoi(val[2])
	// data[i][4] = number_val
	// }
	for i := range data {
		val, _ := strconv.Atoi(data[i][1])
		store = append(store, []interface{}{data[i][0], val * 100, data[i][2]})
		// store[i][0] = data[i][0]
		// num, _ := strconv.Atoi(data[i][1])
		// store[i][1] = num * 100
		// store[i][2] = data[i][2]
	}
	return store, nil

}

func CheckSheetNames(filename string, sheet_number int) (string, error) {
	// start from 0 index
	var file *excelize.File = &excelize.File{}
	var err error
	opts := excelize.Options{
		Password: "",
	}

	if filename == "" {
		return "", fmt.Errorf("filename is empty")
	}

	file, err = excelize.OpenFile(filename, opts)
	if err != nil {
		return "", err
	}

	name := file.GetSheetName(sheet_number)
	return name, nil
}
