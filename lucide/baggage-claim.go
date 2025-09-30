package lucide

import (
	html "github.com/plainkit/html"
)

// BaggageClaim creates a Baggage Claim Lucide icon.
func BaggageClaim(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-baggage-claim", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M22 18H6a2 2 0 0 1-2-2V7a2 2 0 0 0-2-2")),
		html.SvgPath(html.AD("M17 14V4a2 2 0 0 0-2-2h-1a2 2 0 0 0-2 2v10")),
		html.SvgRect(html.AWidth("13"), html.AHeight("8"), html.AX("8"), html.AY("6"), html.ARx("1")),
		html.SvgCircle(html.ACx("18"), html.ACy("20"), html.AR("2")),
		html.SvgCircle(html.ACx("9"), html.ACy("20"), html.AR("2")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
