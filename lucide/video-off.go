package lucide

import x "github.com/bloxui/blox"

// VideoOff creates a Video Off Lucide icon.
func VideoOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-video-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10.66 6H14a2 2 0 0 1 2 2v2.5l5.248-3.062A.5.5 0 0 1 22 7.87v8.196"))),
		x.Child(x.Path(x.D("M16 16a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h2"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
	)
	return x.Svg(svgArgs...)
}
