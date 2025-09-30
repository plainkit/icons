package lucide

import (
	html "github.com/plainkit/html"
)

// KeyRound creates a Key Round Lucide icon.
func KeyRound(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-key-round", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2.586 17.414A2 2 0 0 0 2 18.828V21a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h1a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h.172a2 2 0 0 0 1.414-.586l.814-.814a6.5 6.5 0 1 0-4-4z")),
		html.SvgCircle(html.ACx("16.5"), html.ACy("7.5"), html.AR(".5"), html.AFill("currentColor")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
