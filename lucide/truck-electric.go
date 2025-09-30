package lucide

import (
	html "github.com/plainkit/html"
)

// TruckElectric creates a Truck Electric Lucide icon.
func TruckElectric(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-truck-electric", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 19V7a2 2 0 0 0-2-2H9")),
		html.SvgPath(html.AD("M15 19H9")),
		html.SvgPath(html.AD("M19 19h2a1 1 0 0 0 1-1v-3.65a1 1 0 0 0-.22-.62L18.3 9.38a1 1 0 0 0-.78-.38H14")),
		html.SvgPath(html.AD("M2 13v5a1 1 0 0 0 1 1h2")),
		html.SvgPath(html.AD("M4 3 2.15 5.15a.495.495 0 0 0 .35.86h2.15a.47.47 0 0 1 .35.86L3 9.02")),
		html.SvgCircle(html.ACx("17"), html.ACy("19"), html.AR("2")),
		html.SvgCircle(html.ACx("7"), html.ACy("19"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
