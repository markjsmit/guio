package element

import (
	"errors"
	
	"github.com/beevik/etree"
)

func FromXml(fileContent []byte, rootElement string) (Element, error) {
	doc := etree.NewDocument()
	doc.ReadFromBytes(fileContent)
	elem := doc.SelectElement(rootElement)
	
	if(elem==nil){
		return Element{},errors.New("Element "+rootElement+" not found")
	}
	root := Element{
		Tag:      rootElement,
		Attr:     extractXmlAttributes(elem.Attr),
		Children: extractXmlElements(elem.ChildElements()),
	}
	
	return root, nil
	
}

func extractXmlAttributes(attributes []etree.Attr) map[string]string {
	outputMap := map[string]string{}
	if attributes != nil {
		for _, attr := range attributes {
			outputMap[attr.Key] = attr.Value
		}
	}
	return outputMap
}

func extractXmlElements(elements []*etree.Element) []Element {
	outputArray := []Element{}
	if elements != nil {
		for _, elem := range elements {
			newElement := Element{
				Tag:      elem.Tag,
				Attr:     extractXmlAttributes(elem.Attr),
				Children: extractXmlElements(elem.ChildElements()),
			}
			outputArray = append(outputArray, newElement)
		}
	}
	return outputArray
	
}
