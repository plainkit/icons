package lucide

import (
	html "github.com/plainkit/html"
)

// ShipWheel creates a Ship Wheel Lucide icon.
func ShipWheel(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ship-wheel", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("8")),
		html.SvgPath(html.AD("M12 2v7.5")),
		html.SvgPath(html.AD("m19 5-5.23 5.23")),
		html.SvgPath(html.AD("M22 12h-7.5")),
		html.SvgPath(html.AD("m19 19-5.23-5.23")),
		html.SvgPath(html.AD("M12 14.5V22")),
		html.SvgPath(html.AD("M10.23 13.77 5 19")),
		html.SvgPath(html.AD("M9.5 12H2")),
		html.SvgPath(html.AD("M10.23 10.23 5 5")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("2.5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
