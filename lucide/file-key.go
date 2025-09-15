package lucide

import x "github.com/plainkit/html"

// FileKey creates a File Key Lucide icon.
func FileKey(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-key", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("16"), x.R("2"))),
		x.Child(x.Path(x.D("m16 10-4.5 4.5"))),
		x.Child(x.Path(x.D("m15 11 1 1"))),
	)
	return x.Svg(svgArgs...)
}
