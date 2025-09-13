package lucide

import x "github.com/bloxui/blox"

// BatteryMedium creates a Battery Medium Lucide icon.
func BatteryMedium(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-battery-medium", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 14v-4"))),
		x.Child(x.Path(x.D("M22 14v-4"))),
		x.Child(x.Path(x.D("M6 14v-4"))),
		x.Child(x.Rect(x.X("2"), x.Y("6"), x.RectWidth("16"), x.RectHeight("12"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
