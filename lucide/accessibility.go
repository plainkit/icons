package lucide

import (
	html "github.com/plainkit/html"
)

// Accessibility creates a Accessibility Lucide icon.
func Accessibility(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-accessibility", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("16"), html.ACy("4"), html.AR("1")),
		html.SvgPath(html.AD("m18 19 1-7-6 1")),
		html.SvgPath(html.AD("m5 8 3-3 5.5 3-2.36 3.5")),
		html.SvgPath(html.AD("M4.24 14.5a5 5 0 0 0 6.88 6")),
		html.SvgPath(html.AD("M13.76 17.5a5 5 0 0 0-6.88-6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
