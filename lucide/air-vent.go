package lucide

import (
	html "github.com/plainkit/html"
)

// AirVent creates a Air Vent Lucide icon.
func AirVent(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-air-vent", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M18 17.5a2.5 2.5 0 1 1-4 2.03V12"))),
		html.Child(html.SvgPath(html.AD("M6 12H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2"))),
		html.Child(html.SvgPath(html.AD("M6 8h12"))),
		html.Child(html.SvgPath(html.AD("M6.6 15.572A2 2 0 1 0 10 17v-5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
