package lucide

import (
	html "github.com/plainkit/html"
)

// ClipboardType creates a Clipboard Type Lucide icon.
func ClipboardType(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clipboard-type", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("8"), html.AHeight("4"), html.AX("8"), html.AY("2"), html.ARx("1"), html.ARy("1"))),
		html.Child(html.SvgPath(html.AD("M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"))),
		html.Child(html.SvgPath(html.AD("M9 12v-1h6v1"))),
		html.Child(html.SvgPath(html.AD("M11 17h2"))),
		html.Child(html.SvgPath(html.AD("M12 11v6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
