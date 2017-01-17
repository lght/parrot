package export

import (
	"github.com/anthonynsimon/parrot/parrot-api/model"
	"gopkg.in/yaml.v2"
)

type Yaml struct{}

func (e *Yaml) FileExtension() string {
	return "po"
}

func (e *Yaml) Export(locale *model.Locale) ([]byte, error) {
	data := make(map[string]map[string]interface{})
	data[locale.Ident] = make(map[string]interface{})
	for k, v := range locale.Pairs {
		data[locale.Ident][k] = v
	}

	result, err := yaml.Marshal(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// func (e *Yaml) ExportNested(locale *model.Locale) ([]byte, error) {
// 	data := make(map[string]map[string]interface{})
// 	root := data[locale.Ident]
// 	for k, v := range locale.Pairs {
// 		nesting := strings.Split(k, ".")
// 		current := root
// 		for i, nk := range nesting {
// 			if i < len(nesting) {
// 				current[nk] = make(map[string]interface{})
// 				current, ok := current[nk].(map[string]interface{})
// 				if !ok {
// 					return nil,
// 				}
// 				current = current[nk]
// 			} else {
// 				current[nk] = v
// 			}
// 		}
// 	}

// 	result, err := yaml.Marshal(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }
