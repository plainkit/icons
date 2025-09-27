package lucide

import (
	html "github.com/plainkit/html"
)

// Slack creates a Slack Lucide icon.
func Slack(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-slack", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("3"), html.AHeight("8"), html.AX("13"), html.AY("2"), html.ARx("1.5"))),
		html.Child(html.SvgPath(html.AD("M19 8.5V10h1.5A1.5 1.5 0 1 0 19 8.5"))),
		html.Child(html.SvgRect(html.AWidth("3"), html.AHeight("8"), html.AX("8"), html.AY("14"), html.ARx("1.5"))),
		html.Child(html.SvgPath(html.AD("M5 15.5V14H3.5A1.5 1.5 0 1 0 5 15.5"))),
		html.Child(html.SvgRect(html.AWidth("8"), html.AHeight("3"), html.AX("14"), html.AY("13"), html.ARx("1.5"))),
		html.Child(html.SvgPath(html.AD("M15.5 19H14v1.5a1.5 1.5 0 1 0 1.5-1.5"))),
		html.Child(html.SvgRect(html.AWidth("8"), html.AHeight("3"), html.AX("2"), html.AY("8"), html.ARx("1.5"))),
		html.Child(html.SvgPath(html.AD("M8.5 5H10V3.5A1.5 1.5 0 1 0 8.5 5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
