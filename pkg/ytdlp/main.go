package ytdlp

import (
	"context"
	"io"
	"os"
	"os/exec"
	"strings"
)

type DownloadOptions struct {
	AudioFormat      string
	BinaryPath       string
	SponsorblockCats []string
}

func DownloadAudio(ctx context.Context, w io.Writer, url string, opt DownloadOptions) error {
	if err := opt.parse(); err != nil {
		return err
	}
	tmp, err := os.MkdirTemp("", "ytpodproxy")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)
	//nolint:gosec
	cmd := exec.CommandContext(ctx, opt.BinaryPath, opt.prepareArgs(url)...)
	cmd.Stdout = w
	cmd.Stderr = os.Stderr
	cmd.Dir = tmp
	return cmd.Run()
}

func (opt *DownloadOptions) parse() error {
	if len(opt.BinaryPath) != 0 {
		return nil
	}
	path, err := exec.LookPath("yt-dlp")
	if err != nil {
		return err
	}
	opt.BinaryPath = path
	if len(opt.AudioFormat) == 0 {
		opt.AudioFormat = "mp3"
	}
	return nil
}

func (opt *DownloadOptions) prepareArgs(url string) []string {
	args := []string{
		"--quiet",
		"--verbose",
		"-o", "-",
		"--extract-audio",
		"--audio-format", opt.AudioFormat,
		"-f", "bestaudio",
		"--no-playlist",
	}
	if len(opt.SponsorblockCats) != 0 {
		args = append(args, "--sponsorblock-remove", strings.Join(opt.SponsorblockCats, ","))
	}
	args = append(args, url)
	return args
}
