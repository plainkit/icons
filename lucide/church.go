package lucide

import x "github.com/plainkit/html"

// Church creates a Church Lucide icon.
func Church(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-church", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 9h4"))),
		x.Child(x.Path(x.D("M12 7v5"))),
		x.Child(x.Path(x.D("M14 21v-3a2 2 0 0 0-4 0v3"))),
		x.Child(x.Path(x.D("m18 9 3.52 2.147a1 1 0 0 1 .48.854V19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-6.999a1 1 0 0 1 .48-.854L6 9"))),
		x.Child(x.Path(x.D("M6 21V7a1 1 0 0 1 .376-.782l5-3.999a1 1 0 0 1 1.249.001l5 4A1 1 0 0 1 18 7v14"))),
	)
	return x.Svg(svgArgs...)
}
