package utility

import (
    "os"
    "encoding/csv"
)


func StringArraystoCSV(data [][]string){

    file, _ := os.OpenFile("test.csv", os.O_WRONLY|os.O_CREATE, os.ModePerm)
    w := csv.NewWriter(file)

    //write using UTF-8
    file.WriteString("\xEF\xBB\xBF")
    w.WriteAll(data)
    w.Flush()
    file.Close()
}