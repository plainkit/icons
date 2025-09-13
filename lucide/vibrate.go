package lucide

import x "github.com/bloxui/blox"

// Vibrate creates a Vibrate Lucide icon.
func Vibrate(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-vibrate", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m2 8 2 2-2 2 2 2-2 2"))),
		x.Child(x.Path(x.D("m22 8-2 2 2 2-2 2 2 2"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("14"), x.X("8"), x.Y("5"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
