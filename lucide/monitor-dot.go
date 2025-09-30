package lucide

import (
	html "github.com/plainkit/html"
)

// MonitorDot creates a Monitor Dot Lucide icon.
func MonitorDot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-monitor-dot", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 17v4")),
		html.SvgPath(html.AD("M22 12.307V15a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h8.693")),
		html.SvgPath(html.AD("M8 21h8")),
		html.SvgCircle(html.ACx("19"), html.ACy("6"), html.AR("3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
