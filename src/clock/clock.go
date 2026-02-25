package clock

import (
	"fmt"
	"strings"
	"time"
)

var zone = map[string]string{
	"utc":   "UTC",
	"local": "Local",
	"est":   "America/New_York",
	"cst":   "America/Chicago",
	"mst":   "America/Denver",
	"pst":   "America/Los_Angeles",
	"cet":   "Europe/Paris",
	"ist":   "Asia/Kolkata",
	"jst":   "Asia/Tokyo",
}

func load(name string) (*time.Location, string, error) {
	name = strings.ToLower(strings.TrimSpace(name))
	if name == "" {
		name = "local"
	}
	value, ok := zone[name]
	if !ok {
		return nil, "", fmt.Errorf("unknown zone: %s", name)
	}
	if value == "Local" {
		return time.Local, "local", nil
	}
	loc, err := time.LoadLocation(value)
	if err != nil {
		return nil, "", err
	}
	return loc, name, nil
}

func line(now time.Time, alias string) string {
	name, offset := now.Zone()
	sign := "+"
	if offset < 0 {
		sign = "-"
		offset = -offset
	}
	hour := offset / 3600
	minute := (offset % 3600) / 60
	stamp := now.Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s  %s  %s (%s%02d:%02d)", alias, stamp, name, sign, hour, minute)
}

func Render(name string, clock time.Time) (string, error) {
	loc, alias, err := load(name)
	if err != nil {
		return "", err
	}
	return line(clock.In(loc), alias), nil
}

func Help() string {
	return "zones: utc local est cst mst pst cet ist jst"
}
