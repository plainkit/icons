package lucide

import x "github.com/bloxui/blox"

// Combine creates a Combine Lucide icon.
func Combine(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-combine", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 18H5a3 3 0 0 1-3-3v-1"))),
		x.Child(x.Path(x.D("M14 2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("M20 2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2"))),
		x.Child(x.Path(x.D("m7 21 3-3-3-3"))),
		x.Child(x.Rect(x.X("14"), x.Y("14"), x.RectWidth("8"), x.RectHeight("8"), x.Rx("2"))),
		x.Child(x.Rect(x.X("2"), x.Y("2"), x.RectWidth("8"), x.RectHeight("8"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}