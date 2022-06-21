package mbtiles

import (
	"fmt"
	"log"
	"math/rand"
	"path/filepath"
	"strings"

	"github.com/lambdajack/sequentially-generate-planet-mbtiles/internal/docker"
)

func Generate(src, dst, coastline, landcover, config, process string, outAsDir bool, tilemaker *docker.Container) {
	src, err := filepath.Abs(src)
	if err != nil {
		log.Fatal(err)
	}
	dst, err = filepath.Abs(dst)
	if err != nil {
		log.Fatal(err)
	}

	if !outAsDir {
		var sb strings.Builder
		r := rand.Perm(10)
		for _, o := range r {
			sb.WriteString(fmt.Sprintf("%d", o))
		}

		dst = filepath.Join(dst, sb.String()+".mbtiles")
	}

	tilemaker.Volumes = []docker.Volume{
		{
			Container: "/src",
			Host:      filepath.Dir(src),
		},
		{
			Container: "/dst",
			Host:      filepath.Dir(dst),
		},
		{
			Container: "/coastline",
			Host:      coastline,
		},
		{
			Container: "/landcover",
			Host:      landcover,
		},
		{
			Container: "/config",
			Host:      filepath.Dir(config),
		},
		{
			Container: "/process",
			Host:      filepath.Dir(process),
		},
	}

	err = tilemaker.Execute([]string{"--input", "/src/" + filepath.Base(src), "--output", "/dst/" + filepath.Base(dst), "--config", "/config/" + filepath.Base(config), "--process", "/process/" + filepath.Base(process)})
	if err != nil {
		log.Fatal(err)
	}
}
