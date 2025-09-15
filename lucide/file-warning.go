package lucide

import x "github.com/plainkit/html"

// FileWarning creates a File Warning Lucide icon.
func FileWarning(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-warning", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Path(x.D("M12 9v4"))),
		x.Child(x.Path(x.D("M12 17h.01"))),
	)
	return x.Svg(svgArgs...)
}
