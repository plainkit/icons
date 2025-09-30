package lucide

import (
	html "github.com/plainkit/html"
)

// Banknote creates a Banknote Lucide icon.
func Banknote(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-banknote", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("20"), html.AHeight("12"), html.AX("2"), html.AY("6"), html.ARx("2")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("2")),
		html.SvgPath(html.AD("M6 12h.01M18 12h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
