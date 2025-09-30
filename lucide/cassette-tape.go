package lucide

import (
	html "github.com/plainkit/html"
)

// CassetteTape creates a Cassette Tape Lucide icon.
func CassetteTape(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cassette-tape", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("20"), html.AHeight("16"), html.AX("2"), html.AY("4"), html.ARx("2")),
		html.SvgCircle(html.ACx("8"), html.ACy("10"), html.AR("2")),
		html.SvgPath(html.AD("M8 12h8")),
		html.SvgCircle(html.ACx("16"), html.ACy("10"), html.AR("2")),
		html.SvgPath(html.AD("m6 20 .7-2.9A1.4 1.4 0 0 1 8.1 16h7.8a1.4 1.4 0 0 1 1.4 1l.7 3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
