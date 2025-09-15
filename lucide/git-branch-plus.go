package lucide

import x "github.com/plainkit/html"

// GitBranchPlus creates a Git Branch Plus Lucide icon.
func GitBranchPlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-git-branch-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 3v12"))),
		x.Child(x.Path(x.D("M18 9a3 3 0 1 0 0-6 3 3 0 0 0 0 6z"))),
		x.Child(x.Path(x.D("M6 21a3 3 0 1 0 0-6 3 3 0 0 0 0 6z"))),
		x.Child(x.Path(x.D("M15 6a9 9 0 0 0-9 9"))),
		x.Child(x.Path(x.D("M18 15v6"))),
		x.Child(x.Path(x.D("M21 18h-6"))),
	)
	return x.Svg(svgArgs...)
}
