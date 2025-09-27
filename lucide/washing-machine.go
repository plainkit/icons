package lucide

import (
	html "github.com/plainkit/html"
)

// WashingMachine creates a Washing Machine Lucide icon.
func WashingMachine(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-washing-machine", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M3 6h3"))),
		html.Child(html.SvgPath(html.AD("M17 6h.01"))),
		html.Child(html.SvgRect(html.AWidth("18"), html.AHeight("20"), html.AX("3"), html.AY("2"), html.ARx("2"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("13"), html.AR("5"))),
		html.Child(html.SvgPath(html.AD("M12 18a2.5 2.5 0 0 0 0-5 2.5 2.5 0 0 1 0-5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
