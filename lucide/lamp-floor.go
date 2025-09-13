package lucide

import x "github.com/bloxui/blox"

// LampFloor creates a Lamp Floor Lucide icon.
func LampFloor(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-lamp-floor", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 10v12"))),
		x.Child(x.Path(x.D("M17.929 7.629A1 1 0 0 1 17 9H7a1 1 0 0 1-.928-1.371l2-5A1 1 0 0 1 9 2h6a1 1 0 0 1 .928.629z"))),
		x.Child(x.Path(x.D("M9 22h6"))),
	)
	return x.Svg(svgArgs...)
}
