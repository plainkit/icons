package lucide

import x "github.com/plainkit/html"

// MicOff creates a Mic Off Lucide icon.
func MicOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-mic-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 19v3"))),
		x.Child(x.Path(x.D("M15 9.34V5a3 3 0 0 0-5.68-1.33"))),
		x.Child(x.Path(x.D("M16.95 16.95A7 7 0 0 1 5 12v-2"))),
		x.Child(x.Path(x.D("M18.89 13.23A7 7 0 0 0 19 12v-2"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M9 9v3a3 3 0 0 0 5.12 2.12"))),
	)
	return x.Svg(svgArgs...)
}
