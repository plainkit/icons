package lucide

import (
	html "github.com/plainkit/html"
)

// FerrisWheel creates a Ferris Wheel Lucide icon.
func FerrisWheel(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ferris-wheel", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("2")),
		html.SvgPath(html.AD("M12 2v4")),
		html.SvgPath(html.AD("m6.8 15-3.5 2")),
		html.SvgPath(html.AD("m20.7 7-3.5 2")),
		html.SvgPath(html.AD("M6.8 9 3.3 7")),
		html.SvgPath(html.AD("m20.7 17-3.5-2")),
		html.SvgPath(html.AD("m9 22 3-8 3 8")),
		html.SvgPath(html.AD("M8 22h8")),
		html.SvgPath(html.AD("M18 18.7a9 9 0 1 0-12 0")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
