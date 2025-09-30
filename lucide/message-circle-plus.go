package lucide

import (
	html "github.com/plainkit/html"
)

// MessageCirclePlus creates a Message Circle Plus Lucide icon.
func MessageCirclePlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-circle-plus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2.992 16.342a2 2 0 0 1 .094 1.167l-1.065 3.29a1 1 0 0 0 1.236 1.168l3.413-.998a2 2 0 0 1 1.099.092 10 10 0 1 0-4.777-4.719")),
		html.SvgPath(html.AD("M8 12h8")),
		html.SvgPath(html.AD("M12 8v8")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
