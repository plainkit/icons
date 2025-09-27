package lucide

import (
	html "github.com/plainkit/html"
)

// Plus creates a Plus Lucide icon.
func Plus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-plus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M5 12h14"))),
		html.Child(html.SvgPath(html.AD("M12 5v14"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
