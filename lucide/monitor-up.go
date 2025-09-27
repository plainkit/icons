package lucide

import (
	html "github.com/plainkit/html"
)

// MonitorUp creates a Monitor Up Lucide icon.
func MonitorUp(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-monitor-up", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m9 10 3-3 3 3"))),
		html.Child(html.SvgPath(html.AD("M12 13V7"))),
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M12 17v4"))),
		html.Child(html.SvgPath(html.AD("M8 21h8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
