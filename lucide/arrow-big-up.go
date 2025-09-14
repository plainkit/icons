package lucide

import x "github.com/bloxui/blox"

// ArrowBigUp creates a Arrow Big Up Lucide icon.
func ArrowBigUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-big-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 13a1 1 0 0 0-1-1H5.061a1 1 0 0 1-.75-1.811l6.836-6.835a1.207 1.207 0 0 1 1.707 0l6.835 6.835a1 1 0 0 1-.75 1.811H16a1 1 0 0 0-1 1v6a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1z"))),
	)
	return x.Svg(svgArgs...)
}
