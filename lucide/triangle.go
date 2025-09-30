package lucide

import (
	html "github.com/plainkit/html"
)

// Triangle creates a Triangle Lucide icon.
func Triangle(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-triangle", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13.73 4a2 2 0 0 0-3.46 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
