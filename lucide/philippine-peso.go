package lucide

import x "github.com/bloxui/blox"

// PhilippinePeso creates a Philippine Peso Lucide icon.
func PhilippinePeso(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-philippine-peso", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M20 11H4"))),
		x.Child(x.Path(x.D("M20 7H4"))),
		x.Child(x.Path(x.D("M7 21V4a1 1 0 0 1 1-1h4a1 1 0 0 1 0 12H7"))),
	)
	return x.Svg(svgArgs...)
}
