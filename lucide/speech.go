package lucide

import x "github.com/plainkit/blox"

// Speech creates a Speech Lucide icon.
func Speech(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-speech", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8.8 20v-4.1l1.9.2a2.3 2.3 0 0 0 2.164-2.1V8.3A5.37 5.37 0 0 0 2 8.25c0 2.8.656 3.054 1 4.55a5.77 5.77 0 0 1 .029 2.758L2 20"))),
		x.Child(x.Path(x.D("M19.8 17.8a7.5 7.5 0 0 0 .003-10.603"))),
		x.Child(x.Path(x.D("M17 15a3.5 3.5 0 0 0-.025-4.975"))),
	)
	return x.Svg(svgArgs...)
}
