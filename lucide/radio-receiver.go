package lucide

import (
	html "github.com/plainkit/html"
)

// RadioReceiver creates a Radio Receiver Lucide icon.
func RadioReceiver(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-radio-receiver", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M5 16v2")),
		html.SvgPath(html.AD("M19 16v2")),
		html.SvgRect(html.AWidth("20"), html.AHeight("8"), html.AX("2"), html.AY("8"), html.ARx("2")),
		html.SvgPath(html.AD("M18 12h.01")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
