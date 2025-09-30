package lucide

import (
	html "github.com/plainkit/html"
)

// VibrateOff creates a Vibrate Off Lucide icon.
func VibrateOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-vibrate-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m2 8 2 2-2 2 2 2-2 2")),
		html.SvgPath(html.AD("m22 8-2 2 2 2-2 2 2 2")),
		html.SvgPath(html.AD("M8 8v10c0 .55.45 1 1 1h6c.55 0 1-.45 1-1v-2")),
		html.SvgPath(html.AD("M16 10.34V6c0-.55-.45-1-1-1h-4.34")),
		html.SvgLine(html.AX1("2"), html.AX2("22"), html.AY1("2"), html.AY2("22")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
