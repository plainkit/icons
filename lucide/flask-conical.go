package lucide

import (
	html "github.com/plainkit/html"
)

// FlaskConical creates a Flask Conical Lucide icon.
func FlaskConical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-flask-conical", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14 2v6a2 2 0 0 0 .245.96l5.51 10.08A2 2 0 0 1 18 22H6a2 2 0 0 1-1.755-2.96l5.51-10.08A2 2 0 0 0 10 8V2"))),
		html.Child(html.SvgPath(html.AD("M6.453 15h11.094"))),
		html.Child(html.SvgPath(html.AD("M8.5 2h7"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
