package lucide

import x "github.com/plainkit/html"

// MouseOff creates a Mouse Off Lucide icon.
func MouseOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-mouse-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 6v.343"))),
		x.Child(x.Path(x.D("M18.218 18.218A7 7 0 0 1 5 15V9a7 7 0 0 1 .782-3.218"))),
		x.Child(x.Path(x.D("M19 13.343V9A7 7 0 0 0 8.56 2.902"))),
		x.Child(x.Path(x.D("M22 22 2 2"))),
	)
	return x.Svg(svgArgs...)
}
