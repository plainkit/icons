package lucide

import x "github.com/bloxui/blox"

// TestTubeDiagonal creates a Test Tube Diagonal Lucide icon.
func TestTubeDiagonal(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-test-tube-diagonal", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M21 7 6.82 21.18a2.83 2.83 0 0 1-3.99-.01a2.83 2.83 0 0 1 0-4L17 3"))),
		x.Child(x.Path(x.D("m16 2 6 6"))),
		x.Child(x.Path(x.D("M12 16H4"))),
	)
	return x.Svg(svgArgs...)
}
