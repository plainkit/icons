package lucide

import x "github.com/bloxui/blox"

// Baby creates a Baby Lucide icon.
func Baby(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-baby", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 16c.5.3 1.2.5 2 .5s1.5-.2 2-.5"))),
		x.Child(x.Path(x.D("M15 12h.01"))),
		x.Child(x.Path(x.D("M19.38 6.813A9 9 0 0 1 20.8 10.2a2 2 0 0 1 0 3.6 9 9 0 0 1-17.6 0 2 2 0 0 1 0-3.6A9 9 0 0 1 12 3c2 0 3.5 1.1 3.5 2.5s-.9 2.5-2 2.5c-.8 0-1.5-.4-1.5-1"))),
		x.Child(x.Path(x.D("M9 12h.01"))),
	)
	return x.Svg(svgArgs...)
}
