package lucide

import x "github.com/plainkit/blox"

// CircleFadingPlus creates a Circle Fading Plus Lucide icon.
func CircleFadingPlus(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-circle-fading-plus", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 2a10 10 0 0 1 7.38 16.75"))),
		x.Child(x.Path(x.D("M12 8v8"))),
		x.Child(x.Path(x.D("M16 12H8"))),
		x.Child(x.Path(x.D("M2.5 8.875a10 10 0 0 0-.5 3"))),
		x.Child(x.Path(x.D("M2.83 16a10 10 0 0 0 2.43 3.4"))),
		x.Child(x.Path(x.D("M4.636 5.235a10 10 0 0 1 .891-.857"))),
		x.Child(x.Path(x.D("M8.644 21.42a10 10 0 0 0 7.631-.38"))),
	)
	return x.Svg(svgArgs...)
}
