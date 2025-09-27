package lucide

import (
	html "github.com/plainkit/html"
)

// ClipboardClock creates a Clipboard Clock Lucide icon.
func ClipboardClock(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clipboard-clock", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 14v2.2l1.6 1"))),
		html.Child(html.SvgPath(html.AD("M16 4h2a2 2 0 0 1 2 2v.832"))),
		html.Child(html.SvgPath(html.AD("M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h2"))),
		html.Child(html.SvgCircle(html.ACx("16"), html.ACy("16"), html.AR("6"))),
		html.Child(html.SvgRect(html.AWidth("8"), html.AHeight("4"), html.AX("8"), html.AY("2"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
