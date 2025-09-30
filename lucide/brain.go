package lucide

import (
	html "github.com/plainkit/html"
)

// Brain creates a Brain Lucide icon.
func Brain(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-brain", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 18V5")),
		html.SvgPath(html.AD("M15 13a4.17 4.17 0 0 1-3-4 4.17 4.17 0 0 1-3 4")),
		html.SvgPath(html.AD("M17.598 6.5A3 3 0 1 0 12 5a3 3 0 1 0-5.598 1.5")),
		html.SvgPath(html.AD("M17.997 5.125a4 4 0 0 1 2.526 5.77")),
		html.SvgPath(html.AD("M18 18a4 4 0 0 0 2-7.464")),
		html.SvgPath(html.AD("M19.967 17.483A4 4 0 1 1 12 18a4 4 0 1 1-7.967-.517")),
		html.SvgPath(html.AD("M6 18a4 4 0 0 1-2-7.464")),
		html.SvgPath(html.AD("M6.003 5.125a4 4 0 0 0-2.526 5.77")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
