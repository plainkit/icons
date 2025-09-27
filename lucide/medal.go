package lucide

import (
	html "github.com/plainkit/html"
)

// Medal creates a Medal Lucide icon.
func Medal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-medal", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M7.21 15 2.66 7.14a2 2 0 0 1 .13-2.2L4.4 2.8A2 2 0 0 1 6 2h12a2 2 0 0 1 1.6.8l1.6 2.14a2 2 0 0 1 .14 2.2L16.79 15"))),
		html.Child(html.SvgPath(html.AD("M11 12 5.12 2.2"))),
		html.Child(html.SvgPath(html.AD("m13 12 5.88-9.8"))),
		html.Child(html.SvgPath(html.AD("M8 7h8"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("17"), html.AR("5"))),
		html.Child(html.SvgPath(html.AD("M12 18v-2h-.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
