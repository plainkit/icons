package lucide

import x "github.com/plainkit/blox"

// SunMedium creates a Sun Medium Lucide icon.
func SunMedium(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-sun-medium", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("12"), x.R("4"))),
		x.Child(x.Path(x.D("M12 3v1"))),
		x.Child(x.Path(x.D("M12 20v1"))),
		x.Child(x.Path(x.D("M3 12h1"))),
		x.Child(x.Path(x.D("M20 12h1"))),
		x.Child(x.Path(x.D("m18.364 5.636-.707.707"))),
		x.Child(x.Path(x.D("m6.343 17.657-.707.707"))),
		x.Child(x.Path(x.D("m5.636 5.636.707.707"))),
		x.Child(x.Path(x.D("m17.657 17.657.707.707"))),
	)
	return x.Svg(svgArgs...)
}
