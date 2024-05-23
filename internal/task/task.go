package task

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

var (
	ErrFormat   = errors.New("task format error")
	ErrNotFound = errors.New("not found")
)

func NextDate(now time.Time, startDate string, repeat string) (string, error) {

	repeat = strings.ToLower(strings.TrimSpace(repeat))
	startDt, err := time.Parse("20060102", startDate)
	if err != nil {
		return "", fmt.Errorf("%w: unexpected date value", ErrFormat)
	}

	var nextDate time.Time

	switch {
	case repeat == "y":
		nextDate = startDt
		nextDate = nextDate.AddDate(1, 0, 0)
		for nextDate.Before(now) {
			nextDate = nextDate.AddDate(1, 0, 0)
		}

	case strings.HasPrefix(repeat, "d"):
		d, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(repeat, "d ")))
		if err != nil || d < 1 || d > 400 {
			return "", fmt.Errorf("%w: invalid days %w", ErrFormat, err)
		}
		nextDate = startDt
		nextDate = nextDate.AddDate(0, 0, d)
		for nextDate.Before(now) {
			nextDate = nextDate.AddDate(0, 0, d)
		}
	default:
		return "", fmt.Errorf("%w: unexpected repeat value", ErrFormat)

	}
	return nextDate.Format("20060102"), nil
}

func validateData(tsk *Task) error {

	if tsk.Title == "" || tsk.Title == " " {
		return fmt.Errorf("%w: title is empty", ErrFormat)
	}

	if tsk.Date == "" {
		tsk.Date = time.Now().Format("20060102")
	}

	_, err := time.Parse("20060102", tsk.Date)
	if err != nil {
		return fmt.Errorf("%w: wrong data format", ErrFormat)
	}

	return nil

}
