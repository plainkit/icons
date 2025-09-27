package lucide

import (
	html "github.com/plainkit/html"
)

// ScreenShare creates a Screen Share Lucide icon.
func ScreenShare(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-screen-share", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M13 3H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-3"))),
		html.Child(html.SvgPath(html.AD("M8 21h8"))),
		html.Child(html.SvgPath(html.AD("M12 17v4"))),
		html.Child(html.SvgPath(html.AD("m17 8 5-5"))),
		html.Child(html.SvgPath(html.AD("M17 3h5v5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
