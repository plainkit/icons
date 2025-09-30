package lucide

import (
	html "github.com/plainkit/html"
)

// Glasses creates a Glasses Lucide icon.
func Glasses(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-glasses", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("6"), html.ACy("15"), html.AR("4")),
		html.SvgCircle(html.ACx("18"), html.ACy("15"), html.AR("4")),
		html.SvgPath(html.AD("M14 15a2 2 0 0 0-2-2 2 2 0 0 0-2 2")),
		html.SvgPath(html.AD("M2.5 13 5 7c.7-1.3 1.4-2 3-2")),
		html.SvgPath(html.AD("M21.5 13 19 7c-.7-1.3-1.5-2-3-2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
