package lucide

import x "github.com/bloxui/blox"

// ArrowBigDown creates a Arrow Big Down Lucide icon.
func ArrowBigDown(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-big-down", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 11a1 1 0 0 0 1 1h2.939a1 1 0 0 1 .75 1.811l-6.835 6.836a1.207 1.207 0 0 1-1.707 0L4.31 13.81a1 1 0 0 1 .75-1.811H8a1 1 0 0 0 1-1V5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1z"))),
	)
	return x.Svg(svgArgs...)
}
