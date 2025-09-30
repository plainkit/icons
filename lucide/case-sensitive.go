package lucide

import (
	html "github.com/plainkit/html"
)

// CaseSensitive creates a Case Sensitive Lucide icon.
func CaseSensitive(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-case-sensitive", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m2 16 4.039-9.69a.5.5 0 0 1 .923 0L11 16")),
		html.SvgPath(html.AD("M22 9v7")),
		html.SvgPath(html.AD("M3.304 13h6.392")),
		html.SvgCircle(html.ACx("18.5"), html.ACy("12.5"), html.AR("3.5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
