package lucide

import x "github.com/bloxui/blox"

// MousePointer creates a Mouse Pointer Lucide icon.
func MousePointer(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-mouse-pointer", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12.586 12.586 19 19"))),
		x.Child(x.Path(x.D("M3.688 3.037a.497.497 0 0 0-.651.651l6.5 15.999a.501.501 0 0 0 .947-.062l1.569-6.083a2 2 0 0 1 1.448-1.479l6.124-1.579a.5.5 0 0 0 .063-.947z"))),
	)
	return x.Svg(svgArgs...)
}
