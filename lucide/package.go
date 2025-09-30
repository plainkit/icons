package lucide

import (
	html "github.com/plainkit/html"
)

// Package creates a Package Lucide icon.
func Package(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-package", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M11 21.73a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73z")),
		html.SvgPath(html.AD("M12 22V12")),
		html.SvgPolyline(html.APoints("3.29 7 12 12 20.71 7")),
		html.SvgPath(html.AD("m7.5 4.27 9 5.15")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
