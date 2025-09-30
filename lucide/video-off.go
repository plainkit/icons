package lucide

import (
	html "github.com/plainkit/html"
)

// VideoOff creates a Video Off Lucide icon.
func VideoOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-video-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M10.66 6H14a2 2 0 0 1 2 2v2.5l5.248-3.062A.5.5 0 0 1 22 7.87v8.196")),
		html.SvgPath(html.AD("M16 16a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h2")),
		html.SvgPath(html.AD("m2 2 20 20")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
