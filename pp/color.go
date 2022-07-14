package pp

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var (
	// colorsCache is used to reduce the count of created Color objects and
	// allows to reuse already created objects with required Attribute.
	colorsCache   = make(map[Attribute]*Color)
	colorsCacheMu sync.Mutex // protects colorsCache
)

// Color defines a custom color object which is defined by SGR parameters.
type Color struct {
	params   []Attribute
	format   string
	unformat string
}

// Attribute defines a single SGR Code
type Attribute int

const escape = "\x1b"

// Base attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

// Sprint is just like Print, but returns a string instead of printing it.
func (c *Color) Sprint(noColor bool, a ...interface{}) string {
	return c.wrap(fmt.Sprint(a...), noColor)
}

// Sprintln is just like Println, but returns a string instead of printing it.
func (c *Color) Sprintln(noColor bool, a ...interface{}) string {
	return c.wrap(fmt.Sprintln(a...), noColor)
}

// Sprintf is just like Printf, but returns a string instead of printing it.
func (c *Color) Sprintf(noColor bool, format string, a ...interface{}) string {
	return c.wrap(fmt.Sprintf(format, a...), noColor)
}

// wrap wraps the s string with the colors attributes. The string is ready to
// be printed.
func (c *Color) wrap(s string, noColor bool) string {
	if noColor || c == nil {
		return s
	}
	return c.format + s + c.unformat
}

// Equals returns a boolean value indicating whether two colors are equal.
func (c *Color) Equals(c2 *Color) bool {
	if len(c.params) != len(c2.params) {
		return false
	}

	for _, attr := range c.params {
		if !c2.attrExists(attr) {
			return false
		}
	}

	return true
}

func (c *Color) attrExists(a Attribute) bool {
	for _, attr := range c.params {
		if attr == a {
			return true
		}
	}

	return false
}

// GetColor 利用 attribute 获取 Color
func GetColor(attributes ...Attribute) *Color {
	switch len(attributes) {
	case 0:
		return nil
	case 1:
		return getColorCached(attributes[0])
	default:
		return newColor(attributes...)
	}
}

// newColor returns a newly created color object.
func newColor(value ...Attribute) *Color {
	c := &Color{
		params:   value,
		format:   fmt.Sprintf("%s[%sm", escape, sequence(value)),
		unformat: fmt.Sprintf("%s[%dm", escape, Reset),
	}

	return c
}

// sequence returns a formatted SGR sequence to be plugged into a "\x1b[...m"
// an example output might be: "1;36" -> bold cyan
func sequence(params []Attribute) string {
	format := make([]string, len(params))
	for i, v := range params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

func getColorCached(p Attribute) *Color {
	colorsCacheMu.Lock()
	defer colorsCacheMu.Unlock()

	c, ok := colorsCache[p]
	if !ok {
		c = newColor(p)
		colorsCache[p] = c
	}

	return c
}
