package lucide

import (
	html "github.com/plainkit/html"
)

// Hash creates a Hash Lucide icon.
func Hash(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-hash", args)
	children := []html.SvgArg{
		html.SvgLine(html.AX1("4"), html.AX2("20"), html.AY1("9"), html.AY2("9")),
		html.SvgLine(html.AX1("4"), html.AX2("20"), html.AY1("15"), html.AY2("15")),
		html.SvgLine(html.AX1("10"), html.AX2("8"), html.AY1("3"), html.AY2("21")),
		html.SvgLine(html.AX1("16"), html.AX2("14"), html.AY1("3"), html.AY2("21")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
