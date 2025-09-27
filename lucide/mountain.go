package lucide

import (
	html "github.com/plainkit/html"
)

// Mountain creates a Mountain Lucide icon.
func Mountain(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mountain", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m8 3 4 8 5-5 5 15H2L8 3z"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
