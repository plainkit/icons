package lucide

import (
	html "github.com/plainkit/html"
)

// GeorgianLari creates a Georgian Lari Lucide icon.
func GeorgianLari(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-georgian-lari", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M11.5 21a7.5 7.5 0 1 1 7.35-9")),
		html.SvgPath(html.AD("M13 12V3")),
		html.SvgPath(html.AD("M4 21h16")),
		html.SvgPath(html.AD("M9 12V3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
