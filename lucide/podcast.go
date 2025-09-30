package lucide

import (
	html "github.com/plainkit/html"
)

// Podcast creates a Podcast Lucide icon.
func Podcast(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-podcast", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13 17a1 1 0 1 0-2 0l.5 4.5a0.5 0.5 0 0 0 1 0z"), html.AFill("currentColor")),
		html.SvgPath(html.AD("M16.85 18.58a9 9 0 1 0-9.7 0")),
		html.SvgPath(html.AD("M8 14a5 5 0 1 1 8 0")),
		html.SvgCircle(html.ACx("12"), html.ACy("11"), html.AR("1"), html.AFill("currentColor")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
