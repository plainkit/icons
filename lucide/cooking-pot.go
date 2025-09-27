package lucide

import (
	html "github.com/plainkit/html"
)

// CookingPot creates a Cooking Pot Lucide icon.
func CookingPot(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-cooking-pot", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M2 12h20"))),
		html.Child(html.SvgPath(html.AD("M20 12v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8"))),
		html.Child(html.SvgPath(html.AD("m4 8 16-4"))),
		html.Child(html.SvgPath(html.AD("m8.86 6.78-.45-1.81a2 2 0 0 1 1.45-2.43l1.94-.48a2 2 0 0 1 2.43 1.46l.45 1.8"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
