package lucide

import (
	html "github.com/plainkit/html"
)

// MonitorStop creates a Monitor Stop Lucide icon.
func MonitorStop(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-monitor-stop", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 17v4"))),
		html.Child(html.SvgPath(html.AD("M8 21h8"))),
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("3"), html.ARx("2"))),
		html.Child(html.SvgRect(html.AWidth("6"), html.AHeight("6"), html.AX("9"), html.AY("7"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
