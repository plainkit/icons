package lucide

import (
	html "github.com/plainkit/html"
)

// MonitorCheck creates a Monitor Check Lucide icon.
func MonitorCheck(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-monitor-check", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m9 10 2 2 4-4"))),
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M12 17v4"))),
		html.Child(html.SvgPath(html.AD("M8 21h8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
