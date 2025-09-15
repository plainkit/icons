package lucide

import x "github.com/plainkit/blox"

// FileOutput creates a File Output Lucide icon.
func FileOutput(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-output", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M4 7V4a2 2 0 0 1 2-2 2 2 0 0 0-2 2"))),
		x.Child(x.Path(x.D("M4.063 20.999a2 2 0 0 0 2 1L18 22a2 2 0 0 0 2-2V7l-5-5H6"))),
		x.Child(x.Path(x.D("m5 11-3 3"))),
		x.Child(x.Path(x.D("m5 17-3-3h10"))),
	)
	return x.Svg(svgArgs...)
}
