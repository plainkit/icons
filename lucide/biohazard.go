package lucide

import x "github.com/bloxui/blox"

// Biohazard creates a Biohazard Lucide icon.
func Biohazard(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-biohazard", args)
	svgArgs = append(svgArgs,
		x.Child(x.Circle(x.Cx("12"), x.Cy("11.9"), x.R("2"))),
		x.Child(x.Path(x.D("M6.7 3.4c-.9 2.5 0 5.2 2.2 6.7C6.5 9 3.7 9.6 2 11.6"))),
		x.Child(x.Path(x.D("m8.9 10.1 1.4.8"))),
		x.Child(x.Path(x.D("M17.3 3.4c.9 2.5 0 5.2-2.2 6.7 2.4-1.2 5.2-.6 6.9 1.5"))),
		x.Child(x.Path(x.D("m15.1 10.1-1.4.8"))),
		x.Child(x.Path(x.D("M16.7 20.8c-2.6-.4-4.6-2.6-4.7-5.3-.2 2.6-2.1 4.8-4.7 5.2"))),
		x.Child(x.Path(x.D("M12 13.9v1.6"))),
		x.Child(x.Path(x.D("M13.5 5.4c-1-.2-2-.2-3 0"))),
		x.Child(x.Path(x.D("M17 16.4c.7-.7 1.2-1.6 1.5-2.5"))),
		x.Child(x.Path(x.D("M5.5 13.9c.3.9.8 1.8 1.5 2.5"))),
	)
	return x.Svg(svgArgs...)
}
