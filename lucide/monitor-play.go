package lucide

import (
	html "github.com/plainkit/html"
)

// MonitorPlay creates a Monitor Play Lucide icon.
func MonitorPlay(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-monitor-play", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M15.033 9.44a.647.647 0 0 1 0 1.12l-4.065 2.352a.645.645 0 0 1-.968-.56V7.648a.645.645 0 0 1 .967-.56z")),
		html.SvgPath(html.AD("M12 17v4")),
		html.SvgPath(html.AD("M8 21h8")),
		html.SvgRect(html.AWidth("20"), html.AHeight("14"), html.AX("2"), html.AY("3"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
