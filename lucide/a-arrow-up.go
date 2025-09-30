package lucide

import (
	html "github.com/plainkit/html"
)

// AArrowUp creates a A Arrow Up Lucide icon.
func AArrowUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-a-arrow-up", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m14 11 4-4 4 4")),
		html.SvgPath(html.AD("M18 16V7")),
		html.SvgPath(html.AD("m2 16 4.039-9.69a.5.5 0 0 1 .923 0L11 16")),
		html.SvgPath(html.AD("M3.304 13h6.392")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
