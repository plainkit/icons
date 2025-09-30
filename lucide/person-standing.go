package lucide

import (
	html "github.com/plainkit/html"
)

// PersonStanding creates a Person Standing Lucide icon.
func PersonStanding(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-person-standing", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("5"), html.AR("1")),
		html.SvgPath(html.AD("m9 20 3-6 3 6")),
		html.SvgPath(html.AD("m6 8 6 2 6-2")),
		html.SvgPath(html.AD("M12 10v4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
