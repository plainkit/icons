package lucide

import x "github.com/bloxui/blox"

// Pause creates a Pause Lucide icon.
func Pause(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-pause", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("5"), x.RectHeight("18"), x.X("14"), x.Y("3"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("5"), x.RectHeight("18"), x.X("5"), x.Y("3"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
