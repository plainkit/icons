package lucide

import (
	html "github.com/plainkit/html"
)

// MessageSquarePlus creates a Message Square Plus Lucide icon.
func MessageSquarePlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-message-square-plus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M22 17a2 2 0 0 1-2 2H6.828a2 2 0 0 0-1.414.586l-2.202 2.202A.71.71 0 0 1 2 21.286V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2z")),
		html.SvgPath(html.AD("M12 8v6")),
		html.SvgPath(html.AD("M9 11h6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
