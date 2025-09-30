package lucide

import (
	html "github.com/plainkit/html"
)

// Columns3 creates a Columns 3 Lucide icon.
func Columns3(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-columns-3", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M9 3v18")),
		html.SvgPath(html.AD("M15 3v18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
