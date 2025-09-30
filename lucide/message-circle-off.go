package lucide

import (
	html "github.com/plainkit/html"
)

// MessageCircleOff creates a Message Circle Off Lucide icon.
func MessageCircleOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-circle-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("m2 2 20 20")),
		html.SvgPath(html.AD("M4.93 4.929a10 10 0 0 0-1.938 11.412 2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 0 0 11.302-1.989")),
		html.SvgPath(html.AD("M8.35 2.69A10 10 0 0 1 21.3 15.65")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
