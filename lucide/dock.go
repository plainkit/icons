package lucide

import (
	html "github.com/plainkit/html"
)

// Dock creates a Dock Lucide icon.
func Dock(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-dock", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2 8h20")),
		html.SvgRect(html.AWidth("20"), html.AHeight("16"), html.AX("2"), html.AY("4"), html.ARx("2")),
		html.SvgPath(html.AD("M6 16h12")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
