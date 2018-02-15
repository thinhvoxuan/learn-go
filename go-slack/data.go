package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func file2line(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return lines
}

func line2UserLogwork(line string) (user UserLogWork) {
	infors := strings.Split(line, "|")
	if len(infors) < 13 {
		log.Fatalln("Record error: %s", line)
	}
	user.UserAlias = infors[0]
	user.TotalLogwork, _ = strconv.ParseFloat(infors[10], 64)
	user.Commitment, _ = strconv.ParseFloat(infors[9], 64)
	days := []string{"Monday", "Tuesday", "Wednesday", "Thusday", "Friday", "Saturday", "Sunday"}
	for k, d := range days {
		logHour, _ := strconv.ParseFloat(infors[1+k], 64)
		logDay := LogInDay{DateInWeek: d, Hour: logHour}
		user.Logs = append(user.Logs, logDay)
	}
	user.FromDate = infors[11]
	user.ToDate = infors[12]
	return
}

// ReadDataFromFile User logwork with alias
func ReadDataFromFile(filePath string) (users []UserLogWork) {
	lines := file2line(filePath)
	for _, line := range lines {
		user := line2UserLogwork(line)
		users = append(users, user)
	}
	return
}
