package lucide

import (
	html "github.com/plainkit/html"
)

// Transgender creates a Transgender Lucide icon.
func Transgender(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-transgender", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 16v6"))),
		html.Child(html.SvgPath(html.AD("M14 20h-4"))),
		html.Child(html.SvgPath(html.AD("M18 2h4v4"))),
		html.Child(html.SvgPath(html.AD("m2 2 7.17 7.17"))),
		html.Child(html.SvgPath(html.AD("M2 5.355V2h3.357"))),
		html.Child(html.SvgPath(html.AD("m22 2-7.17 7.17"))),
		html.Child(html.SvgPath(html.AD("M8 5 5 8"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
