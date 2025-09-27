package lucide

import (
	html "github.com/plainkit/html"
)

// EarthLock creates a Earth Lock Lucide icon.
func EarthLock(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-earth-lock", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M7 3.34V5a3 3 0 0 0 3 3"))),
		html.Child(html.SvgPath(html.AD("M11 21.95V18a2 2 0 0 0-2-2 2 2 0 0 1-2-2v-1a2 2 0 0 0-2-2H2.05"))),
		html.Child(html.SvgPath(html.AD("M21.54 15H17a2 2 0 0 0-2 2v4.54"))),
		html.Child(html.SvgPath(html.AD("M12 2a10 10 0 1 0 9.54 13"))),
		html.Child(html.SvgPath(html.AD("M20 6V4a2 2 0 1 0-4 0v2"))),
		html.Child(html.SvgRect(html.AWidth("8"), html.AHeight("5"), html.AX("14"), html.AY("6"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
