package lucide

import (
	html "github.com/plainkit/html"
)

// WavesLadder creates a Waves Ladder Lucide icon.
func WavesLadder(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-waves-ladder", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M19 5a2 2 0 0 0-2 2v11")),
		html.SvgPath(html.AD("M2 18c.6.5 1.2 1 2.5 1 2.5 0 2.5-2 5-2 2.6 0 2.4 2 5 2 2.5 0 2.5-2 5-2 1.3 0 1.9.5 2.5 1")),
		html.SvgPath(html.AD("M7 13h10")),
		html.SvgPath(html.AD("M7 9h10")),
		html.SvgPath(html.AD("M9 5a2 2 0 0 0-2 2v11")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
