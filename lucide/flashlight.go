package lucide

import (
	html "github.com/plainkit/html"
)

// Flashlight creates a Flashlight Lucide icon.
func Flashlight(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-flashlight", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18 6c0 2-2 2-2 4v10a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V10c0-2-2-2-2-4V2h12z"))),
		html.Child(html.SvgLine(html.AX1("6"), html.AX2("18"), html.AY1("6"), html.AY2("6"))),
		html.Child(html.SvgLine(html.AX1("12"), html.AX2("12"), html.AY1("12"), html.AY2("12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
