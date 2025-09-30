package lucide

import (
	html "github.com/plainkit/html"
)

// Shapes creates a Shapes Lucide icon.
func Shapes(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-shapes", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M8.3 10a.7.7 0 0 1-.626-1.079L11.4 3a.7.7 0 0 1 1.198-.043L16.3 8.9a.7.7 0 0 1-.572 1.1Z")),
		html.SvgRect(html.AWidth("7"), html.AHeight("7"), html.AX("3"), html.AY("14"), html.ARx("1")),
		html.SvgCircle(html.ACx("17.5"), html.ACy("17.5"), html.AR("3.5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
