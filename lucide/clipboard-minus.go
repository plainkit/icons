package lucide

import (
	html "github.com/plainkit/html"
)

// ClipboardMinus creates a Clipboard Minus Lucide icon.
func ClipboardMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-clipboard-minus", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("8"), html.AHeight("4"), html.AX("8"), html.AY("2"), html.ARx("1"), html.ARy("1")),
		html.SvgPath(html.AD("M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2")),
		html.SvgPath(html.AD("M9 14h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
