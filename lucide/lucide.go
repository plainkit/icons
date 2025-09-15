package lucide

import x "github.com/plainkit/html"

// iconSizeArg represents a size argument that sets both width and height to the same value.
type iconSizeArg struct {
	x.Global
	size string
}

// Size creates an icon size argument that sets both width and height to the specified value.
// This is the preferred way to size Lucide icons uniformly.
//
// Example:
//
//	lucide.Heart(lucide.Size("16"))
//	lucide.Star(lucide.Size("24"), x.Class("text-yellow-500"))
func Size(size string) x.SvgArg {
	return iconSizeArg{
		Global: x.Class(""),
		size:   size,
	}
}

// buildLucideArgs constructs the final SVG arguments for a Lucide icon.
// It processes user arguments, applies defaults, and ensures proper ordering.
func buildLucideArgs(className string, userArgs []x.SvgArg) []x.SvgArg {
	hasWidth := false
	hasHeight := false
	hasStrokeWidth := false
	hasClass := false
	var iconSize string

	processedArgs := make([]x.SvgArg, 0, len(userArgs))

	for _, arg := range userArgs {
		switch v := arg.(type) {
		case iconSizeArg:
			iconSize = v.size
			hasWidth = true
			hasHeight = true
		case x.SvgWidthOpt:
			hasWidth = true
			processedArgs = append(processedArgs, arg)
		case x.SvgHeightOpt:
			hasHeight = true
			processedArgs = append(processedArgs, arg)
		case x.StrokeWidthOpt:
			hasStrokeWidth = true
			processedArgs = append(processedArgs, arg)
		case x.Global:
			hasClass = true
			processedArgs = append(processedArgs, arg)
		default:
			processedArgs = append(processedArgs, arg)
		}
	}

	result := make([]x.SvgArg, 0, len(processedArgs)+10)

	result = append(result, x.Xmlns("http://www.w3.org/2000/svg"))

	if iconSize != "" {
		result = append(result, x.SvgWidth(iconSize))
		result = append(result, x.SvgHeight(iconSize))
	} else {
		if !hasWidth {
			result = append(result, x.SvgWidth("24"))
		}
		if !hasHeight {
			result = append(result, x.SvgHeight("24"))
		}
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

	result = append(result, processedArgs...)

	return result
}
