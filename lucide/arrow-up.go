package lucide

import x "github.com/plainkit/html"

// ArrowUp creates a Arrow Up Lucide icon.
func ArrowUp(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-arrow-up", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m5 12 7-7 7 7"))),
		x.Child(x.Path(x.D("M12 19V5"))),
	)
	return x.Svg(svgArgs...)
}
