package lucide

import (
	html "github.com/plainkit/html"
)

// MonitorDown creates a Monitor Down Lucide icon.
func MonitorDown(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-monitor-down", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 13V7")),
		html.SvgPath(html.AD("m15 10-3 3-3-3")),
		html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M12 17v4")),
		html.SvgPath(html.AD("M8 21h8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
