package lucide

import x "github.com/bloxui/blox"

// Theater creates a Theater Lucide icon.
func Theater(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-theater", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M2 10s3-3 3-8"))),
		x.Child(x.Path(x.D("M22 10s-3-3-3-8"))),
		x.Child(x.Path(x.D("M10 2c0 4.4-3.6 8-8 8"))),
		x.Child(x.Path(x.D("M14 2c0 4.4 3.6 8 8 8"))),
		x.Child(x.Path(x.D("M2 10s2 2 2 5"))),
		x.Child(x.Path(x.D("M22 10s-2 2-2 5"))),
		x.Child(x.Path(x.D("M8 15h8"))),
		x.Child(x.Path(x.D("M2 22v-1a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1"))),
		x.Child(x.Path(x.D("M14 22v-1a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1"))),
	)
	return x.Svg(svgArgs...)
}
