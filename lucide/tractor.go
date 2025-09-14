package lucide

import x "github.com/bloxui/blox"

// Tractor creates a Tractor Lucide icon.
func Tractor(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-tractor", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m10 11 11 .9a1 1 0 0 1 .8 1.1l-.665 4.158a1 1 0 0 1-.988.842H20"))),
		x.Child(x.Path(x.D("M16 18h-5"))),
		x.Child(x.Path(x.D("M18 5a1 1 0 0 0-1 1v5.573"))),
		x.Child(x.Path(x.D("M3 4h8.129a1 1 0 0 1 .99.863L13 11.246"))),
		x.Child(x.Path(x.D("M4 11V4"))),
		x.Child(x.Path(x.D("M7 15h.01"))),
		x.Child(x.Path(x.D("M8 10.1V4"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("18"), x.R("2"))),
		x.Child(x.Circle(x.Cx("7"), x.Cy("15"), x.R("5"))),
	)
	return x.Svg(svgArgs...)
}
