package lucide

import x "github.com/plainkit/html"

// FileSearch2 creates a File Search 2 Lucide icon.
func FileSearch2(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-search-2", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Circle(x.Cx("11.5"), x.Cy("14.5"), x.R("2.5"))),
		x.Child(x.Path(x.D("M13.3 16.3 15 18"))),
	)
	return x.Svg(svgArgs...)
}
