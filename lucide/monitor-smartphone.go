package lucide

import (
	html "github.com/plainkit/html"
)

// MonitorSmartphone creates a Monitor Smartphone Lucide icon.
func MonitorSmartphone(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-monitor-smartphone", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M18 8V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h8")),
		html.SvgPath(html.AD("M10 19v-3.96 3.15")),
		html.SvgPath(html.AD("M7 19h5")),
		html.SvgRect(html.AWidth("6"), html.AHeight("10"), html.AX("16"), html.AY("12"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
