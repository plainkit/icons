package lucide

import (
	html "github.com/plainkit/html"
)

// CloudOff creates a Cloud Off Lucide icon.
func CloudOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cloud-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m2 2 20 20")),
		html.SvgPath(html.AD("M5.782 5.782A7 7 0 0 0 9 19h8.5a4.5 4.5 0 0 0 1.307-.193")),
		html.SvgPath(html.AD("M21.532 16.5A4.5 4.5 0 0 0 17.5 10h-1.79A7.008 7.008 0 0 0 10 5.07")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
