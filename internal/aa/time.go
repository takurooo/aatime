package aa

import (
	"embed"
	"fmt"
	"strings"
	"time"
)

//go:embed static/*
var asciiArts embed.FS

type Font string

const (
	FontInvita   Font = "invita"
	FontTrain    Font = "train"
	FontBlubhead Font = "bulbhead"
)

func FontFromString(s string) Font {
	switch s {
	case "invita":
		return FontInvita
	case "train":
		return FontTrain
	case "bulbhead":
		return FontBlubhead
	default:
		return FontInvita
	}
}

type asciiArtTime struct {
	utc       bool
	numbers   map[string]string
	separator string
}

func (a *asciiArtTime) now() string {
	var now time.Time
	if a.utc {
		now = time.Now().UTC()
	} else {
		now = time.Now()
	}

	year := fmt.Sprintf("%04d", now.Year())
	month := fmt.Sprintf("%02d", now.Month())
	day := fmt.Sprintf("%02d", now.Day())
	hour := fmt.Sprintf("%02d", now.Hour())
	minute := fmt.Sprintf("%02d", now.Minute())
	second := fmt.Sprintf("%02d", now.Second())

	sep := strings.Split(a.separator, "\n")
	y1 := strings.Split(a.numbers[year[0:1]], "\n")
	y2 := strings.Split(a.numbers[year[1:2]], "\n")
	y3 := strings.Split(a.numbers[year[2:3]], "\n")
	y4 := strings.Split(a.numbers[year[3:4]], "\n")
	m1 := strings.Split(a.numbers[month[0:1]], "\n")
	m2 := strings.Split(a.numbers[month[1:2]], "\n")
	d1 := strings.Split(a.numbers[day[0:1]], "\n")
	d2 := strings.Split(a.numbers[day[1:2]], "\n")
	h1 := strings.Split(a.numbers[hour[0:1]], "\n")
	h2 := strings.Split(a.numbers[hour[1:2]], "\n")
	min1 := strings.Split(a.numbers[minute[0:1]], "\n")
	min2 := strings.Split(a.numbers[minute[1:2]], "\n")
	s1 := strings.Split(a.numbers[second[0:1]], "\n")
	s2 := strings.Split(a.numbers[second[1:2]], "\n")

	out := ""
	for i := 0; i < len(y1); i++ {
		out += y1[i] + y2[i] + y3[i] + y4[i] + sep[i] +
			m1[i] + m2[i] + sep[i] +
			d1[i] + d2[i] + sep[i] +
			h1[i] + h2[i] + sep[i] +
			min1[i] + min2[i] + sep[i] +
			s1[i] + s2[i] + "\n"
	}
	return out
}

func newASCIIArtTime(font Font, utc bool) *asciiArtTime {
	dir := "static/" + font
	n := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	asciiArtNumbers := make(map[string]string)
	for _, nn := range n {
		b, _ := asciiArts.ReadFile(fmt.Sprintf("%s/%s.txt", dir, nn))
		asciiArtNumbers[nn] = string(b)
	}
	b, _ := asciiArts.ReadFile(fmt.Sprintf("%s/sep.txt", dir))
	asciiArtSep := string(b)
	return &asciiArtTime{
		utc:       utc,
		numbers:   asciiArtNumbers,
		separator: asciiArtSep,
	}
}

func Now(font Font, utc bool) string {
	aaTime := newASCIIArtTime(font, utc)
	return aaTime.now()
}
