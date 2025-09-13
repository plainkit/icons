package lucide

import x "github.com/bloxui/blox"

// EyeOff creates a Eye Off Lucide icon.
func EyeOff(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-eye-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.733 5.076a10.744 10.744 0 0 1 11.205 6.575 1 1 0 0 1 0 .696 10.747 10.747 0 0 1-1.444 2.49"))),
		x.Child(x.Path(x.D("M14.084 14.158a3 3 0 0 1-4.242-4.242"))),
		x.Child(x.Path(x.D("M17.479 17.499a10.75 10.75 0 0 1-15.417-5.151 1 1 0 0 1 0-.696 10.75 10.75 0 0 1 4.446-5.143"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
	)
	return x.Svg(svgArgs...)
}
