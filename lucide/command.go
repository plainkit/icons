package lucide

import (
	html "github.com/plainkit/html"
)

// Command creates a Command Lucide icon.
func Command(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-command", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M15 6v12a3 3 0 1 0 3-3H6a3 3 0 1 0 3 3V6a3 3 0 1 0-3 3h12a3 3 0 1 0-3-3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
