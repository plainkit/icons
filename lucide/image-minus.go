package lucide

import (
	html "github.com/plainkit/html"
)

// ImageMinus creates a Image Minus Lucide icon.
func ImageMinus(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-image-minus", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M21 9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7")),
		html.SvgLine(html.AX1("16"), html.AX2("22"), html.AY1("5"), html.AY2("5")),
		html.SvgCircle(html.ACx("9"), html.ACy("9"), html.AR("2")),
		html.SvgPath(html.AD("m21 15-3.086-3.086a2 2 0 0 0-2.828 0L6 21")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
