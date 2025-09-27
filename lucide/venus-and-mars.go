package lucide

import (
	html "github.com/plainkit/html"
)

// VenusAndMars creates a Venus And Mars Lucide icon.
func VenusAndMars(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-venus-and-mars", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 20h4"))),
		html.Child(html.SvgPath(html.AD("M12 16v6"))),
		html.Child(html.SvgPath(html.AD("M17 2h4v4"))),
		html.Child(html.SvgPath(html.AD("m21 2-5.46 5.46"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("11"), html.AR("5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
