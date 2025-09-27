package lucide

import (
	html "github.com/plainkit/html"
)

// BrushCleaning creates a Brush Cleaning Lucide icon.
func BrushCleaning(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-brush-cleaning", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m16 22-1-4"))),
		html.Child(html.SvgPath(html.AD("M19 13.99a1 1 0 0 0 1-1V12a2 2 0 0 0-2-2h-3a1 1 0 0 1-1-1V4a2 2 0 0 0-4 0v5a1 1 0 0 1-1 1H6a2 2 0 0 0-2 2v.99a1 1 0 0 0 1 1"))),
		html.Child(html.SvgPath(html.AD("M5 14h14l1.973 6.767A1 1 0 0 1 20 22H4a1 1 0 0 1-.973-1.233z"))),
		html.Child(html.SvgPath(html.AD("m8 22 1-4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
