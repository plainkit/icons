package lucide

import x "github.com/plainkit/blox"

// FileCheck creates a File Check Lucide icon.
func FileCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("m9 15 2 2 4-4"))),
	)
	return x.Svg(svgArgs...)
}
