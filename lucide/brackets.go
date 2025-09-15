package lucide

import x "github.com/plainkit/blox"

// Brackets creates a Brackets Lucide icon.
func Brackets(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-brackets", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 3h3a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-3"))),
		x.Child(x.Path(x.D("M8 21H5a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h3"))),
	)
	return x.Svg(svgArgs...)
}
