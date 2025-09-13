package lucide

import x "github.com/bloxui/blox"

// ServerOff creates a Server Off Lucide icon.
func ServerOff(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-server-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M7 2h13a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-5"))),
		x.Child(x.Path(x.D("M10 10 2.5 2.5C2 2 2 2.5 2 5v3a2 2 0 0 0 2 2h6z"))),
		x.Child(x.Path(x.D("M22 17v-1a2 2 0 0 0-2-2h-1"))),
		x.Child(x.Path(x.D("M4 14a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16.5l1-.5.5.5-8-8H4z"))),
		x.Child(x.Path(x.D("M6 18h.01"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
	)
	return x.Svg(svgArgs...)
}
