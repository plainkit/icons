package lucide

import (
	html "github.com/plainkit/html"
)

// Euro creates a Euro Lucide icon.
func Euro(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-euro", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M4 10h12")),
		html.SvgPath(html.AD("M4 14h9")),
		html.SvgPath(html.AD("M19 6a7.7 7.7 0 0 0-5.2-2A7.9 7.9 0 0 0 6 12c0 4.4 3.5 8 7.8 8 2 0 3.8-.8 5.2-2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
