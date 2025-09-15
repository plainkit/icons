package lucide

import x "github.com/plainkit/blox"

// SquircleDashed creates a Squircle Dashed Lucide icon.
func SquircleDashed(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-squircle-dashed", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M13.77 3.043a34 34 0 0 0-3.54 0"))),
		x.Child(x.Path(x.D("M13.771 20.956a33 33 0 0 1-3.541.001"))),
		x.Child(x.Path(x.D("M20.18 17.74c-.51 1.15-1.29 1.93-2.439 2.44"))),
		x.Child(x.Path(x.D("M20.18 6.259c-.51-1.148-1.291-1.929-2.44-2.438"))),
		x.Child(x.Path(x.D("M20.957 10.23a33 33 0 0 1 0 3.54"))),
		x.Child(x.Path(x.D("M3.043 10.23a34 34 0 0 0 .001 3.541"))),
		x.Child(x.Path(x.D("M6.26 20.179c-1.15-.508-1.93-1.29-2.44-2.438"))),
		x.Child(x.Path(x.D("M6.26 3.82c-1.149.51-1.93 1.291-2.44 2.44"))),
	)
	return x.Svg(svgArgs...)
}
