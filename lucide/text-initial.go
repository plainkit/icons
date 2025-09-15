package lucide

import x "github.com/plainkit/blox"

// TextInitial creates a Text Initial Lucide icon.
func TextInitial(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-text-initial", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 5h6"))),
		x.Child(x.Path(x.D("M15 12h6"))),
		x.Child(x.Path(x.D("M3 19h18"))),
		x.Child(x.Path(x.D("m3 12 3.553-7.724a.5.5 0 0 1 .894 0L11 12"))),
		x.Child(x.Path(x.D("M3.92 10h6.16"))),
	)
	return x.Svg(svgArgs...)
}
