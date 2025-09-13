package lucide

import x "github.com/bloxui/blox"

// ServerCrash creates a Server Crash Lucide icon.
func ServerCrash(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-server-crash", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 10H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-2"))),
		x.Child(x.Path(x.D("M6 14H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-2"))),
		x.Child(x.Path(x.D("M6 6h.01"))),
		x.Child(x.Path(x.D("M6 18h.01"))),
		x.Child(x.Path(x.D("m13 6-4 6h6l-4 6"))),
	)
	return x.Svg(svgArgs...)
}
