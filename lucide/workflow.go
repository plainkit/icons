package lucide

import (
	html "github.com/plainkit/html"
)

// Workflow creates a Workflow Lucide icon.
func Workflow(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-workflow", args)
	children := []html.SvgArg{
		html.SvgRect(html.AWidth("8"), html.AHeight("8"), html.AX("3"), html.AY("3"), html.ARx("2")),
		html.SvgPath(html.AD("M7 11v4a2 2 0 0 0 2 2h4")),
		html.SvgRect(html.AWidth("8"), html.AHeight("8"), html.AX("13"), html.AY("13"), html.ARx("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
