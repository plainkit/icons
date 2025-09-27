package lucide

import (
	html "github.com/plainkit/html"
)

// Menu creates a Menu Lucide icon.
func Menu(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-menu", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M4 5h16"))),
		html.Child(html.SvgPath(html.AD("M4 12h16"))),
		html.Child(html.SvgPath(html.AD("M4 19h16"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
