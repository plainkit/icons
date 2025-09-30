package lucide

import (
	html "github.com/plainkit/html"
)

// Mail creates a Mail Lucide icon.
func Mail(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mail", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m22 7-8.991 5.727a2 2 0 0 1-2.009 0L2 7")),
		html.SvgRect(html.AWidth("20"), html.AHeight("16"), html.AX("2"), html.AY("4"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
