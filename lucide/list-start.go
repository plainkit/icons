package lucide

import (
	html "github.com/plainkit/html"
)

// ListStart creates a List Start Lucide icon.
func ListStart(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-start", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 5h6"))),
		html.Child(html.SvgPath(html.AD("M3 12h13"))),
		html.Child(html.SvgPath(html.AD("M3 19h13"))),
		html.Child(html.SvgPath(html.AD("m16 8-3-3 3-3"))),
		html.Child(html.SvgPath(html.AD("M21 19V7a2 2 0 0 0-2-2h-6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
