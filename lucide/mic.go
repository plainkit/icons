package lucide

import (
	html "github.com/plainkit/html"
)

// Mic creates a Mic Lucide icon.
func Mic(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mic", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 19v3"))),
		html.Child(html.SvgPath(html.AD("M19 10v2a7 7 0 0 1-14 0v-2"))),
		html.Child(html.SvgRect(html.AWidth("6"), html.AHeight("13"), html.AX("9"), html.AY("2"), html.ARx("3"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
