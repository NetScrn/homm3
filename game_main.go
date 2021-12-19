package main

import (
	_ "embed"
	"flag"
	"game_try/game"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"runtime/pprof"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed maps/simple_sample.yaml
var sampleMap []byte

//go:embed configs/translations/all.yaml
var translationsYaml []byte

const language = "ru"

var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Heroes of Might and Magic III")

	var translations map[string]map[string]string
	err := yaml.Unmarshal(translationsYaml, &translations)
	if err != nil {
		log.Fatal(err)
	}

	translationsForLang := map[string]string{}
	for k, t := range translations {
		translationsForLang[k] = t[language]
	}

	homm3Game, err := game.NewHoMM3Game(sampleMap, translationsForLang)
	if err != nil {
		log.Fatal(err)
	}

	if err = ebiten.RunGame(homm3Game); err != nil {
		log.Fatal(err)
	}
}
