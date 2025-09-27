package lucide

import (
	html "github.com/plainkit/html"
)

// Bike creates a Bike Lucide icon.
func Bike(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-bike", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("18.5"), html.ACy("17.5"), html.AR("3.5"))),
		html.Child(html.SvgCircle(html.ACx("5.5"), html.ACy("17.5"), html.AR("3.5"))),
		html.Child(html.SvgCircle(html.ACx("15"), html.ACy("5"), html.AR("1"))),
		html.Child(html.SvgPath(html.AD("M12 17.5V14l-3-3 4-3 2 3h2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
