package lucide

import (
	html "github.com/plainkit/html"
)

// ListTodo creates a List Todo Lucide icon.
func ListTodo(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-list-todo", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13 5h8")),
		html.SvgPath(html.AD("M13 12h8")),
		html.SvgPath(html.AD("M13 19h8")),
		html.SvgPath(html.AD("m3 17 2 2 4-4")),
		html.SvgRect(html.AWidth("6"), html.AHeight("6"), html.AX("3"), html.AY("4"), html.ARx("1")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
