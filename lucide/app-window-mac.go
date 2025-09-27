package lucide

import (
	html "github.com/plainkit/html"
)

// AppWindowMac creates a App Window Mac Lucide icon.
func AppWindowMac(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-app-window-mac", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("16"), html.AX("2"), html.AY("4"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M6 8h.01"))),
		html.Child(html.SvgPath(html.AD("M10 8h.01"))),
		html.Child(html.SvgPath(html.AD("M14 8h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
