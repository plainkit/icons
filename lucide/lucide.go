package lucide

import (
	html "github.com/plainkit/html"
)

// Size creates an icon size argument that sets both width and height to the specified value.
// This is the preferred way to size Lucide icons uniformly.
//
// Example:
//
//	lucide.Heart(lucide.Size("16"))
//	lucide.Star(lucide.Size("24"), html.AClass("text-yellow-500"))
func Size(size string) []html.SvgArg {
	return []html.SvgArg{
		html.AWidth(size),
		html.AHeight(size),
	}
}

// withLucideDefaults constructs SVG arguments with Lucide icon defaults.
// It prepends standard Lucide styling and attributes to user-provided arguments.
func withLucideDefaults(className string, userArgs []html.SvgArg) []html.SvgArg {
	defaults := []html.SvgArg{
		html.AWidth("24"),
		html.AHeight("24"),
		html.AViewBox("0 0 24 24"),
		html.AFill("none"),
		html.AStroke("currentColor"),
		html.AStrokeWidth("2"),
		html.AStrokeLinecap("round"),
		html.AStrokeLinejoin("round"),
		html.AXmlns("http://www.w3.org/2000/svg"),
		html.AClass(className),
	}

	return append(defaults, userArgs...)
}
