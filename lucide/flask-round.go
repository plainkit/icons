package lucide

import (
	html "github.com/plainkit/html"
)

// FlaskRound creates a Flask Round Lucide icon.
func FlaskRound(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-flask-round", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10 2v6.292a7 7 0 1 0 4 0V2")),
		html.SvgPath(html.AD("M5 15h14")),
		html.SvgPath(html.AD("M8.5 2h7")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
