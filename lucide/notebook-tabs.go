package lucide

import (
	html "github.com/plainkit/html"
)

// NotebookTabs creates a Notebook Tabs Lucide icon.
func NotebookTabs(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-notebook-tabs", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 6h4"))),
		html.Child(html.SvgPath(html.AD("M2 10h4"))),
		html.Child(html.SvgPath(html.AD("M2 14h4"))),
		html.Child(html.SvgPath(html.AD("M2 18h4"))),
		html.Child(html.SvgRect(html.AWidth("16"), html.AHeight("20"), html.AX("4"), html.AY("2"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M15 2v20"))),
		html.Child(html.SvgPath(html.AD("M15 7h5"))),
		html.Child(html.SvgPath(html.AD("M15 12h5"))),
		html.Child(html.SvgPath(html.AD("M15 17h5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
