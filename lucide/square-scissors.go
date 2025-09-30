package lucide

import (
	html "github.com/plainkit/html"
)

// SquareScissors creates a Square Scissors Lucide icon.
func SquareScissors(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-square-scissors", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("20"), html.AHeight("20"), html.AX("2"), html.AY("2"), html.ARx("2")),
		html.SvgCircle(html.ACx("8"), html.ACy("8"), html.AR("2")),
		html.SvgPath(html.AD("M9.414 9.414 12 12")),
		html.SvgPath(html.AD("M14.8 14.8 18 18")),
		html.SvgCircle(html.ACx("8"), html.ACy("16"), html.AR("2")),
		html.SvgPath(html.AD("m18 6-8.586 8.586")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
