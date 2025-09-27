package lucide

import (
	html "github.com/plainkit/html"
)

// DraftingCompass creates a Drafting Compass Lucide icon.
func DraftingCompass(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-drafting-compass", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m12.99 6.74 1.93 3.44"))),
		html.Child(html.SvgPath(html.AD("M19.136 12a10 10 0 0 1-14.271 0"))),
		html.Child(html.SvgPath(html.AD("m21 21-2.16-3.84"))),
		html.Child(html.SvgPath(html.AD("m3 21 8.02-14.26"))),
		html.Child(html.SvgCircle(html.ACx("12"), html.ACy("5"), html.AR("2"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
