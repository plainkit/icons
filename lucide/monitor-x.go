package lucide

import (
	html "github.com/plainkit/html"
)

// MonitorX creates a Monitor X Lucide icon.
func MonitorX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-monitor-x", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m14.5 12.5-5-5")),
		html.SvgPath(html.AD("m9.5 12.5 5-5")),
		html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M12 17v4")),
		html.SvgPath(html.AD("M8 21h8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
