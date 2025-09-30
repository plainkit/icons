package lucide

import (
	html "github.com/plainkit/html"
)

// HardDrive creates a Hard Drive Lucide icon.
func HardDrive(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-hard-drive", args)
	children := []html.SvgArg{
		html.SvgLine(html.AX1("22"), html.AX2("2"), html.AY1("12"), html.AY2("12")),
		html.SvgPath(html.AD("M5.45 5.11 2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11z")),
		html.SvgLine(html.AX1("6"), html.AX2("6.01"), html.AY1("16"), html.AY2("16")),
		html.SvgLine(html.AX1("10"), html.AX2("10.01"), html.AY1("16"), html.AY2("16")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
