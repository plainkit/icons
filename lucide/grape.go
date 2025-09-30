package lucide

import (
	html "github.com/plainkit/html"
)

// Grape creates a Grape Lucide icon.
func Grape(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-grape", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M22 5V2l-5.89 5.89")),
		html.SvgCircle(html.ACx("16.6"), html.ACy("15.89"), html.AR("3")),
		html.SvgCircle(html.ACx("8.11"), html.ACy("7.4"), html.AR("3")),
		html.SvgCircle(html.ACx("12.35"), html.ACy("11.65"), html.AR("3")),
		html.SvgCircle(html.ACx("13.91"), html.ACy("5.85"), html.AR("3")),
		html.SvgCircle(html.ACx("18.15"), html.ACy("10.09"), html.AR("3")),
		html.SvgCircle(html.ACx("6.56"), html.ACy("13.2"), html.AR("3")),
		html.SvgCircle(html.ACx("10.8"), html.ACy("17.44"), html.AR("3")),
		html.SvgCircle(html.ACx("5"), html.ACy("19"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
