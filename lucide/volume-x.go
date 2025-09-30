package lucide

import (
	html "github.com/plainkit/html"
)

// VolumeX creates a Volume X Lucide icon.
func VolumeX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-volume-x", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M11 4.702a.705.705 0 0 0-1.203-.498L6.413 7.587A1.4 1.4 0 0 1 5.416 8H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2.416a1.4 1.4 0 0 1 .997.413l3.383 3.384A.705.705 0 0 0 11 19.298z")),
		html.SvgLine(html.AX1("22"), html.AX2("16"), html.AY1("9"), html.AY2("15")),
		html.SvgLine(html.AX1("16"), html.AX2("22"), html.AY1("9"), html.AY2("15")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
