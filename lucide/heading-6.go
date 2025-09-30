package lucide

import (
	html "github.com/plainkit/html"
)

// Heading6 creates a Heading 6 Lucide icon.
func Heading6(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-heading-6", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 12h8")),
		html.SvgPath(html.AD("M4 18V6")),
		html.SvgPath(html.AD("M12 18V6")),
		html.SvgCircle(html.ACx("19"), html.ACy("16"), html.AR("2")),
		html.SvgPath(html.AD("M20 10c-2 2-3 3.5-3 6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
