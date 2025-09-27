package lucide

import (
	html "github.com/plainkit/html"
)

// Keyboard creates a Keyboard Lucide icon.
func Keyboard(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-keyboard", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10 8h.01"))),
		html.Child(html.SvgPath(html.AD("M12 12h.01"))),
		html.Child(html.SvgPath(html.AD("M14 8h.01"))),
		html.Child(html.SvgPath(html.AD("M16 12h.01"))),
		html.Child(html.SvgPath(html.AD("M18 8h.01"))),
		html.Child(html.SvgPath(html.AD("M6 8h.01"))),
		html.Child(html.SvgPath(html.AD("M7 16h10"))),
		html.Child(html.SvgPath(html.AD("M8 12h.01"))),
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("16"), html.AX("2"), html.AY("4"), html.ARx("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
