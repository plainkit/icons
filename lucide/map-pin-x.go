package lucide

import x "github.com/plainkit/blox"

// MapPinX creates a Map Pin X Lucide icon.
func MapPinX(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-map-pin-x", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M19.752 11.901A7.78 7.78 0 0 0 20 10a8 8 0 0 0-16 0c0 4.993 5.539 10.193 7.399 11.799a1 1 0 0 0 1.202 0 19 19 0 0 0 .09-.077"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("10"), x.R("3"))),
		x.Child(x.Path(x.D("m21.5 15.5-5 5"))),
		x.Child(x.Path(x.D("m21.5 20.5-5-5"))),
	)
	return x.Svg(svgArgs...)
}
