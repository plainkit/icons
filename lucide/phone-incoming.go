package lucide

import (
	html "github.com/plainkit/html"
)

// PhoneIncoming creates a Phone Incoming Lucide icon.
func PhoneIncoming(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-phone-incoming", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M16 2v6h6"))),
		html.Child(html.SvgPath(html.AD("m22 2-6 6"))),
		html.Child(html.SvgPath(html.AD("M13.832 16.568a1 1 0 0 0 1.213-.303l.355-.465A2 2 0 0 1 17 15h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2A18 18 0 0 1 2 4a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v3a2 2 0 0 1-.8 1.6l-.468.351a1 1 0 0 0-.292 1.233 14 14 0 0 0 6.392 6.384"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
