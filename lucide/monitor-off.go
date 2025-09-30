package lucide

import (
	html "github.com/plainkit/html"
)

// MonitorOff creates a Monitor Off Lucide icon.
func MonitorOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-monitor-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M17 17H4a2 2 0 0 1-2-2V5c0-1.5 1-2 1-2")),
		html.SvgPath(html.AD("M22 15V5a2 2 0 0 0-2-2H9")),
		html.SvgPath(html.AD("M8 21h8")),
		html.SvgPath(html.AD("M12 17v4")),
		html.SvgPath(html.AD("m2 2 20 20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
