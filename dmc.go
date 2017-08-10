package dmc

import (
	"strconv"
	"strings"
	"time"

	"github.com/get-ion/ion/core/errors"
)

// yy-mm-dd
// mm-dd
func parseDate(s string) (time.Time, error) {
	ts := strings.Split(s, "-")

	if len(ts) == 3 {
		year, err := strconv.Atoi(s[0])
		if err != nil {
			return time.Time{}, err
		}
		month, err := strconv.Atoi(s[1])
		if err != nil {
			return time.Time{}, err
		}
		if 12 < month {
			return time.Time{}, errors.New("parse error")
		}
		day, err := strconv.Atoi(s[2])
		if err != nil {
			return time.Time{}, err
		}
		if 32 < day {
			return time.Time{}, errors.New("parse error")
		}
		return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local), nil

	} else if len(ts) == 2 {
		month, err := strconv.Atoi(s[0])
		if err != nil {
			return time.Time{}, err
		}
		if 12 < month {
			return time.Time{}, errors.New("parse error")
		}
		day, err := strconv.Atoi(s[1])
		if err != nil {
			return time.Time{}, err
		}
		if 32 < day {
			return time.Time{}, errors.New("parse error")
		}
		return time.Date(time.Now().Year(), time.Month(month), day, 0, 0, 0, 0, time.Local), nil
	}

	return time.Time{}, errors.New("parse error")
}

// hh:mm:ss
// mm:ss
// ss
func parseTime(str string) (time.Time, error) {
	h := 0
	m := 0
	s := 0
	var err error
	ts := strings.Split(str, ":")
	if len(ts) == 3 {
		h, err = strconv.Atoi(ts[0])
		if err != nil {
			return time.Time{}, errors.New("parse error")
		}
		m, err = strconv.Atoi(ts[1])
		if err != nil {
			return time.Time{}, errors.New("parse error")
		}
		s, err = strconv.Atoi(ts[2])
		if err != nil {
			return time.Time{}, errors.New("parse error")
		}
	} else if len(ts) == 2 {
		m, err = strconv.Atoi(ts[0])
		if err != nil {
			return time.Time{}, errors.New("parse error")
		}
		s, err = strconv.Atoi(ts[1])
		if err != nil {
			return time.Time{}, errors.New("parse error")
		}
	} else if len(ts) == 1 {
		s, err = strconv.Atoi(ts[0])
		if err != nil {
			return time.Time{}, errors.New("parse error")
		}
	} else {
		return time.Time{}, errors.New("parse error")
	}

	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), h, m, s, 0, time.Local), nil
}


func DimensionalTransfer(s string) (time.Time, error) {
	if strings.Contains(s, "-") || strings.Contains(s, "/") {
		ts := strings.Split(s, "/")
		if len(ts) == 2 {
			d, err := parseDate(ts[0])
			if err != nil {
				return time.Time{}, err
			}
			t, err := parseTime(ts[1])
			if err != nil {
				return time.Time{}, err
			}
			return time.Date(d.Year(), d.Month(), d.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.Local), nil
		} else if len(ts) == 1 {
			return parseDate(ts[0])
		}
		return time.Time{}, errors.New("parse error")
	} else if strings.Contains(s, ":") {
		return parseTime(s)
	} else if strings.Contains(s, "h") || strings.Contains(s, "m") || strings.Contains(s, "s") {
		d, err := time.ParseDuration(s)
		if err != nil {
			return time.Time{}, err
		}

		return time.Now().Add(-1*d), nil
	}
	return time.Time{}, errors.New("parse error")
}
