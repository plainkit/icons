package lucide

import x "github.com/bloxui/blox"

// ClockFading creates a Clock Fading Lucide icon.
func ClockFading(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-clock-fading", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2a10 10 0 0 1 7.38 16.75"))),
		x.Child(x.Path(x.D("M12 6v6l4 2"))),
		x.Child(x.Path(x.D("M2.5 8.875a10 10 0 0 0-.5 3"))),
		x.Child(x.Path(x.D("M2.83 16a10 10 0 0 0 2.43 3.4"))),
		x.Child(x.Path(x.D("M4.636 5.235a10 10 0 0 1 .891-.857"))),
		x.Child(x.Path(x.D("M8.644 21.42a10 10 0 0 0 7.631-.38"))),
	)
	return x.Svg(svgArgs...)
}
