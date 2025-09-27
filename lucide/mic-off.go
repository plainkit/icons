package lucide

import (
	html "github.com/plainkit/html"
)

// MicOff creates a Mic Off Lucide icon.
func MicOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mic-off", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12 19v3"))),
		html.Child(html.SvgPath(html.AD("M15 9.34V5a3 3 0 0 0-5.68-1.33"))),
		html.Child(html.SvgPath(html.AD("M16.95 16.95A7 7 0 0 1 5 12v-2"))),
		html.Child(html.SvgPath(html.AD("M18.89 13.23A7 7 0 0 0 19 12v-2"))),
		html.Child(html.SvgPath(html.AD("m2 2 20 20"))),
		html.Child(html.SvgPath(html.AD("M9 9v3a3 3 0 0 0 5.12 2.12"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
