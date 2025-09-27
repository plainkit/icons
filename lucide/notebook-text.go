package lucide

import (
	html "github.com/plainkit/html"
)

// NotebookText creates a Notebook Text Lucide icon.
func NotebookText(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-notebook-text", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 6h4"))),
		html.Child(html.SvgPath(html.AD("M2 10h4"))),
		html.Child(html.SvgPath(html.AD("M2 14h4"))),
		html.Child(html.SvgPath(html.AD("M2 18h4"))),
		html.Child(html.SvgRect(html.AWidth("16"), html.AHeight("20"), html.AX("4"), html.AY("2"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M9.5 8h5"))),
		html.Child(html.SvgPath(html.AD("M9.5 12H16"))),
		html.Child(html.SvgPath(html.AD("M9.5 16H14"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
