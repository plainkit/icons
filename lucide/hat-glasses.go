package lucide

import x "github.com/bloxui/blox"

// HatGlasses creates a Hat Glasses Lucide icon.
func HatGlasses(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-hat-glasses", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 18a2 2 0 0 0-4 0"))),
		x.Child(x.Path(x.D("m19 11-2.11-6.657a2 2 0 0 0-2.752-1.148l-1.276.61A2 2 0 0 1 12 4H8.5a2 2 0 0 0-1.925 1.456L5 11"))),
		x.Child(x.Path(x.D("M2 11h20"))),
		x.Child(x.Circle(x.Cx("17"), x.Cy("18"), x.R("3"))),
		x.Child(x.Circle(x.Cx("7"), x.Cy("18"), x.R("3"))),
	)
	return x.Svg(svgArgs...)
}
