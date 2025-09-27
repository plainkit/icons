package lucide

import (
	html "github.com/plainkit/html"
)

// KeyboardOff creates a Keyboard Off Lucide icon.
func KeyboardOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-keyboard-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M 20 4 A2 2 0 0 1 22 6"))),
		html.Child(html.SvgPath(html.AD("M 22 6 L 22 16.41"))),
		html.Child(html.SvgPath(html.AD("M 7 16 L 16 16"))),
		html.Child(html.SvgPath(html.AD("M 9.69 4 L 20 4"))),
		html.Child(html.SvgPath(html.AD("M14 8h.01"))),
		html.Child(html.SvgPath(html.AD("M18 8h.01"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M20 20H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2"))),
		html.Child(html.SvgPath(html.AD("M6 8h.01"))),
		html.Child(html.SvgPath(html.AD("M8 12h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
