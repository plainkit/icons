package lucide

import (
	html "github.com/plainkit/html"
)

// FlashlightOff creates a Flashlight Off Lucide icon.
func FlashlightOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-flashlight-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 16v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V10c0-2-2-2-2-4"))),
		html.Child(html.SvgPath(html.AD("M7 2h11v4c0 2-2 2-2 4v1"))),
		html.Child(html.SvgLine(html.AX1("11"), html.AX2("18"), html.AY1("6"), html.AY2("6"))),
		html.Child(html.SvgLine(html.AX1("2"), html.AX2("22"), html.AY1("2"), html.AY2("22"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
