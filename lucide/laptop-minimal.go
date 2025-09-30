package lucide

import (
	html "github.com/plainkit/html"
)

// LaptopMinimal creates a Laptop Minimal Lucide icon.
func LaptopMinimal(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-laptop-minimal", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("12"), html.AX("3"), html.AY("4"), html.ARx("2"), html.ARy("2")),
		html.SvgLine(html.AX1("2"), html.AX2("22"), html.AY1("20"), html.AY2("20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
