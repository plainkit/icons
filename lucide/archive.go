package lucide

import (
	html "github.com/plainkit/html"
)

// Archive creates a Archive Lucide icon.
func Archive(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-archive", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("5"), html.AX("2"), html.AY("3"), html.ARx("1"))),
		html.Child(html.SvgPath(html.AD("M4 8v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8"))),
		html.Child(html.SvgPath(html.AD("M10 12h4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
