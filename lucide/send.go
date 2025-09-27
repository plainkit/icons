package lucide

import (
	html "github.com/plainkit/html"
)

// Send creates a Send Lucide icon.
func Send(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-send", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14.536 21.686a.5.5 0 0 0 .937-.024l6.5-19a.496.496 0 0 0-.635-.635l-19 6.5a.5.5 0 0 0-.024.937l7.93 3.18a2 2 0 0 1 1.112 1.11z"))),
		html.Child(html.SvgPath(html.AD("m21.854 2.147-10.94 10.939"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
