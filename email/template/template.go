package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
)

type EmailTemplate string

const (
	TransactionDetail EmailTemplate = "transaction_detail"
	TransactionToken  EmailTemplate = "transaction_token"
)

func GenerateEmail(temp EmailTemplate, itf interface{}, interfaces ...interface{}) string {
	tmpl := template.Must(template.ParseFiles(fmt.Sprintf("email/template/%s.html", temp)))
	var tpl bytes.Buffer
	data := toMap(itf)
	for _, itf := range interfaces {
		data = mergeMaps(data, toMap(itf))
	}

	if err := tmpl.Execute(&tpl, data); err != nil {
		log.Println(err)
	}

	return tpl.String()
}

func toMap(i interface{}) map[string]interface{} {
	b, _ := json.Marshal(&i)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	return m
}

func mergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
