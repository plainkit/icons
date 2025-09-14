package lucide

import x "github.com/bloxui/blox"

// ArrowBigRight creates a Arrow Big Right Lucide icon.
func ArrowBigRight(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-big-right", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 9a1 1 0 0 0 1-1V5.061a1 1 0 0 1 1.811-.75l6.836 6.836a1.207 1.207 0 0 1 0 1.707l-6.836 6.835a1 1 0 0 1-1.811-.75V16a1 1 0 0 0-1-1H5a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1z"))),
	)
	return x.Svg(svgArgs...)
}
