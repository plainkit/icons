package lucide

import (
	html "github.com/plainkit/html"
)

// Sunrise creates a Sunrise Lucide icon.
func Sunrise(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sunrise", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 2v8")),
		html.SvgPath(html.AD("m4.93 10.93 1.41 1.41")),
		html.SvgPath(html.AD("M2 18h2")),
		html.SvgPath(html.AD("M20 18h2")),
		html.SvgPath(html.AD("m19.07 10.93-1.41 1.41")),
		html.SvgPath(html.AD("M22 22H2")),
		html.SvgPath(html.AD("m8 6 4-4 4 4")),
		html.SvgPath(html.AD("M16 18a4 4 0 0 0-8 0")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
