package lucide

import (
	html "github.com/plainkit/html"
)

// Earth creates a Earth Lucide icon.
func Earth(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-earth", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M21.54 15H17a2 2 0 0 0-2 2v4.54"))),
		html.Child(html.SvgPath(html.AD("M7 3.34V5a3 3 0 0 0 3 3a2 2 0 0 1 2 2c0 1.1.9 2 2 2a2 2 0 0 0 2-2c0-1.1.9-2 2-2h3.17"))),
		html.Child(html.SvgPath(html.AD("M11 21.95V18a2 2 0 0 0-2-2a2 2 0 0 1-2-2v-1a2 2 0 0 0-2-2H2.05"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
