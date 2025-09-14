package lucide

import x "github.com/bloxui/blox"

// Factory creates a Factory Lucide icon.
func Factory(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-factory", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 16h.01"))),
		x.Child(x.Path(x.D("M16 16h.01"))),
		x.Child(x.Path(x.D("M3 19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8.5a.5.5 0 0 0-.769-.422l-4.462 2.844A.5.5 0 0 1 15 10.5v-2a.5.5 0 0 0-.769-.422L9.77 10.922A.5.5 0 0 1 9 10.5V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2z"))),
		x.Child(x.Path(x.D("M8 16h.01"))),
	)
	return x.Svg(svgArgs...)
}
