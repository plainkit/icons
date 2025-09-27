package lucide

import (
	html "github.com/plainkit/html"
)

// MonitorSpeaker creates a Monitor Speaker Lucide icon.
func MonitorSpeaker(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-monitor-speaker", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M5.5 20H8"))),
		html.Child(html.SvgPath(html.AD("M17 9h.01"))),
		html.Child(html.SvgRect(html.AWidth("10"), html.AHeight("16"), html.AX("12"), html.AY("4"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M8 6H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h4"))),
		html.Child(html.SvgCircle(html.ACx("17"), html.ACy("15"), html.AR("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
