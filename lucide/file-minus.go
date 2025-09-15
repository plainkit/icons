package lucide

import x "github.com/plainkit/html"

// FileMinus creates a File Minus Lucide icon.
func FileMinus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-minus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M9 15h6"))),
	)
	return x.Svg(svgArgs...)
}
