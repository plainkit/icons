package lucide

import x "github.com/plainkit/blox"

// BookOpen creates a Book Open Lucide icon.
func BookOpen(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-book-open", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 7v14"))),
		x.Child(x.Path(x.D("M3 18a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h5a4 4 0 0 1 4 4 4 4 0 0 1 4-4h5a1 1 0 0 1 1 1v13a1 1 0 0 1-1 1h-6a3 3 0 0 0-3 3 3 3 0 0 0-3-3z"))),
	)
	return x.Svg(svgArgs...)
}
