package lucide

import x "github.com/plainkit/blox"

// School creates a School Lucide icon.
func School(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-school", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 21v-3a2 2 0 0 0-4 0v3"))),
		x.Child(x.Path(x.D("M18 5v16"))),
		x.Child(x.Path(x.D("m4 6 7.106-3.79a2 2 0 0 1 1.788 0L20 6"))),
		x.Child(x.Path(x.D("m6 11-3.52 2.147a1 1 0 0 0-.48.854V19a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-5a1 1 0 0 0-.48-.853L18 11"))),
		x.Child(x.Path(x.D("M6 5v16"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("9"), x.R("2"))),
	)
	return x.Svg(svgArgs...)
}
