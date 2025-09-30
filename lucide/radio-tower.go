package lucide

import (
	html "github.com/plainkit/html"
)

// RadioTower creates a Radio Tower Lucide icon.
func RadioTower(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-radio-tower", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4.9 16.1C1 12.2 1 5.8 4.9 1.9")),
		html.SvgPath(html.AD("M7.8 4.7a6.14 6.14 0 0 0-.8 7.5")),
		html.SvgCircle(html.ACx("12"), html.ACy("9"), html.AR("2")),
		html.SvgPath(html.AD("M16.2 4.8c2 2 2.26 5.11.8 7.47")),
		html.SvgPath(html.AD("M19.1 1.9a9.96 9.96 0 0 1 0 14.1")),
		html.SvgPath(html.AD("M9.5 18h5")),
		html.SvgPath(html.AD("m8 22 4-11 4 11")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
