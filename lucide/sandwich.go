package lucide

import (
	html "github.com/plainkit/html"
)

// Sandwich creates a Sandwich Lucide icon.
func Sandwich(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-sandwich", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m2.37 11.223 8.372-6.777a2 2 0 0 1 2.516 0l8.371 6.777")),
		html.SvgPath(html.AD("M21 15a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-5.25")),
		html.SvgPath(html.AD("M3 15a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h9")),
		html.SvgPath(html.AD("m6.67 15 6.13 4.6a2 2 0 0 0 2.8-.4l3.15-4.2")),
		html.SvgRect(html.AWidth("20"), html.AHeight("4"), html.AX("2"), html.AY("11"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
