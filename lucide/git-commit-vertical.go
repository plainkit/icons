package lucide

import (
	html "github.com/plainkit/html"
)

// GitCommitVertical creates a Git Commit Vertical Lucide icon.
func GitCommitVertical(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-git-commit-vertical", args)
	children := []html.SvgArg{
		html.SvgPath(html.AD("M12 3v6")),
		html.SvgCircle(html.ACx("12"), html.ACy("12"), html.AR("3")),
		html.SvgPath(html.AD("M12 15v6")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
