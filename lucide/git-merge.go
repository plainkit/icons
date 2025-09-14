package lucide

import x "github.com/bloxui/blox"

// GitMerge creates a Git Merge Lucide icon.
func GitMerge(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-git-merge", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("18"), x.Cy("18"), x.R("3"))),
		x.Child(x.Circle(x.Cx("6"), x.Cy("6"), x.R("3"))),
		x.Child(x.Path(x.D("M6 21V9a9 9 0 0 0 9 9"))),
	)
	return x.Svg(svgArgs...)
}
