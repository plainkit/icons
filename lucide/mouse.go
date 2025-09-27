package lucide

import (
	html "github.com/plainkit/html"
)

// Mouse creates a Mouse Lucide icon.
func Mouse(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mouse", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("14"), html.AHeight("20"), html.AX("5"), html.AY("2"), html.ARx("7"))),
		html.Child(html.SvgPath(html.AD("M12 6v4"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
