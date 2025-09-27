package lucide

import (
	html "github.com/plainkit/html"
)

// CircleStop creates a Circle Stop Lucide icon.
func CircleStop(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-circle-stop", args)
	children := []html.SvgArg{
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("10"))),
		html.Child(html.SvgRect(html.AWidth("6"), html.AHeight("6"), html.AX("9"), html.AY("9"), html.ARx("1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
