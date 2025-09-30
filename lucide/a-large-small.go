package lucide

import (
	html "github.com/plainkit/html"
)

// ALargeSmall creates a A Large Small Lucide icon.
func ALargeSmall(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-a-large-small", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m15 16 2.536-7.328a1.02 1.02 1 0 1 1.928 0L22 16")),
		html.SvgPath(html.AD("M15.697 14h5.606")),
		html.SvgPath(html.AD("m2 16 4.039-9.69a.5.5 0 0 1 .923 0L11 16")),
		html.SvgPath(html.AD("M3.304 13h6.392")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
