package lucide

import (
	html "github.com/plainkit/html"
)

// HouseWifi creates a House Wifi Lucide icon.
func HouseWifi(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-house-wifi", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M9.5 13.866a4 4 0 0 1 5 .01"))),
		html.Child(html.SvgPath(html.AD("M12 17h.01"))),
		html.Child(html.SvgPath(html.AD("M3 10a2 2 0 0 1 .709-1.528l7-6a2 2 0 0 1 2.582 0l7 6A2 2 0 0 1 21 10v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"))),
		html.Child(html.SvgPath(html.AD("M7 10.754a8 8 0 0 1 10 0"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
