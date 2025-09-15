package lucide

import x "github.com/plainkit/blox"

// FileText creates a File Text Lucide icon.
func FileText(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-text", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M10 9H8"))),
		x.Child(x.Path(x.D("M16 13H8"))),
		x.Child(x.Path(x.D("M16 17H8"))),
	)
	return x.Svg(svgArgs...)
}
