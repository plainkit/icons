package lucide

import (
	html "github.com/plainkit/html"
)

// ParkingMeter creates a Parking Meter Lucide icon.
func ParkingMeter(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-parking-meter", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M11 15h2")),
		html.SvgPath(html.AD("M12 12v3")),
		html.SvgPath(html.AD("M12 19v3")),
		html.SvgPath(html.AD("M15.282 19a1 1 0 0 0 .948-.68l2.37-6.988a7 7 0 1 0-13.2 0l2.37 6.988a1 1 0 0 0 .948.68z")),
		html.SvgPath(html.AD("M9 9a3 3 0 1 1 6 0")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
