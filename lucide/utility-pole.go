package lucide

import x "github.com/bloxui/blox"

// UtilityPole creates a Utility Pole Lucide icon.
func UtilityPole(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-utility-pole", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2v20"))),
		x.Child(x.Path(x.D("M2 5h20"))),
		x.Child(x.Path(x.D("M3 3v2"))),
		x.Child(x.Path(x.D("M7 3v2"))),
		x.Child(x.Path(x.D("M17 3v2"))),
		x.Child(x.Path(x.D("M21 3v2"))),
		x.Child(x.Path(x.D("m19 5-7 7-7-7"))),
	)
	return x.Svg(svgArgs...)
}
