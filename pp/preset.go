package pp

var (
	RED    = getColorCached(FgRed)
	GREEN  = getColorCached(FgGreen)
	BLUE   = getColorCached(FgBlue)
	CYAN   = getColorCached(FgCyan)
	YELLOW = getColorCached(FgYellow)

	BOLD      = getColorCached(Bold)
	UNDERLINE = getColorCached(Underline)
	ITALIC    = getColorCached(Italic)

	BOLDUNDERLINE   = GetColor(Bold, Underline)
	REDUNDERLINE    = GetColor(FgRed, Underline)
	GREENUNDERLINE  = GetColor(FgGreen, Underline)
	BLUEUNDERLINE   = GetColor(FgBlue, Underline)
	CYANUNDERLINE   = GetColor(FgCyan, Underline)
	YELLOWUNDERLINE = GetColor(FgYellow, Underline)
)
