package lucide

import (
	html "github.com/plainkit/html"
)

// Proportions creates a Proportions Lucide icon.
func Proportions(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-proportions", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("20"), html.AHeight("16"), html.AX("2"), html.AY("4"), html.ARx("2")),
		html.SvgPath(html.AD("M12 9v11")),
		html.SvgPath(html.AD("M2 9h13a2 2 0 0 1 2 2v9")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
