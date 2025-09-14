package lucide

import x "github.com/bloxui/blox"

// Ungroup creates a Ungroup Lucide icon.
func Ungroup(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ungroup", args)
	svgArgs = append(svgArgs,
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("6"), x.X("5"), x.Y("4"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("8"), x.RectHeight("6"), x.X("11"), x.Y("14"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
