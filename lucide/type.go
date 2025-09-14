package lucide

import x "github.com/bloxui/blox"

// Type creates a Type Lucide icon.
func Type(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-type", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 4v16"))),
		x.Child(x.Path(x.D("M4 7V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v2"))),
		x.Child(x.Path(x.D("M9 20h6"))),
	)
	return x.Svg(svgArgs...)
}
