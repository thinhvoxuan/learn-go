package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// LogInDay struct
type LogInDay struct {
	DateInWeek string
	Hour       float64
}

// UserLogWork information
type UserLogWork struct {
	UserAlias    string
	Commitment   float64
	TotalLogwork float64
	FromDate     string
	ToDate       string
	Logs         []LogInDay
}

func (user UserLogWork) completePercent() float64 {
	return user.TotalLogwork / user.Commitment * 100
}

func (user UserLogWork) summaryMessage() string {
	return fmt.Sprintf("Total: *%.2fh*.\nComparing to your comit: %.2f/%.2f (%0.2f%%).\n", user.TotalLogwork, user.TotalLogwork, user.Commitment, user.completePercent())
}

func (user UserLogWork) haventLogDate() string {
	days := []string{}
	for _, day := range user.Logs {
		if day.DateInWeek != "Saturday" && day.DateInWeek != "Sunday" && day.Hour < 1 {
			days = append(days, day.DateInWeek)
		}
	}

	if len(days) == 0 {
		return ""
	}

	return fmt.Sprintf("Haven't log days: %s.\n", strings.Join(days, ", "))
}

func (user UserLogWork) loggedDayMessage() string {
	days := []string{}
	for _, day := range user.Logs {
		if day.Hour != 0 {
			d := fmt.Sprintf("%s - %0.1f", day.DateInWeek[:3], day.Hour)
			days = append(days, d)
		}
	}

	if len(days) == 0 {
		return "No logged day.\n"
	}

	return fmt.Sprintf("Logged day: %s.\n", strings.Join(days, ", "))
}

func (user UserLogWork) isCompleteLogWork() bool {
	return user.TotalLogwork >= user.Commitment
}

func (user UserLogWork) hiMessage() string {
	return fmt.Sprintf("Hi @%s\nSummary logwork information %s (%s to %s).\n", user.UserAlias, weekNumber(user.FromDate), user.FromDate, user.ToDate)
}

func (user UserLogWork) call4ActionMessage() string {
	if user.isCompleteLogWork() {
		return fmt.Sprintf("Thank for completing logwork *%s* :heart: ", weekNumber(user.FromDate))
	}
	return fmt.Sprintf("Please complete your logwork *%s* before 8:30 AM next Monday", weekNumber(user.FromDate))
}

func weekNumber(dateString string) string {
	dateLog, errorParseDate := time.Parse("1/02/2006", dateString)
	if errorParseDate != nil {
		log.Fatalf("Error date: %s", errorParseDate)
	}
	year, week := dateLog.ISOWeek()
	return fmt.Sprintf("W%.2d-%d", week, year)
}

func (user UserLogWork) messageNotice() (message string) {

	lines := []string{
		user.hiMessage(),
		user.summaryMessage(),
		user.loggedDayMessage(),
		user.haventLogDate(),
		user.call4ActionMessage(),
	}

	message = strings.Join(lines, "")
	return
}
