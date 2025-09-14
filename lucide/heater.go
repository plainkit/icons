package lucide

import x "github.com/bloxui/blox"

// Heater creates a Heater Lucide icon.
func Heater(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-heater", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 8c2-3-2-3 0-6"))),
		x.Child(x.Path(x.D("M15.5 8c2-3-2-3 0-6"))),
		x.Child(x.Path(x.D("M6 10h.01"))),
		x.Child(x.Path(x.D("M6 14h.01"))),
		x.Child(x.Path(x.D("M10 16v-4"))),
		x.Child(x.Path(x.D("M14 16v-4"))),
		x.Child(x.Path(x.D("M18 16v-4"))),
		x.Child(x.Path(x.D("M20 6a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h3"))),
		x.Child(x.Path(x.D("M5 20v2"))),
		x.Child(x.Path(x.D("M19 20v2"))),
	)
	return x.Svg(svgArgs...)
}
