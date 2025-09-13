package lucide

import x "github.com/bloxui/blox"

// Size sets both width and height to the same value for square icons
func Size(size string) []x.SvgArg {
	return []x.SvgArg{
		x.SvgWidth(size),
		x.SvgHeight(size),
	}
}

// StrokeWidth sets the stroke width for the icon
func StrokeWidth(width string) x.SvgArg {
	return x.StrokeWidth(width)
}

// buildLucideArgs creates the final argument list for a Lucide icon
// User arguments override defaults
func buildLucideArgs(className string, userArgs []x.SvgArg) []x.SvgArg {
	// Check what user has provided
	hasWidth := false
	hasHeight := false
	hasStrokeWidth := false
	hasClass := false

	for _, arg := range userArgs {
		switch arg.(type) {
		case x.SvgWidthOpt:
			hasWidth = true
		case x.SvgHeightOpt:
			hasHeight = true
		case x.StrokeWidthOpt:
			hasStrokeWidth = true
		case x.Global:
			// For Global type (which includes Class), we need to check if it's a class
			// Since we can't easily inspect the Global's internal function, we'll assume
			// any Global might be setting class and let it override
			hasClass = true
		}
	}

	// Start with user arguments
	result := make([]x.SvgArg, 0, len(userArgs)+10)
	result = append(result, userArgs...)

	// Add defaults only if not provided by user
	result = append(result, x.Xmlns("http://www.w3.org/2000/svg"))
	if !hasWidth {
		result = append(result, x.SvgWidth("24"))
	}
	if !hasHeight {
		result = append(result, x.SvgHeight("24"))
	}
	result = append(result, x.ViewBox("0 0 24 24"))
	result = append(result, x.Fill("none"))
	result = append(result, x.Stroke("currentColor"))
	if !hasStrokeWidth {
		result = append(result, x.StrokeWidth("2"))
	}
	result = append(result, x.StrokeLinecap("round"))
	result = append(result, x.StrokeLinejoin("round"))
	if !hasClass {
		result = append(result, x.Class(className))
	}

	return result
}
