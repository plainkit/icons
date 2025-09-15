package lucide

import x "github.com/plainkit/html"

// Ghost creates a Ghost Lucide icon.
func Ghost(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-ghost", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 10h.01"))),
		x.Child(x.Path(x.D("M15 10h.01"))),
		x.Child(x.Path(x.D("M12 2a8 8 0 0 0-8 8v12l3-3 2.5 2.5L12 19l2.5 2.5L17 19l3 3V10a8 8 0 0 0-8-8z"))),
	)
	return x.Svg(svgArgs...)
}
