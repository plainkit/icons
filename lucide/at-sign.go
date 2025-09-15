package lucide

import x "github.com/plainkit/blox"

// AtSign creates a At Sign Lucide icon.
func AtSign(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-at-sign", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("4"))),
		x.Child(x.Path(x.D("M16 8v5a3 3 0 0 0 6 0v-1a10 10 0 1 0-4 8"))),
	)
	return x.Svg(svgArgs...)
}
