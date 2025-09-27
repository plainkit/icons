package lucide

import (
	html "github.com/plainkit/html"
)

// Touchpad creates a Touchpad Lucide icon.
func Touchpad(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-touchpad", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("20"), html.AHeight("16"), html.AX("2"), html.AY("4"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M2 14h20"))),
		html.Child(html.SvgPath(html.AD("M12 20v-6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
