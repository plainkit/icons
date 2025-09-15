package lucide

import x "github.com/plainkit/html"

// Loader creates a Loader Lucide icon.
func Loader(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-loader", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2v4"))),
		x.Child(x.Path(x.D("m16.2 7.8 2.9-2.9"))),
		x.Child(x.Path(x.D("M18 12h4"))),
		x.Child(x.Path(x.D("m16.2 16.2 2.9 2.9"))),
		x.Child(x.Path(x.D("M12 18v4"))),
		x.Child(x.Path(x.D("m4.9 19.1 2.9-2.9"))),
		x.Child(x.Path(x.D("M2 12h4"))),
		x.Child(x.Path(x.D("m4.9 4.9 2.9 2.9"))),
	)
	return x.Svg(svgArgs...)
}
