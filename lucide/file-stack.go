package lucide

import x "github.com/plainkit/html"

// FileStack creates a File Stack Lucide icon.
func FileStack(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-stack", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M11 21a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1"))),
		x.Child(x.Path(x.D("M16 16a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1"))),
		x.Child(x.Path(x.D("M21 6a2 2 0 0 0-.586-1.414l-2-2A2 2 0 0 0 17 2h-3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1z"))),
	)
	return x.Svg(svgArgs...)
}
