package lucide

import (
	html "github.com/plainkit/html"
)

// Computer creates a Computer Lucide icon.
func Computer(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-computer", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("14"), html.AHeight("8"), html.AX("5"), html.AY("2"), html.ARx("2")),
		html.SvgRect(html.AWidth("20"), html.AHeight("8"), html.AX("2"), html.AY("14"), html.ARx("2")),
		html.SvgPath(html.AD("M6 18h2")),
		html.SvgPath(html.AD("M12 18h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
