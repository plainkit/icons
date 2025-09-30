package lucide

import (
	html "github.com/plainkit/html"
)

// MousePointerBan creates a Mouse Pointer Ban Lucide icon.
func MousePointerBan(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-mouse-pointer-ban", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M2.034 2.681a.498.498 0 0 1 .647-.647l9 3.5a.5.5 0 0 1-.033.944L8.204 7.545a1 1 0 0 0-.66.66l-1.066 3.443a.5.5 0 0 1-.944.033z")),
		html.SvgCircle(html.ACx("16"), html.ACy("16"), html.AR("6")),
		html.SvgPath(html.AD("m11.8 11.8 8.4 8.4")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
