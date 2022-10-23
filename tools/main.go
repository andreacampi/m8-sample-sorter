package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-audio/wav"
	"github.com/h2non/filetype"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Pack struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type Ctx struct {
	// the base path of the output tree
	Outpath string
	// the path to the current directory in the output tree
	Curpath string
}

type m8processor func(ctx Ctx, path string) error

func main() {
	// log.SetLevel(log.DebugLevel) // FIXME use a flag

	packs, err := Load("data/packs.yaml")
	if err != nil {
		log.Fatalf("loading packs: %v", err)
	}

	output, err := ioutil.TempDir(".", "m8")
	if err != nil {
		log.Fatalf("opening output: %v", err)
	}

	for i, pack := range packs {
		log.Infof("Processing '%s' [%d/%d] from %s", pack.Name, i, len(packs), pack.Path)

		err := Process(pack, wavProcessor, output)
		if err != nil {
			log.Infof("processing pack: %v", err)
		}
	}
}

func Load(path string) ([]Pack, error) {
	var packs []Pack

	in, err := os.Open("data/packs.yaml")
	if err != nil {
		return packs, fmt.Errorf("opening input file: %v", err)
	}

	data, err := ioutil.ReadAll(in)
	if err != nil {
		return packs, fmt.Errorf("reading input file: %v", err)
	}

	if err = yaml.Unmarshal(data, &packs); err != nil {
		return packs, fmt.Errorf("unmarshaling: %v", err)
	}

	return packs, nil
}

func Process(pack Pack, processor m8processor, output string) error {
	outpath := filepath.Join(output, pack.Name)
	err := os.Mkdir(outpath, 0755)
	if err != nil {
		return fmt.Errorf("creating dir %s: %v", outpath, err)
	}

	ctx := Ctx{
		Outpath: output,
		Curpath: outpath,
	}

	err = filepath.Walk(pack.Path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("walking %s: %v", path, err)
			}
			if info.IsDir() {
				return nil
			}

			if strings.HasSuffix(path, "/.DS_Store") {
				if err = os.Remove(path); err == nil {
					log.Infof("deleted %s", path)
				}
				// deliberately ignore errors
				return nil
			}

			// ignore custom icon
			if strings.HasSuffix(path, "/Icon\r") {
				return nil
			}

			buf, err := ioutil.ReadFile(path)
			if err != nil {
				return fmt.Errorf("reading %s: %v", path, err)
			}

			kind, err := filetype.Match(buf)
			if err != nil {
				return fmt.Errorf("matching %s: %v", path, err)
			}

			if kind == filetype.Unknown {
				log.Debugf("%s: unknown file type", path)
				return nil
			}
			if !filetype.IsAudio(buf) {
				log.Debugf("%s: not an audio file", path)
				return nil
			}

			switch kind.MIME.Value {
			case "audio/x-wav":
				if err = processor(ctx, path); err != nil {
					return fmt.Errorf("processing %s: %v", path, err)
				}

			default:
				log.Debugf("%s: not a wav file", path)
				return nil
			}

			// panic("stop")
			return nil
		})
	if err != nil {
		return fmt.Errorf("listing directory: %v", err)
	}

	return nil
}

func wavProcessor(ctx Ctx, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("opening %s: %v", path, err)
	}
	defer f.Close()

	d := wav.NewDecoder(f)

	// first check that it's valid...
	if !d.IsValidFile() {
		return fmt.Errorf("%s is not a valid wav file")
	}

	// ... then the format...
	format := d.Format()
	if format.NumChannels != 1 && format.NumChannels != 2 {
		return fmt.Errorf("%s has invalid format: NumChannels=%d", path, format.NumChannels)
	}
	if format.SampleRate != 44100 {
		return fmt.Errorf("%s has invalid format: SampleRate=%d", path, format.SampleRate)
	}
	if d.SampleBitDepth() != 8 && d.SampleBitDepth() != 16 && d.SampleBitDepth() != 24 {
		return fmt.Errorf("%s has invalid format: SampleBitDepth=%d", path, d.SampleBitDepth())
	}

	d.ReadMetadata()
	if d.Err() != nil {
		return fmt.Errorf("read metadata %s: %v", path, err)
	}

	dur, err := d.Duration()
	if err != nil {
		return fmt.Errorf("extracting length %s: %v", path, err)
	}
	if dur > 30*time.Second {
		log.Infof("%s looks like a loop, metadata: %#v, duration: %s", path, d.Metadata, dur)
	}

	log.Infof("%s metadata: %#v duration: %s", path, d.Metadata, dur)

	base := filepath.Base(path)
	target := filepath.Join(ctx.Curpath, base)
	err = os.Symlink(path, target)
	if err != nil {
		return fmt.Errorf("creating symlink %s -> %s: %v", target, path, err)
	}

	return nil
}
