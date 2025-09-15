package lucide

import x "github.com/plainkit/html"

// BrainCircuit creates a Brain Circuit Lucide icon.
func BrainCircuit(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-brain-circuit", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 5a3 3 0 1 0-5.997.125 4 4 0 0 0-2.526 5.77 4 4 0 0 0 .556 6.588A4 4 0 1 0 12 18Z"))),
		x.Child(x.Path(x.D("M9 13a4.5 4.5 0 0 0 3-4"))),
		x.Child(x.Path(x.D("M6.003 5.125A3 3 0 0 0 6.401 6.5"))),
		x.Child(x.Path(x.D("M3.477 10.896a4 4 0 0 1 .585-.396"))),
		x.Child(x.Path(x.D("M6 18a4 4 0 0 1-1.967-.516"))),
		x.Child(x.Path(x.D("M12 13h4"))),
		x.Child(x.Path(x.D("M12 18h6a2 2 0 0 1 2 2v1"))),
		x.Child(x.Path(x.D("M12 8h8"))),
		x.Child(x.Path(x.D("M16 8V5a2 2 0 0 1 2-2"))),
		x.Child(x.Circle(x.Cx("16"), x.Cy("13"), x.R(".5"))),
		x.Child(x.Circle(x.Cx("18"), x.Cy("3"), x.R(".5"))),
		x.Child(x.Circle(x.Cx("20"), x.Cy("21"), x.R(".5"))),
		x.Child(x.Circle(x.Cx("20"), x.Cy("8"), x.R(".5"))),
	)
	return x.Svg(svgArgs...)
}
