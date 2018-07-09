package output

import (
	"bytes"
	"protoapi/generator/data"
	"text/template"
)

var javaTypes = map[string]string{
	// https://developers.google.com/protocol-buffers/docs/proto#scalar
	"double":   "double",
	"float":    "float",
	"int32":    "int",
	"int64":    "long",
	"uint32":   "int",
	"uint64":   "long",
	"sint32":   "int",
	"sint64":   "long",
	"fixed32":  "int",
	"fixed64":  "long",
	"sfixed32": "int",
	"sfixed64": "long",
	"bool":     "boolean",
	"string":   "String",
	"bytes":    "ByteString",
}

func toJavaType(dataType string) string {
	var javaType = javaTypes[dataType]
	if javaType != "" {
		return javaType
	}
	// if not primary type return data type and ignore the . in the data type
	return dataType[1:]
}

type springGen struct {
	ApplicationName string
	PackageName     string
	structTpl       *template.Template
	serviceTpl      *template.Template
}

func newSpringGen(applicationName, packageName string) *springGen {
	gen := &springGen{
		ApplicationName: applicationName,
		PackageName:     packageName,
	}
	gen.init()
	return gen
}

func (g *springGen) getTpl(path string) *template.Template {
	var err error
	tpl := template.New("tpl")
	tplStr := data.LoadTpl(path)
	result, err := tpl.Parse(tplStr)
	if err != nil {
		panic(err)
	}
	return result
}

func (g *springGen) init() {
	g.structTpl = g.getTpl("/generator/template/spring_struct.gojava")
	g.serviceTpl = g.getTpl("/generator/template/spring_service.gojava")
}

func (g *springGen) getStructFilename(msg *data.MessageData) string {
	return msg.Name + ".java"
}

func (g *springGen) genStruct(msg *data.MessageData) string {
	buf := bytes.NewBufferString("")

	obj := newSpringStruct(msg, g.PackageName)
	err := g.structTpl.Execute(buf, obj)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (g *springGen) genServie(service *data.ServiceData) string {
	buf := bytes.NewBufferString("")

	obj := newSpringService(service, g.PackageName)
	err := g.serviceTpl.Execute(buf, obj)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func genSpringCode(applicationName string, packageName string, service *data.ServiceData, messages []*data.MessageData, enums []*data.EnumData) (result map[string]string, err error) {
	gen := newSpringGen(applicationName, packageName)
	result = make(map[string]string)

	for _, msg := range messages {
		filename := gen.getStructFilename(msg)
		content := gen.genStruct(msg)

		result[filename] = content
	}

	// make file name same as java class name
	filename := service.Name + "Base.java"
	content := gen.genServie(service)
	result[filename] = content

	return
}

func init() {
	data.OutputMap["spring"] = genSpringCode
}