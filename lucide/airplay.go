package lucide

import (
	html "github.com/plainkit/html"
)

// Airplay creates a Airplay Lucide icon.
func Airplay(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-airplay", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M5 17H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1")),
		html.SvgPath(html.AD("m12 15 5 6H7Z")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
