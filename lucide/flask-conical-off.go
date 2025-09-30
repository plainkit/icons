package lucide

import (
	html "github.com/plainkit/html"
)

// FlaskConicalOff creates a Flask Conical Off Lucide icon.
func FlaskConicalOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-flask-conical-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 2v2.343")),
		html.SvgPath(html.AD("M14 2v6.343")),
		html.SvgPath(html.AD("m2 2 20 20")),
		html.SvgPath(html.AD("M20 20a2 2 0 0 1-2 2H6a2 2 0 0 1-1.755-2.96l5.227-9.563")),
		html.SvgPath(html.AD("M6.453 15H15")),
		html.SvgPath(html.AD("M8.5 2h7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
