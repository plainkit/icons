package lucide

import x "github.com/plainkit/html"

// TrainFront creates a Train Front Lucide icon.
func TrainFront(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-train-front", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 3.1V7a4 4 0 0 0 8 0V3.1"))),
		x.Child(x.Path(x.D("m9 15-1-1"))),
		x.Child(x.Path(x.D("m15 15 1-1"))),
		x.Child(x.Path(x.D("M9 19c-2.8 0-5-2.2-5-5v-4a8 8 0 0 1 16 0v4c0 2.8-2.2 5-5 5Z"))),
		x.Child(x.Path(x.D("m8 19-2 3"))),
		x.Child(x.Path(x.D("m16 19 2 3"))),
	)
	return x.Svg(svgArgs...)
}
