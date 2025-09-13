package lucide

import x "github.com/bloxui/blox"

// FlashlightOff creates a Flashlight Off Lucide icon.
func FlashlightOff(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-flashlight-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M16 16v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V10c0-2-2-2-2-4"))),
		x.Child(x.Path(x.D("M7 2h11v4c0 2-2 2-2 4v1"))),
		x.Child(x.Line(x.X1("11"), x.X2("18"), x.Y1("6"), x.Y2("6"))),
		x.Child(x.Line(x.X1("2"), x.X2("22"), x.Y1("2"), x.Y2("22"))),
	)
	return x.Svg(svgArgs...)
}
