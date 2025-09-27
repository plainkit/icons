package lucide

import (
	html "github.com/plainkit/html"
)

// Haze creates a Haze Lucide icon.
func Haze(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-haze", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m5.2 6.2 1.4 1.4"))),
		html.Child(html.SvgPath(html.AD("M2 13h2"))),
		html.Child(html.SvgPath(html.AD("M20 13h2"))),
		html.Child(html.SvgPath(html.AD("m17.4 7.6 1.4-1.4"))),
		html.Child(html.SvgPath(html.AD("M22 17H2"))),
		html.Child(html.SvgPath(html.AD("M22 21H2"))),
		html.Child(html.SvgPath(html.AD("M16 13a4 4 0 0 0-8 0"))),
		html.Child(html.SvgPath(html.AD("M12 5V2.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
