package pp

import "fmt"

type String struct {
	s     string
	color *Color
}

// Pure 返回不包括颜色的底层字符串
func (s String) Pure() string {
	return s.s
}

// Get 返回字符串
func (s String) Get(noColor bool) string {
	return s.color.Sprint(noColor, s.s)
}

func (s String) GetForStdout() string {
	return s.Get(Stdout.NoColor)
}

func (s String) GetForStderr() string {
	return s.Get(Stdout.NoColor)
}

// Wrap 返回添加了颜色头尾的底层字符串
func (s String) Wrap() string {
	return s.color.Sprint(false, s.s)
}

// Color 返回底层颜色数据
func (s String) Color() *Color {
	return s.color
}

func (s *String) ChangeColor(newColor *Color) {
	s.color = newColor
}

func colorString(color *Color, format string, a ...interface{}) String {
	if len(a) == 0 {
		return String{format, color}
	} else {
		return String{fmt.Sprintf(format, a...), color}
	}
}
