package lucide

import x "github.com/plainkit/html"

// Cookie creates a Cookie Lucide icon.
func Cookie(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cookie", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2a10 10 0 1 0 10 10 4 4 0 0 1-5-5 4 4 0 0 1-5-5"))),
		x.Child(x.Path(x.D("M8.5 8.5v.01"))),
		x.Child(x.Path(x.D("M16 15.5v.01"))),
		x.Child(x.Path(x.D("M12 12v.01"))),
		x.Child(x.Path(x.D("M11 17v.01"))),
		x.Child(x.Path(x.D("M7 14v.01"))),
	)
	return x.Svg(svgArgs...)
}
