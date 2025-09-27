package lucide

import (
	html "github.com/plainkit/html"
)

// Luggage creates a Luggage Lucide icon.
func Luggage(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-luggage", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M6 20a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2"))),
		html.Child(html.SvgPath(html.AD("M8 18V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v14"))),
		html.Child(html.SvgPath(html.AD("M10 20h4"))),
		html.Child(html.SvgCircle(html.ACx("16"), html.ACy("20"), html.AR("2"))),
		html.Child(html.SvgCircle(html.ACx("8"), html.ACy("20"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
