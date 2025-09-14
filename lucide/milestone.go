package lucide

import x "github.com/bloxui/blox"

// Milestone creates a Milestone Lucide icon.
func Milestone(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-milestone", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 13v8"))),
		x.Child(x.Path(x.D("M12 3v3"))),
		x.Child(x.Path(x.D("M4 6a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h13a2 2 0 0 0 1.152-.365l3.424-2.317a1 1 0 0 0 0-1.635l-3.424-2.318A2 2 0 0 0 17 6z"))),
	)
	return x.Svg(svgArgs...)
}
