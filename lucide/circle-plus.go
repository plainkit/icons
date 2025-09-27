package lucide

import (
	html "github.com/plainkit/html"
)

// CirclePlus creates a Circle Plus Lucide icon.
func CirclePlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-plus", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgPath(html.AD("M8 12h8"))),
		html.Child(html.SvgPath(html.AD("M12 8v8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
