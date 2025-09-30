package lucide

import (
	html "github.com/plainkit/html"
)

// GitFork creates a Git Fork Lucide icon.
func GitFork(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-git-fork", args)
	children := []html.SvgArg{
		html.SvgCircle(html.ACx("12"), html.ACy("18"), html.AR("3")),
		html.SvgCircle(html.ACx("6"), html.ACy("6"), html.AR("3")),
		html.SvgCircle(html.ACx("18"), html.ACy("6"), html.AR("3")),
		html.SvgPath(html.AD("M18 9v2c0 .6-.4 1-1 1H7c-.6 0-1-.4-1-1V9")),
		html.SvgPath(html.AD("M12 12v3")),
	}
	return html.Svg(append(svgArgs, children...)...)
}
