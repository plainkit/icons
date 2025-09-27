package lucide

import (
	html "github.com/plainkit/html"
)

// Box creates a Box Lucide icon.
func Box(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-box", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21 8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16Z"))),
		html.Child(html.SvgPath(html.AD("m3.3 7 8.7 5 8.7-5"))),
		html.Child(html.SvgPath(html.AD("M12 22V12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
