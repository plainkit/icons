package lucide

import x "github.com/plainkit/html"

// FileCode creates a File Code Lucide icon.
func FileCode(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-code", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 12.5 8 15l2 2.5"))),
		x.Child(x.Path(x.D("m14 12.5 2 2.5-2 2.5"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7z"))),
	)
	return x.Svg(svgArgs...)
}
