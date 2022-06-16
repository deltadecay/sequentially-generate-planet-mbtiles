package sequentiallygenerateplanetmbtiles

import (
	"log"
	"path/filepath"

	"github.com/lambdajack/lj_go/pkg/lj_archive"
)

type unzipInformation struct {
	srcPath  string
	destPath string
}

func unzipSourceData() {
	var unzipInfo = []unzipInformation{
		{srcPath: filepath.Join(pth.workingDir, "water-polygons-split-4326.zip"), destPath: pth.coastlineDir},
		{srcPath: filepath.Join(pth.workingDir, "ne_10m_urban_areas.zip"), destPath: pth.landCoverUrbanDepth},
		{srcPath: filepath.Join(pth.workingDir, "ne_10m_antarctic_ice_shelves_polys.zip"), destPath: pth.landCoverIceShelvesDepth},
		{srcPath: filepath.Join(pth.workingDir, "ne_10m_glaciated_areas.zip"), destPath: pth.landCoverGlaciatedDepth},
	}

	for _, info := range unzipInfo {
		log.Println("Unzipping", info.srcPath)
		lj_archive.Unzip(info.srcPath, info.destPath)
	}
}
