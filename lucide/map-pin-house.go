package lucide

import x "github.com/plainkit/blox"

// MapPinHouse creates a Map Pin House Lucide icon.
func MapPinHouse(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-map-pin-house", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M15 22a1 1 0 0 1-1-1v-4a1 1 0 0 1 .445-.832l3-2a1 1 0 0 1 1.11 0l3 2A1 1 0 0 1 22 17v4a1 1 0 0 1-1 1z"))),
		x.Child(x.Path(x.D("M18 10a8 8 0 0 0-16 0c0 4.993 5.539 10.193 7.399 11.799a1 1 0 0 0 .601.2"))),
		x.Child(x.Path(x.D("M18 22v-3"))),
		x.Child(x.Circle(x.Cx("10"), x.Cy("10"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
