package lucide

import x "github.com/plainkit/html"

// Pointer creates a Pointer Lucide icon.
func Pointer(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-pointer", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M22 14a8 8 0 0 1-8 8"))),
		x.Child(x.Path(x.D("M18 11v-1a2 2 0 0 0-2-2a2 2 0 0 0-2 2"))),
		x.Child(x.Path(x.D("M14 10V9a2 2 0 0 0-2-2a2 2 0 0 0-2 2v1"))),
		x.Child(x.Path(x.D("M10 9.5V4a2 2 0 0 0-2-2a2 2 0 0 0-2 2v10"))),
		x.Child(x.Path(x.D("M18 11a2 2 0 1 1 4 0v3a8 8 0 0 1-8 8h-2c-2.8 0-4.5-.86-5.99-2.34l-3.6-3.6a2 2 0 0 1 2.83-2.82L7 15"))),
	)
	return x.Svg(svgArgs...)
}
