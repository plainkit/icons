package lucide

import (
	html "github.com/plainkit/html"
)

// EvCharger creates a Ev Charger Lucide icon.
func EvCharger(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-ev-charger", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14 13h2a2 2 0 0 1 2 2v2a2 2 0 0 0 4 0v-6.998a2 2 0 0 0-.59-1.42L18 5"))),
		html.Child(html.SvgPath(html.AD("M14 21V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v16"))),
		html.Child(html.SvgPath(html.AD("M2 21h13"))),
		html.Child(html.SvgPath(html.AD("M3 7h11"))),
		html.Child(html.SvgPath(html.AD("m9 11-2 3h3l-2 3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
