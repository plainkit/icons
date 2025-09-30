package lucide

import (
	html "github.com/plainkit/html"
)

// Chromium creates a Chromium Lucide icon.
func Chromium(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-chromium", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10.88 21.94 15.46 14")),
		html.SvgPath(html.AD("M21.17 8H12")),
		html.SvgPath(html.AD("M3.95 6.06 8.54 14")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
