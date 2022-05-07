package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"text/template"

	"github.com/iancoleman/strcase"
	"golang.org/x/mod/modfile"
)

func GetModuleName() string {
	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		log.Fatalln("error load go.mod")
	}

	modName := modfile.ModulePath(goModBytes)
	return modName
}

func createFileUsingTemplate(t *template.Template, filename string, data interface{}) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, data)
	if err != nil {
		return err
	}

	return nil
}

func InjectModule(filename string, flag string, data string) error {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	output := bytes.Replace(input, []byte(flag), []byte(data), -1)

	if err = ioutil.WriteFile(filename, output, 0664); err != nil {
		return err
	}
	return nil
}

func TemplateToString(t *template.Template, data interface{}) (string, error) {
	var tmpl bytes.Buffer
	if err := t.Execute(&tmpl, data); err != nil {
		return "", err
	}
	return tmpl.String(), nil
}

type model struct {
	Name     string
	DataType string
	Json     string
	Label    string
}

type databaseModel struct {
	FieldName string
	DataType  string
	Column    string
}

type codeTemplate struct {
	Source      string
	Destination string
	FileName    string
	Result      string
	Flag        string
}

type generator struct {
	ModuleName     string
	PackageName    string
	EntityName     string
	TableName      string
	SingularName   string
	PluralName     string
	DatabaseModels []databaseModel
	Models         []model
	Templates      []codeTemplate
}

func main() {

	arrModels := []model{
		{
			Name:     "title",
			DataType: "string",
			Label:    "title",
		},
		{
			Name:     "description",
			DataType: "string",
			Label:    "description",
		},
	}

	data := generator{
		ModuleName:   GetModuleName(),
		PackageName:  "sensor",
		EntityName:   "Sensor",
		TableName:    "sensors",
		SingularName: "sensor",
		PluralName:   "sensors",
		Models:       arrModels,
	}

	templates := []codeTemplate{
		{
			Source:      "./template/model/database.tmpl",
			Destination: "./model/database/",
			FileName:    fmt.Sprintf("%s.go", data.PackageName),
		},
		{
			Source:      "./template/model/model.tmpl",
			Destination: "./model/",
			FileName:    fmt.Sprintf("%s_model.go", data.PackageName),
		},
		{
			Source:      "./template/repository/repository.tmpl",
			Destination: fmt.Sprintf("./api/repository/%s/", data.PackageName),
			FileName:    "repository.go",
		},
		{
			Source:      "./template/repository/repository.tmpl",
			Destination: fmt.Sprintf("./api/repository/%s/", data.PackageName),
			FileName:    "repository.go",
		},
		{
			Source:      "./template/repository/builder.tmpl",
			Destination: fmt.Sprintf("./api/repository/%s/", data.PackageName),
			FileName:    "builder.go",
		},
		{
			Source:      "./template/repository/delete.tmpl",
			Destination: fmt.Sprintf("./api/repository/%s/", data.PackageName),
			FileName:    fmt.Sprintf("delete%sByID.go", data.EntityName),
		},
		{
			Source:      "./template/repository/getAll.tmpl",
			Destination: fmt.Sprintf("./api/repository/%s/", data.PackageName),
			FileName:    fmt.Sprintf("GetAll%s.go", data.EntityName),
		},
		{
			Source:      "./template/repository/getByID.tmpl",
			Destination: fmt.Sprintf("./api/repository/%s/", data.PackageName),
			FileName:    fmt.Sprintf("Get%sByID.go", data.EntityName),
		},
		{
			Source:      "./template/repository/store.tmpl",
			Destination: fmt.Sprintf("./api/repository/%s/", data.PackageName),
			FileName:    fmt.Sprintf("store%s.go", data.EntityName),
		},
		{
			Source:      "./template/repository/updateByID.tmpl",
			Destination: fmt.Sprintf("./api/repository/%s/", data.PackageName),
			FileName:    fmt.Sprintf("update%sByID.go", data.EntityName),
		},
		{
			Source:      "./template/service/service.tmpl",
			Destination: fmt.Sprintf("./api/service/%s/", data.PackageName),
			FileName:    "service.go",
		},
		{
			Source:      "./template/service/deleteByID.tmpl",
			Destination: fmt.Sprintf("./api/service/%s/", data.PackageName),
			FileName:    fmt.Sprintf("delete%sByID.go", data.EntityName),
		},
		{
			Source:      "./template/service/getAll.tmpl",
			Destination: fmt.Sprintf("./api/service/%s/", data.PackageName),
			FileName:    fmt.Sprintf("getAll%s.go", data.EntityName),
		},
		{
			Source:      "./template/service/getByID.tmpl",
			Destination: fmt.Sprintf("./api/service/%s/", data.PackageName),
			FileName:    fmt.Sprintf("Get%sByID.go", data.EntityName),
		},
		{
			Source:      "./template/service/store.tmpl",
			Destination: fmt.Sprintf("./api/service/%s/", data.PackageName),
			FileName:    fmt.Sprintf("store%s.go", data.EntityName),
		},
		{
			Source:      "./template/service/updateByID.tmpl",
			Destination: fmt.Sprintf("./api/service/%s/", data.PackageName),
			FileName:    fmt.Sprintf("update%sByID.go", data.EntityName),
		},
		{
			Source:      "./template/controller/controller.tmpl",
			Destination: fmt.Sprintf("./api/controller/%s/", data.PackageName),
			FileName:    "controller.go",
		},
		{
			Source:      "./template/controller/handleCreate.tmpl",
			Destination: fmt.Sprintf("./api/controller/%s/", data.PackageName),
			FileName:    fmt.Sprintf("handleCreate%s.go", data.EntityName),
		},
		{
			Source:      "./template/controller/handleDelete.tmpl",
			Destination: fmt.Sprintf("./api/controller/%s/", data.PackageName),
			FileName:    fmt.Sprintf("handleDelete%s.go", data.EntityName),
		},
		{
			Source:      "./template/controller/handleGetAll.tmpl",
			Destination: fmt.Sprintf("./api/controller/%s/", data.PackageName),
			FileName:    fmt.Sprintf("handleGetAll%s.go", data.EntityName),
		},
		{
			Source:      "./template/controller/handleShow.tmpl",
			Destination: fmt.Sprintf("./api/controller/%s/", data.PackageName),
			FileName:    fmt.Sprintf("handleShow%s.go", data.EntityName),
		},
		{
			Source:      "./template/controller/handleUpdate.tmpl",
			Destination: fmt.Sprintf("./api/controller/%s/", data.PackageName),
			FileName:    fmt.Sprintf("handleUpdate%s.go", data.EntityName),
		},
	}

	injector := []codeTemplate{
		{
			Source:      "./template/module/repository/import_repository.tmpl",
			Destination: "./api/repository/",
			FileName:    "repository.go",
			Flag:        "// IMPORT REPOSITORY PACKAGE HERE",
		},
		{
			Source:      "./template/module/repository/inject_repository.tmpl",
			Destination: "./api/repository/",
			FileName:    "repository.go",
			Flag:        "// INJECT REPOSITORY HERE",
		},
		{
			Source:      "./template/module/repository/inject_repository_module.tmpl",
			Destination: "./api/repository/",
			FileName:    "repository.go",
			Flag:        "// INJECT REPOSITORY MODULE HERE",
		},
		{
			Source:      "./template/module/service/import_service.tmpl",
			Destination: "./api/service/",
			FileName:    "service.go",
			Flag:        "// IMPORT SERVICE PACKAGE HERE",
		},
		{
			Source:      "./template/module/service/inject_service.tmpl",
			Destination: "./api/service/",
			FileName:    "service.go",
			Flag:        "// INJECT SERVICE HERE",
		},
		{
			Source:      "./template/module/service/inject_service_module.tmpl",
			Destination: "./api/service/",
			FileName:    "service.go",
			Flag:        "// INJECT SERVICE MODULE HERE",
		},
		{
			Source:      "./template/module/controller/import_controller.tmpl",
			Destination: "./api/controller/",
			FileName:    "controller.go",
			Flag:        "// IMPORT CONTROLLER PACKAGE HERE",
		},
		{
			Source:      "./template/module/controller/inject_controller.tmpl",
			Destination: "./api/controller/",
			FileName:    "controller.go",
			Flag:        "// INJECT CONTROLLER HERE",
		},
		{
			Source:      "./template/module/route/inject_route.tmpl",
			Destination: "./api/route/",
			FileName:    "http_route.go",
			Flag:        "// INJECT ROUTE HERE",
		},
	}

	for i := 0; i < len(templates); i++ {
		result, err := os.ReadFile(templates[i].Source)
		if err != nil {
			log.Fatalf("%s not found \n", templates[i].Source)
		}
		if _, err := os.Stat(templates[i].Destination); os.IsNotExist(err) {
			err := os.Mkdir(templates[i].Destination, 0775)
			if err != nil {
				log.Fatalf("error when create folder %s \n", templates[i].Destination)
			}
		}
		templates[i].Result = string(result)
	}

	arrDbModels := []databaseModel{}

	for i := 0; i < len(arrModels); i++ {
		arrModels[i].Name = strcase.ToCamel(arrModels[i].Name)
		arrModels[i].Json = strcase.ToLowerCamel(arrModels[i].Name)
		arrModels[i].Label = strcase.ToLowerCamel(arrModels[i].Label)
		arrDbModels = append(arrDbModels, databaseModel{
			FieldName: arrModels[i].Name,
			DataType:  arrModels[i].DataType,
			Column:    strcase.ToSnake(arrModels[i].Name),
		})
	}
	data.DatabaseModels = arrDbModels

	for i := 0; i < len(templates); i++ {
		temp := templates[i]
		t := template.Must(template.New("").Parse(string(temp.Result)))
		fullPath := fmt.Sprintf("%s%s", temp.Destination, temp.FileName)

		err := createFileUsingTemplate(t, fullPath, data)
		if err != nil {
			log.Fatalf("error when creating file  %s\n", fullPath)
		}
	}

	for i := 0; i < len(injector); i++ {
		result, err := os.ReadFile(injector[i].Source)
		if err != nil {
			log.Fatalf("%s not found \n", injector[i].Source)
		}
		t := template.Must(template.New("").Parse(string(result)))
		text, err := TemplateToString(t, data)
		if err != nil {
			log.Fatalf("%s inject error \n", injector[i].Source)
		}
		fullPath := fmt.Sprintf("%s%s", injector[i].Destination, injector[i].FileName)
		err = InjectModule(fullPath, injector[i].Flag, text)
		if err != nil {
			log.Fatalf("%s inject error \n", injector[i].Source)
		}
	}

	cmd := exec.Command("gofmt", "-s", "-w", ".")
	_, err := cmd.Output()

	if err != nil {
		log.Fatalf("error formatting \n")
	}

}
