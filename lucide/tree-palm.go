package lucide

import x "github.com/bloxui/blox"

// TreePalm creates a Tree Palm Lucide icon.
func TreePalm(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tree-palm", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13 8c0-2.76-2.46-5-5.5-5S2 5.24 2 8h2l1-1 1 1h4"))),
		x.Child(x.Path(x.D("M13 7.14A5.82 5.82 0 0 1 16.5 6c3.04 0 5.5 2.24 5.5 5h-3l-1-1-1 1h-3"))),
		x.Child(x.Path(x.D("M5.89 9.71c-2.15 2.15-2.3 5.47-.35 7.43l4.24-4.25.7-.7.71-.71 2.12-2.12c-1.95-1.96-5.27-1.8-7.42.35"))),
		x.Child(x.Path(x.D("M11 15.5c.5 2.5-.17 4.5-1 6.5h4c2-5.5-.5-12-1-14"))),
	)
	return x.Svg(svgArgs...)
}
