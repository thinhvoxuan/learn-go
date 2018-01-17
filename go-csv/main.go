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
		log.Println(dateString)
		return dateString
	}
	return dateLog.Format("1/2/2006")
}

func convertDateColIndex(rec []string, index int) {
	rec[index] = convertDate(rec[index])
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
		convertDateColIndex(record, 6)
		convertDateColIndex(record, 8)
		// record[6] = convertDate(record[6])
		// record[8] = convertDate(record[8])
		errWrite := w.Write(record)
		if errWrite != nil {
			log.Fatal(errWrite)
			// panic(errWrite)
		}
	}
	w.Flush()
}
