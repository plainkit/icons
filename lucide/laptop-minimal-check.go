package lucide

import (
	html "github.com/plainkit/html"
)

// LaptopMinimalCheck creates a Laptop Minimal Check Lucide icon.
func LaptopMinimalCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-laptop-minimal-check", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 20h20")),
		html.SvgPath(html.AD("m9 10 2 2 4-4")),
		html.SvgRect(html.AWidth("18"), html.AHeight("12"), html.AX("3"), html.AY("4"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
