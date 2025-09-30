package lucide

import (
	html "github.com/plainkit/html"
)

// BadgeSwissFranc creates a Badge Swiss Franc Lucide icon.
func BadgeSwissFranc(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-badge-swiss-franc", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z")),
		html.SvgPath(html.AD("M11 17V8h4")),
		html.SvgPath(html.AD("M11 12h3")),
		html.SvgPath(html.AD("M9 16h4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
