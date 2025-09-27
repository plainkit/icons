package lucide

import (
	html "github.com/plainkit/html"
)

// FunnelPlus creates a Funnel Plus Lucide icon.
func FunnelPlus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-funnel-plus", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M13.354 3H3a1 1 0 0 0-.742 1.67l7.225 7.989A2 2 0 0 1 10 14v6a1 1 0 0 0 .553.895l2 1A1 1 0 0 0 14 21v-7a2 2 0 0 1 .517-1.341l1.218-1.348"))),
		html.Child(html.SvgPath(html.AD("M16 6h6"))),
		html.Child(html.SvgPath(html.AD("M19 3v6"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
