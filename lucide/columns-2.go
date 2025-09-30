package lucide

import (
	html "github.com/plainkit/html"
)

// Columns2 creates a Columns 2 Lucide icon.
func Columns2(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-columns-2", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("18"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M12 3v18")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
