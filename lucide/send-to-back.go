package lucide

import (
	html "github.com/plainkit/html"
)

// SendToBack creates a Send To Back Lucide icon.
func SendToBack(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-send-to-back", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("8"), html.AHeight("8"), html.AX("14"), html.AY("14"), html.ARx("2")),
		html.SvgRect(html.AWidth("8"), html.AHeight("8"), html.AX("2"), html.AY("2"), html.ARx("2")),
		html.SvgPath(html.AD("M7 14v1a2 2 0 0 0 2 2h1")),
		html.SvgPath(html.AD("M14 7h1a2 2 0 0 1 2 2v1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
