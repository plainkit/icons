package lucide

import x "github.com/plainkit/blox"

// TriangleDashed creates a Triangle Dashed Lucide icon.
func TriangleDashed(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-triangle-dashed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.17 4.193a2 2 0 0 1 3.666.013"))),
		x.Child(x.Path(x.D("M14 21h2"))),
		x.Child(x.Path(x.D("m15.874 7.743 1 1.732"))),
		x.Child(x.Path(x.D("m18.849 12.952 1 1.732"))),
		x.Child(x.Path(x.D("M21.824 18.18a2 2 0 0 1-1.835 2.824"))),
		x.Child(x.Path(x.D("M4.024 21a2 2 0 0 1-1.839-2.839"))),
		x.Child(x.Path(x.D("m5.136 12.952-1 1.732"))),
		x.Child(x.Path(x.D("M8 21h2"))),
		x.Child(x.Path(x.D("m8.102 7.743-1 1.732"))),
	)
	return x.Svg(svgArgs...)
}
