package lucide

import x "github.com/plainkit/html"

// FastForward creates a Fast Forward Lucide icon.
func FastForward(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-fast-forward", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6a2 2 0 0 1 3.414-1.414l6 6a2 2 0 0 1 0 2.828l-6 6A2 2 0 0 1 12 18z"))),
		x.Child(x.Path(x.D("M2 6a2 2 0 0 1 3.414-1.414l6 6a2 2 0 0 1 0 2.828l-6 6A2 2 0 0 1 2 18z"))),
	)
	return x.Svg(svgArgs...)
}
