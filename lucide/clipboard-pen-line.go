package lucide

import (
	html "github.com/plainkit/html"
)

// ClipboardPenLine creates a Clipboard Pen Line Lucide icon.
func ClipboardPenLine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clipboard-pen-line", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("8"), html.AHeight("4"), html.AX("8"), html.AY("2"), html.ARx("1")),
		html.SvgPath(html.AD("M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-.5")),
		html.SvgPath(html.AD("M16 4h2a2 2 0 0 1 1.73 1")),
		html.SvgPath(html.AD("M8 18h1")),
		html.SvgPath(html.AD("M21.378 12.626a1 1 0 0 0-3.004-3.004l-4.01 4.012a2 2 0 0 0-.506.854l-.837 2.87a.5.5 0 0 0 .62.62l2.87-.837a2 2 0 0 0 .854-.506z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
