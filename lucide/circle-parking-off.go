package lucide

import (
	html "github.com/plainkit/html"
)

// CircleParkingOff creates a Circle Parking Off Lucide icon.
func CircleParkingOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-parking-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12.656 7H13a3 3 0 0 1 2.984 3.307")),
		html.SvgPath(html.AD("M13 13H9")),
		html.SvgPath(html.AD("M19.071 19.071A1 1 0 0 1 4.93 4.93")),
		html.SvgPath(html.AD("m2 2 20 20")),
		html.SvgPath(html.AD("M8.357 2.687a10 10 0 0 1 12.956 12.956")),
		html.SvgPath(html.AD("M9 17V9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
