package lucide

import x "github.com/plainkit/html"

// TestTube creates a Test Tube Lucide icon.
func TestTube(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-test-tube", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14.5 2v17.5c0 1.4-1.1 2.5-2.5 2.5c-1.4 0-2.5-1.1-2.5-2.5V2"))),
		x.Child(x.Path(x.D("M8.5 2h7"))),
		x.Child(x.Path(x.D("M14.5 16h-5"))),
	)
	return x.Svg(svgArgs...)
}
