package lucide

import (
	html "github.com/plainkit/html"
)

// Disc3 creates a Disc 3 Lucide icon.
func Disc3(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-disc-3", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgPath(html.AD("M6 12c0-1.7.7-3.2 1.8-4.2")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("2")),
		html.SvgPath(html.AD("M18 12c0 1.7-.7 3.2-1.8 4.2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
