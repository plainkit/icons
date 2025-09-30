package lucide

import (
	html "github.com/plainkit/html"
)

// Baby creates a Baby Lucide icon.
func Baby(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-baby", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 16c.5.3 1.2.5 2 .5s1.5-.2 2-.5")),
		html.SvgPath(html.AD("M15 12h.01")),
		html.SvgPath(html.AD("M19.38 6.813A9 9 0 0 1 20.8 10.2a2 2 0 0 1 0 3.6 9 9 0 0 1-17.6 0 2 2 0 0 1 0-3.6A9 9 0 0 1 12 3c2 0 3.5 1.1 3.5 2.5s-.9 2.5-2 2.5c-.8 0-1.5-.4-1.5-1")),
		html.SvgPath(html.AD("M9 12h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
