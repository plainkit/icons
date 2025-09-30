package lucide

import (
	html "github.com/plainkit/html"
)

// CaseLower creates a Case Lower Lucide icon.
func CaseLower(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-case-lower", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 9v7")),
		html.SvgPath(html.AD("M14 6v10")),
		html.SvgCircle(html.ACx("17.5"), html.ACy("12.5"), html.AR("3.5")),
		html.SvgCircle(html.ACx("6.5"), html.ACy("12.5"), html.AR("3.5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
