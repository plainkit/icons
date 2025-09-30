package lucide

import (
	html "github.com/plainkit/html"
)

// WandSparkles creates a Wand Sparkles Lucide icon.
func WandSparkles(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-wand-sparkles", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m21.64 3.64-1.28-1.28a1.21 1.21 0 0 0-1.72 0L2.36 18.64a1.21 1.21 0 0 0 0 1.72l1.28 1.28a1.2 1.2 0 0 0 1.72 0L21.64 5.36a1.2 1.2 0 0 0 0-1.72")),
		html.SvgPath(html.AD("m14 7 3 3")),
		html.SvgPath(html.AD("M5 6v4")),
		html.SvgPath(html.AD("M19 14v4")),
		html.SvgPath(html.AD("M10 2v2")),
		html.SvgPath(html.AD("M7 8H3")),
		html.SvgPath(html.AD("M21 16h-4")),
		html.SvgPath(html.AD("M11 3H9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
