package lucide

import (
	html "github.com/plainkit/html"
)

// Scale creates a Scale Lucide icon.
func Scale(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-scale", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m16 16 3-8 3 8c-.87.65-1.92 1-3 1s-2.13-.35-3-1Z")),
		html.SvgPath(html.AD("m2 16 3-8 3 8c-.87.65-1.92 1-3 1s-2.13-.35-3-1Z")),
		html.SvgPath(html.AD("M7 21h10")),
		html.SvgPath(html.AD("M12 3v18")),
		html.SvgPath(html.AD("M3 7h2c2 0 5-1 7-2 2 1 5 2 7 2h2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
