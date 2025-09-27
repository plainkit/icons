package lucide

import (
	html "github.com/plainkit/html"
)

// CaptionsOff creates a Captions Off Lucide icon.
func CaptionsOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-captions-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M10.5 5H19a2 2 0 0 1 2 2v8.5"))),
		html.Child(html.SvgPath(html.AD("M17 11h-.5"))),
		html.Child(html.SvgPath(html.AD("M19 19H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M7 11h4"))),
		html.Child(html.SvgPath(html.AD("M7 15h2.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
