package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"
)

func convertDate(dateString string) string {
	dateLog, errorParseDate := time.Parse("02/01/06", dateString)
	if errorParseDate != nil {
		// log.Fatal(errorParseDate)
		return dateString
	}
	return dateLog.Format("01/02/2006")
}

func main() {
	readerFile, errReadFile := os.OpenFile("data/input.csv", os.O_RDONLY, os.ModePerm)
	if errReadFile != nil {
		println(errReadFile)
		panic(errReadFile)
	}
	defer readerFile.Close()

	writerFile, errWriteFile := os.Create("data/Worklog-formated.csv")

	if errWriteFile != nil {
		panic(errWriteFile)
	}
	defer writerFile.Close()

	r := csv.NewReader(bufio.NewReader(readerFile))
	w := csv.NewWriter(writerFile)
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		record[5] = convertDate(record[5])
		record[7] = convertDate(record[7])
		errWrite := w.Write(record)
		if errWrite != nil {
			log.Fatal(errWrite)
			// panic(errWrite)
		}
	}
	w.Flush()
}
