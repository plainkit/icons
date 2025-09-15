package lucide

import x "github.com/plainkit/blox"

// MapPinCheck creates a Map Pin Check Lucide icon.
func MapPinCheck(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-map-pin-check", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M19.43 12.935c.357-.967.57-1.955.57-2.935a8 8 0 0 0-16 0c0 4.993 5.539 10.193 7.399 11.799a1 1 0 0 0 1.202 0 32.197 32.197 0 0 0 .813-.728"))),
		x.Child(x.Circle(x.Cx("12"), x.Cy("10"), x.R("3"))),
		x.Child(x.Path(x.D("m16 18 2 2 4-4"))),
	)
	return x.Svg(svgArgs...)
}
