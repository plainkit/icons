package lucide

import x "github.com/bloxui/blox"

// Minimize creates a Minimize Lucide icon.
func Minimize(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-minimize", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 3v3a2 2 0 0 1-2 2H3"))),
		x.Child(x.Path(x.D("M21 8h-3a2 2 0 0 1-2-2V3"))),
		x.Child(x.Path(x.D("M3 16h3a2 2 0 0 1 2 2v3"))),
		x.Child(x.Path(x.D("M16 21v-3a2 2 0 0 1 2-2h3"))),
	)
	return x.Svg(svgArgs...)
}
