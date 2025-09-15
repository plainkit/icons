package lucide

import x "github.com/plainkit/blox"

// BatteryLow creates a Battery Low Lucide icon.
func BatteryLow(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-battery-low", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 14v-4"))),
		x.Child(x.Path(x.D("M6 14v-4"))),
		x.Child(x.Rect(x.RectWidth("16"), x.RectHeight("12"), x.X("2"), x.Y("6"), x.Rx("2"))),
	)
	return x.Svg(svgArgs...)
}
