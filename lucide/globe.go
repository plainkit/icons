package lucide

import (
	html "github.com/plainkit/html"
)

// Globe creates a Globe Lucide icon.
func Globe(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-globe", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("M12 2a14.5 14.5 0 0 0 0 20 14.5 14.5 0 0 0 0-20"))),
		html.Child(html.SvgPath(html.AD("M2 12h20"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
