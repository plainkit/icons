package lucide

import (
	html "github.com/plainkit/html"
)

// PhoneForwarded creates a Phone Forwarded Lucide icon.
func PhoneForwarded(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-phone-forwarded", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M14 6h8")),
		html.SvgPath(html.AD("m18 2 4 4-4 4")),
		html.SvgPath(html.AD("M13.832 16.568a1 1 0 0 0 1.213-.303l.355-.465A2 2 0 0 1 17 15h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2A18 18 0 0 1 2 4a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v3a2 2 0 0 1-.8 1.6l-.468.351a1 1 0 0 0-.292 1.233 14 14 0 0 0 6.392 6.384")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
