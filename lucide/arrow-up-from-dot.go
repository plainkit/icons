package lucide

import (
	html "github.com/plainkit/html"
)

// ArrowUpFromDot creates a Arrow Up From Dot Lucide icon.
func ArrowUpFromDot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-arrow-up-from-dot", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m5 9 7-7 7 7"))),
		html.Child(html.SvgPath(html.AD("M12 16V2"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("21"), html.AR("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
