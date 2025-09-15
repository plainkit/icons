package lucide

import x "github.com/plainkit/blox"

// FileSymlink creates a File Symlink Lucide icon.
func FileSymlink(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-file-symlink", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("m10 18 3-3-3-3"))),
		x.Child(x.Path(x.D("M14 2v4a2 2 0 0 0 2 2h4"))),
		x.Child(x.Path(x.D("M4 11V4a2 2 0 0 1 2-2h9l5 5v13a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h7"))),
	)
	return x.Svg(svgArgs...)
}
