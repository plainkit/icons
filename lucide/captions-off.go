package lucide

import x "github.com/bloxui/blox"

// CaptionsOff creates a Captions Off Lucide icon.
func CaptionsOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-captions-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.5 5H19a2 2 0 0 1 2 2v8.5"))),
		x.Child(x.Path(x.D("M17 11h-.5"))),
		x.Child(x.Path(x.D("M19 19H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M7 11h4"))),
		x.Child(x.Path(x.D("M7 15h2.5"))),
	)
	return x.Svg(svgArgs...)
}
