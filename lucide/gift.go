package lucide

import (
	html "github.com/plainkit/html"
)

// Gift creates a Gift Lucide icon.
func Gift(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-gift", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("18"), html.AHeight("4"), html.AX("3"), html.AY("8"), html.ARx("1")),
		html.SvgPath(html.AD("M12 8v13")),
		html.SvgPath(html.AD("M19 12v7a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-7")),
		html.SvgPath(html.AD("M7.5 8a2.5 2.5 0 0 1 0-5A4.8 8 0 0 1 12 8a4.8 8 0 0 1 4.5-5 2.5 2.5 0 0 1 0 5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
