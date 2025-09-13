package lucide

import x "github.com/bloxui/blox"

// Delete creates a Delete Lucide icon.
func Delete(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-delete", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 5a2 2 0 0 0-1.344.519l-6.328 5.74a1 1 0 0 0 0 1.481l6.328 5.741A2 2 0 0 0 10 19h10a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2z"))),
		x.Child(x.Path(x.D("m12 9 6 6"))),
		x.Child(x.Path(x.D("m18 9-6 6"))),
	)
	return x.Svg(svgArgs...)
}
