package lucide

import x "github.com/bloxui/blox"

// HandGrab creates a Hand Grab Lucide icon.
func HandGrab(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-hand-grab", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M18 11.5V9a2 2 0 0 0-2-2a2 2 0 0 0-2 2v1.4"))),
		x.Child(x.Path(x.D("M14 10V8a2 2 0 0 0-2-2a2 2 0 0 0-2 2v2"))),
		x.Child(x.Path(x.D("M10 9.9V9a2 2 0 0 0-2-2a2 2 0 0 0-2 2v5"))),
		x.Child(x.Path(x.D("M6 14a2 2 0 0 0-2-2a2 2 0 0 0-2 2"))),
		x.Child(x.Path(x.D("M18 11a2 2 0 1 1 4 0v3a8 8 0 0 1-8 8h-4a8 8 0 0 1-8-8 2 2 0 1 1 4 0"))),
	)
	return x.Svg(svgArgs...)
}
