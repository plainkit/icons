package lucide

import (
	html "github.com/plainkit/html"
)

// ClipboardPen creates a Clipboard Pen Lucide icon.
func ClipboardPen(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clipboard-pen", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("8"), html.AHeight("4"), html.AX("8"), html.AY("2"), html.ARx("1")),
		html.SvgPath(html.AD("M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-5.5")),
		html.SvgPath(html.AD("M4 13.5V6a2 2 0 0 1 2-2h2")),
		html.SvgPath(html.AD("M13.378 15.626a1 1 0 1 0-3.004-3.004l-5.01 5.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
