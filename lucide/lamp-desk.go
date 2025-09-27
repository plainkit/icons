package lucide

import (
	html "github.com/plainkit/html"
)

// LampDesk creates a Lamp Desk Lucide icon.
func LampDesk(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-lamp-desk", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10.293 2.293a1 1 0 0 1 1.414 0l2.5 2.5 5.994 1.227a1 1 0 0 1 .506 1.687l-7 7a1 1 0 0 1-1.687-.506l-1.227-5.994-2.5-2.5a1 1 0 0 1 0-1.414z"))),
		html.Child(html.SvgPath(html.AD("m14.207 4.793-3.414 3.414"))),
		html.Child(html.SvgPath(html.AD("M3 20a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"))),
		html.Child(html.SvgPath(html.AD("m9.086 6.5-4.793 4.793a1 1 0 0 0-.18 1.17L7 18"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
