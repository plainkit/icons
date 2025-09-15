package lucide

import x "github.com/plainkit/html"

// Rewind creates a Rewind Lucide icon.
func Rewind(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-rewind", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6a2 2 0 0 0-3.414-1.414l-6 6a2 2 0 0 0 0 2.828l6 6A2 2 0 0 0 12 18z"))),
		x.Child(x.Path(x.D("M22 6a2 2 0 0 0-3.414-1.414l-6 6a2 2 0 0 0 0 2.828l6 6A2 2 0 0 0 22 18z"))),
	)
	return x.Svg(svgArgs...)
}
