package lucide

import (
	html "github.com/plainkit/html"
)

// ClipboardPaste creates a Clipboard Paste Lucide icon.
func ClipboardPaste(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clipboard-paste", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M11 14h10"))),
		html.Child(html.SvgPath(html.AD("M16 4h2a2 2 0 0 1 2 2v1.344"))),
		html.Child(html.SvgPath(html.AD("m17 18 4-4-4-4"))),
		html.Child(html.SvgPath(html.AD("M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 1.793-1.113"))),
		html.Child(html.SvgRect(html.AWidth("8"), html.AHeight("4"), html.AX("8"), html.AY("2"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
