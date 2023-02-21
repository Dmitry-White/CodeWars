package main

import (
	"fmt"
	"strings"
)

const (
	SECOND = 1
	MINUTE = 60 * SECOND
	HOUR   = 60 * MINUTE
	DAY    = 24 * HOUR
	YEAR   = 365 * DAY
)

type Duration struct {
	seconds  int64
	minutes  int64
	hours    int64
	days     int64
	years    int64
	raw      []string
	readable string
}

type LanguageForms struct {
	single string
	plural string
}

type DateParts struct {
	seconds string
	minutes string
	hours   string
	days    string
	years   string
}

var DurationKeys = DateParts{
	seconds: "seconds",
	minutes: "minutes",
	hours:   "hours",
	days:    "days",
	years:   "years",
}

var FormatMap = map[string]LanguageForms{
	"seconds": {
		single: "second",
		plural: "seconds",
	},
	"minutes": {
		single: "minute",
		plural: "minutes",
	},
	"hours": {
		single: "hour",
		plural: "hours",
	},
	"days": {
		single: "day",
		plural: "days",
	},
	"years": {
		single: "year",
		plural: "years",
	},
}

func (f *Duration) computeMinutes() Duration {
	f.minutes = f.seconds / MINUTE
	f.seconds = f.seconds - f.minutes*MINUTE
	return *f
}

func (f *Duration) computeHours() Duration {
	f.hours = f.seconds / HOUR
	f.seconds = f.seconds - f.hours*HOUR
	return *f
}

func (f *Duration) computeDays() Duration {
	f.days = f.seconds / DAY
	f.seconds = f.seconds - f.days*DAY
	return *f
}

func (f *Duration) computeYears() Duration {
	f.years = f.seconds / YEAR
	f.seconds = f.seconds - f.years*YEAR
	return *f
}

func (f *Duration) computeDate() Duration {
	f.computeYears()
	f.computeDays()
	f.computeHours()
	f.computeMinutes()
	return *f
}

func (f *Duration) computeRaw(time int64, part string) {
	if time != 0 {
		switch time {
		case 1:
			partStr := fmt.Sprintf("1 %v", FormatMap[part].single)
			f.raw = append(f.raw, partStr)
		default:
			partStr := fmt.Sprintf("%v %v", time, FormatMap[part].plural)
			f.raw = append(f.raw, partStr)
		}
	}
}

func (f *Duration) prettify() Duration {
	f.computeRaw(f.years, DurationKeys.years)
	f.computeRaw(f.days, DurationKeys.days)
	f.computeRaw(f.hours, DurationKeys.hours)
	f.computeRaw(f.minutes, DurationKeys.minutes)
	f.computeRaw(f.seconds, DurationKeys.seconds)

	lastPart := f.raw[len(f.raw)-1:][0]
	mainPart := strings.Join(f.raw[:len(f.raw)-1], ", ")

	if len(f.raw) == 1 {
		f.readable = lastPart
	} else {
		f.readable = fmt.Sprintf("%v and %v", mainPart, lastPart)
	}

	return *f
}

func FormatDuration(seconds int64) string {
	humanFormat := Duration{seconds: seconds}

	humanFormat.computeDate()

	humanFormat.prettify()

	return humanFormat.readable
}
