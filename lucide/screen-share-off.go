package lucide

import (
	html "github.com/plainkit/html"
)

// ScreenShareOff creates a Screen Share Off Lucide icon.
func ScreenShareOff(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-screen-share-off", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M13 3H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-3")),
		html.SvgPath(html.AD("M8 21h8")),
		html.SvgPath(html.AD("M12 17v4")),
		html.SvgPath(html.AD("m22 3-5 5")),
		html.SvgPath(html.AD("m17 3 5 5")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
