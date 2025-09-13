package lucide

import x "github.com/bloxui/blox"

// Recycle creates a Recycle Lucide icon.
func Recycle(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-recycle", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M7 19H4.815a1.83 1.83 0 0 1-1.57-.881 1.785 1.785 0 0 1-.004-1.784L7.196 9.5"))),
		x.Child(x.Path(x.D("M11 19h8.203a1.83 1.83 0 0 0 1.556-.89 1.784 1.784 0 0 0 0-1.775l-1.226-2.12"))),
		x.Child(x.Path(x.D("m14 16-3 3 3 3"))),
		x.Child(x.Path(x.D("M8.293 13.596 7.196 9.5 3.1 10.598"))),
		x.Child(x.Path(x.D("m9.344 5.811 1.093-1.892A1.83 1.83 0 0 1 11.985 3a1.784 1.784 0 0 1 1.546.888l3.943 6.843"))),
		x.Child(x.Path(x.D("m13.378 9.633 4.096 1.098 1.097-4.096"))),
	)
	return x.Svg(svgArgs...)
}
