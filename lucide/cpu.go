package lucide

import (
	html "github.com/plainkit/html"
)

// Cpu creates a Cpu Lucide icon.
func Cpu(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cpu", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 20v2")),
		html.SvgPath(html.AD("M12 2v2")),
		html.SvgPath(html.AD("M17 20v2")),
		html.SvgPath(html.AD("M17 2v2")),
		html.SvgPath(html.AD("M2 12h2")),
		html.SvgPath(html.AD("M2 17h2")),
		html.SvgPath(html.AD("M2 7h2")),
		html.SvgPath(html.AD("M20 12h2")),
		html.SvgPath(html.AD("M20 17h2")),
		html.SvgPath(html.AD("M20 7h2")),
		html.SvgPath(html.AD("M7 20v2")),
		html.SvgPath(html.AD("M7 2v2")),
		html.SvgRect(html.AWidth("16"), html.AHeight("16"), html.AX("4"), html.AY("4"), html.ARx("2")),
		html.SvgRect(html.AWidth("8"), html.AHeight("8"), html.AX("8"), html.AY("8"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
