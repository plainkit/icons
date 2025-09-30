package lucide

import (
	html "github.com/plainkit/html"
)

// UtilityPole creates a Utility Pole Lucide icon.
func UtilityPole(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-utility-pole", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 2v20")),
		html.SvgPath(html.AD("M2 5h20")),
		html.SvgPath(html.AD("M3 3v2")),
		html.SvgPath(html.AD("M7 3v2")),
		html.SvgPath(html.AD("M17 3v2")),
		html.SvgPath(html.AD("M21 3v2")),
		html.SvgPath(html.AD("m19 5-7 7-7-7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
