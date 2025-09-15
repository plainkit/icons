package lucide

import x "github.com/plainkit/html"

// TestTubes creates a Test Tubes Lucide icon.
func TestTubes(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-test-tubes", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M9 2v17.5A2.5 2.5 0 0 1 6.5 22A2.5 2.5 0 0 1 4 19.5V2"))),
		x.Child(x.Path(x.D("M20 2v17.5a2.5 2.5 0 0 1-2.5 2.5a2.5 2.5 0 0 1-2.5-2.5V2"))),
		x.Child(x.Path(x.D("M3 2h7"))),
		x.Child(x.Path(x.D("M14 2h7"))),
		x.Child(x.Path(x.D("M9 16H4"))),
		x.Child(x.Path(x.D("M20 16h-5"))),
	)
	return x.Svg(svgArgs...)
}
