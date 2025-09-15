package lucide

import x "github.com/plainkit/blox"

// FileDown creates a File Down Lucide icon.
func FileDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M12 18v-6"))),
		x.Child(x.Path(x.D("m9 15 3 3 3-3"))),
	)
	return x.Svg(svgArgs...)
}
