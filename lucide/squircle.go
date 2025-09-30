package lucide

import (
	html "github.com/plainkit/html"
)

// Squircle creates a Squircle Lucide icon.
func Squircle(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-squircle", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 3c7.2 0 9 1.8 9 9s-1.8 9-9 9-9-1.8-9-9 1.8-9 9-9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
