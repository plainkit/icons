package lucide

import (
	html "github.com/plainkit/html"
)

// Notebook creates a Notebook Lucide icon.
func Notebook(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-notebook", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 6h4")),
		html.SvgPath(html.AD("M2 10h4")),
		html.SvgPath(html.AD("M2 14h4")),
		html.SvgPath(html.AD("M2 18h4")),
		html.SvgRect(html.AWidth("16"), html.AHeight("20"), html.AX("4"), html.AY("2"), html.ARx("2")),
		html.SvgPath(html.AD("M16 2v20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
