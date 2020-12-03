package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {

	var rimes = map[string]string{
		"sourire":  "rougir",
		"soupire":  "bouillir",
		"haleine":  "poleine",
		"cœur":     "beurre",
		"toujours": "jamais",
		"Signé":    "Joséphine Vladimir",
	}
	t := template.Must(template.New("letter").Funcs(funcMap).Parse(letter))
	err := t.Execute(os.Stdout, rimes)
	if err != nil {
		log.Println("executing template:", err)
	}
}

var funcMap = template.FuncMap{
	"title": strings.Title,
}

// Define a template.
const letter = `
Le bonheur d'aimer.
Recueil : Les élégies et poésies diverses (1813)

Il est auprès de moi, sa main presse ma main,
Sa bouche s'embellit du plus charmant {{ .sourire}},
Son teint s'anime, je {{ .soupire }},
Sa tête mollement vient tomber sur mon sein ;
Là je respire son {{ .haleine }},
Son {{ .haleine }} en parfum plus douce que la fleur.

De ses bras l'amoureuse chaîne
Rapproche mon {{ .cœur }} de son {{ .cœur }} ;
Bientôt nos baisers se confondent,
Ils sont purs comme nos amours :
Nous demeurons sans voix ;
Seuls nos yeux se répondent ;
Ils se disent tout bas :
{{ title .toujours }}, {{ .toujours }}, {{ .toujours }} !


{{ .Signé }}

`
