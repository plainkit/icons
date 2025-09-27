package lucide

import (
	html "github.com/plainkit/html"
)

// FunnelX creates a Funnel X Lucide icon.
func FunnelX(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-funnel-x", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M12.531 3H3a1 1 0 0 0-.742 1.67l7.225 7.989A2 2 0 0 1 10 14v6a1 1 0 0 0 .553.895l2 1A1 1 0 0 0 14 21v-7a2 2 0 0 1 .517-1.341l.427-.473"))),
		html.Child(html.SvgPath(html.AD("m16.5 3.5 5 5"))),
		html.Child(html.SvgPath(html.AD("m21.5 3.5-5 5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
