package lucide

import x "github.com/plainkit/html"

// Goal creates a Goal Lucide icon.
func Goal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-goal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 13V2l8 4-8 4"))),
		x.Child(x.Path(x.D("M20.561 10.222a9 9 0 1 1-12.55-5.29"))),
		x.Child(x.Path(x.D("M8.002 9.997a5 5 0 1 0 8.9 2.02"))),
	)
	return x.Svg(svgArgs...)
}
