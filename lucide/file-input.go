package lucide

import x "github.com/plainkit/html"

// FileInput creates a File Input Lucide icon.
func FileInput(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-input", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M4 22h14a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v4"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M2 15h10"))),
		x.Child(x.Path(x.D("m9 18 3-3-3-3"))),
	)
	return x.Svg(svgArgs...)
}
