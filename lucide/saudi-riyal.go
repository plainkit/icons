package lucide

import (
	html "github.com/plainkit/html"
)

// SaudiRiyal creates a Saudi Riyal Lucide icon.
func SaudiRiyal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-saudi-riyal", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m20 19.5-5.5 1.2"))),
		html.Child(html.SvgPath(html.AD("M14.5 4v11.22a1 1 0 0 0 1.242.97L20 15.2"))),
		html.Child(html.SvgPath(html.AD("m2.978 19.351 5.549-1.363A2 2 0 0 0 10 16V2"))),
		html.Child(html.SvgPath(html.AD("M20 10 4 13.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
