package lucide

import (
	html "github.com/plainkit/html"
)

// Sailboat creates a Sailboat Lucide icon.
func Sailboat(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sailboat", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 2v15")),
		html.SvgPath(html.AD("M7 22a4 4 0 0 1-4-4 1 1 0 0 1 1-1h16a1 1 0 0 1 1 1 4 4 0 0 1-4 4z")),
		html.SvgPath(html.AD("M9.159 2.46a1 1 0 0 1 1.521-.193l9.977 8.98A1 1 0 0 1 20 13H4a1 1 0 0 1-.824-1.567z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
