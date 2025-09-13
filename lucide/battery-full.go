package lucide

import x "github.com/bloxui/blox"

// BatteryFull creates a Battery Full Lucide icon.
func BatteryFull(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-battery-full", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 10v4"))),
		x.Child(x.Path(x.D("M14 10v4"))),
		x.Child(x.Path(x.D("M22 14v-4"))),
		x.Child(x.Path(x.D("M6 10v4"))),
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("12"), x.X("2"), x.Y("6"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
