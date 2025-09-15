package lucide

import x "github.com/plainkit/html"

// MemoryStick creates a Memory Stick Lucide icon.
func MemoryStick(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-memory-stick", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M6 19v-3"))),
		x.Child(x.Path(x.D("M10 19v-3"))),
		x.Child(x.Path(x.D("M14 19v-3"))),
		x.Child(x.Path(x.D("M18 19v-3"))),
		x.Child(x.Path(x.D("M8 11V9"))),
		x.Child(x.Path(x.D("M16 11V9"))),
		x.Child(x.Path(x.D("M12 11V9"))),
		x.Child(x.Path(x.D("M2 15h20"))),
		x.Child(x.Path(x.D("M2 7a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v1.1a2 2 0 0 0 0 3.837V17a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-5.1a2 2 0 0 0 0-3.837Z"))),
	)
	return x.Svg(svgArgs...)
}
