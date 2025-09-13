package lucide

import x "github.com/bloxui/blox"

// Group creates a Group Lucide icon.
func Group(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-group", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M3 7V5c0-1.1.9-2 2-2h2"))),
		x.Child(x.Path(x.D("M17 3h2c1.1 0 2 .9 2 2v2"))),
		x.Child(x.Path(x.D("M21 17v2c0 1.1-.9 2-2 2h-2"))),
		x.Child(x.Path(x.D("M7 21H5c-1.1 0-2-.9-2-2v-2"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("5"), x.X("7"), x.Y("7"), x.Rx("1"))),
		x.Child(x.Rect(x.RectWidth("7"), x.RectHeight("5"), x.X("10"), x.Y("12"), x.Rx("1"))),
	)
	return x.Svg(svgArgs...)
}
