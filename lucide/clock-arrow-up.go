package lucide

import x "github.com/bloxui/blox"

// ClockArrowUp creates a Clock Arrow Up Lucide icon.
func ClockArrowUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clock-arrow-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v6l1.56.78"))),
		x.Child(x.Path(x.D("M13.227 21.925a10 10 0 1 1 8.767-9.588"))),
		x.Child(x.Path(x.D("m14 18 4-4 4 4"))),
		x.Child(x.Path(x.D("M18 22v-8"))),
	)
	return x.Svg(svgArgs...)
}
