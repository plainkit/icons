package lucide

import (
	html "github.com/plainkit/html"
)

// AppWindow creates a App Window Lucide icon.
func AppWindow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-app-window", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("20"), html.AHeight("16"), html.AX("2"), html.AY("4"), html.ARx("2")),
		html.SvgPath(html.AD("M10 4v4")),
		html.SvgPath(html.AD("M2 8h20")),
		html.SvgPath(html.AD("M6 4v4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
