package lucide

import (
	html "github.com/plainkit/html"
)

// SprayCan creates a Spray Can Lucide icon.
func SprayCan(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-spray-can", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 3h.01"))),
		html.Child(html.SvgPath(html.AD("M7 5h.01"))),
		html.Child(html.SvgPath(html.AD("M11 7h.01"))),
		html.Child(html.SvgPath(html.AD("M3 7h.01"))),
		html.Child(html.SvgPath(html.AD("M7 9h.01"))),
		html.Child(html.SvgPath(html.AD("M3 11h.01"))),
		html.Child(html.SvgRect(html.AWidth("4"), html.AHeight("4"), html.AX("15"), html.AY("5"))),
		html.Child(html.SvgPath(html.AD("m19 9 2 2v10c0 .6-.4 1-1 1h-6c-.6 0-1-.4-1-1V11l2-2"))),
		html.Child(html.SvgPath(html.AD("m13 14 8-2"))),
		html.Child(html.SvgPath(html.AD("m13 19 8-2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
