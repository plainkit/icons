package lucide

import (
	html "github.com/plainkit/html"
)

// ClipboardCopy creates a Clipboard Copy Lucide icon.
func ClipboardCopy(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clipboard-copy", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("8"), html.AHeight("4"), html.AX("8"), html.AY("2"), html.ARx("1"), html.ARy("1")),
		html.SvgPath(html.AD("M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-2")),
		html.SvgPath(html.AD("M16 4h2a2 2 0 0 1 2 2v4")),
		html.SvgPath(html.AD("M21 14H11")),
		html.SvgPath(html.AD("m15 10-4 4 4 4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
