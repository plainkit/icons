package lucide

import x "github.com/bloxui/blox"

// Search creates a Search Lucide icon.
func Search(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-search", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m21 21-4.34-4.34"))),
		x.Child(x.Circle(x.Cx("11"), x.Cy("11"), x.R("8"))),
	)
	return x.Svg(svgArgs...)
}
