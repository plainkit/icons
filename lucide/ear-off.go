package lucide

import (
	html "github.com/plainkit/html"
)

// EarOff creates a Ear Off Lucide icon.
func EarOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ear-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M6 18.5a3.5 3.5 0 1 0 7 0c0-1.57.92-2.52 2.04-3.46")),
		html.SvgPath(html.AD("M6 8.5c0-.75.13-1.47.36-2.14")),
		html.SvgPath(html.AD("M8.8 3.15A6.5 6.5 0 0 1 19 8.5c0 1.63-.44 2.81-1.09 3.76")),
		html.SvgPath(html.AD("M12.5 6A2.5 2.5 0 0 1 15 8.5M10 13a2 2 0 0 0 1.82-1.18")),
		html.SvgLine(html.AX1("2"), html.AX2("22"), html.AY1("2"), html.AY2("22")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
