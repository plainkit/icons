package lucide

import (
	html "github.com/plainkit/html"
)

// Speaker creates a Speaker Lucide icon.
func Speaker(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-speaker", args)
	children := []html.SvgArg{
		html.Child(html.SvgRect(html.AWidth("16"), html.AHeight("20"), html.AX("4"), html.AY("2"), html.ARx("2"))),
		html.Child(html.SvgPath(html.AD("M12 6h.01"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("14"), html.AR("4"))),
		html.Child(html.SvgPath(html.AD("M12 14h.01"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
