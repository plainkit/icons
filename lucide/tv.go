package lucide

import (
	html "github.com/plainkit/html"
)

// Tv creates a Tv Lucide icon.
func Tv(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-tv", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m17 2-5 5-5-5"))),
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("15"), html.AX("2"), html.AY("7"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
