package lucide

import (
	html "github.com/plainkit/html"
)

// Bolt creates a Bolt Lucide icon.
func Bolt(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bolt", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
