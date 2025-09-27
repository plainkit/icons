package lucide

import (
	html "github.com/plainkit/html"
)

// LocateOff creates a Locate Off Lucide icon.
func LocateOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-locate-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 19v3"))),
		html.Child(html.SvgPath(html.AD("M12 2v3"))),
		html.Child(html.SvgPath(html.AD("M18.89 13.24a7 7 0 0 0-8.13-8.13"))),
		html.Child(html.SvgPath(html.AD("M19 12h3"))),
		html.Child(html.SvgPath(html.AD("M2 12h3"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M7.05 7.05a7 7 0 0 0 9.9 9.9"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
