package lucide

import x "github.com/plainkit/html"

// Scale creates a Scale Lucide icon.
func Scale(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-scale", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m16 16 3-8 3 8c-.87.65-1.92 1-3 1s-2.13-.35-3-1Z"))),
		x.Child(x.Path(x.D("m2 16 3-8 3 8c-.87.65-1.92 1-3 1s-2.13-.35-3-1Z"))),
		x.Child(x.Path(x.D("M7 21h10"))),
		x.Child(x.Path(x.D("M12 3v18"))),
		x.Child(x.Path(x.D("M3 7h2c2 0 5-1 7-2 2 1 5 2 7 2h2"))),
	)
	return x.Svg(svgArgs...)
}
