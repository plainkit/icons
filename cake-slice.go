package lucide

import x "github.com/bloxui/blox"

// CakeSlice creates a Cake Slice Lucide icon.
func CakeSlice(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-cake-slice", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 13H3"))),
		x.Child(x.Path(x.D("M16 17H3"))),
		x.Child(x.Path(x.D("m7.2 7.9-3.388 2.5A2 2 0 0 0 3 12.01V20a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-8.654c0-2-2.44-6.026-6.44-8.026a1 1 0 0 0-1.082.057L10.4 5.6"))),
		x.Child(x.Circle(x.Cx("9"), x.Cy("7"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
