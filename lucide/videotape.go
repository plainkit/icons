package lucide

import (
	html "github.com/plainkit/html"
)

// Videotape creates a Videotape Lucide icon.
func Videotape(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-videotape", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("16"), html.AX("2"), html.AY("4"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M2 8h20"))),
		html.Child(html.SvgCircle(html.ACx("8"), html.ACy("14"), html.AR("2"))),
		html.Child(html.SvgPath(html.AD("M8 12h8"))),
		html.Child(html.SvgCircle(html.ACx("16"), html.ACy("14"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
