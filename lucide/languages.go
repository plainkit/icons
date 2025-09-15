package lucide

import x "github.com/plainkit/blox"

// Languages creates a Languages Lucide icon.
func Languages(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-languages", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m5 8 6 6"))),
		x.Child(x.Path(x.D("m4 14 6-6 2-3"))),
		x.Child(x.Path(x.D("M2 5h12"))),
		x.Child(x.Path(x.D("M7 2h1"))),
		x.Child(x.Path(x.D("m22 22-5-10-5 10"))),
		x.Child(x.Path(x.D("M14 18h6"))),
	)
	return x.Svg(svgArgs...)
}
