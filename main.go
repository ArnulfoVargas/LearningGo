package main

import (
	printing "main/001_Print"
	variables "main/002_Variables"
	inputs "main/003_Input"
	conditionals "main/004_Conditionals"
	loops "main/005_Loops"
	pointers "main/006_Pointers"
	functions "main/007_Functions"
	datastructures "main/008_DataStructures"
	mapsObject "main/009_Maps"
	datatypesdefinition "main/010_DataTypesDefinition"
	structs "main/011_Structs"
	interfaces "main/012_Interfaces"
	iowriterreader "main/013_IOWriterReader"

	"github.com/mariomac/analizador"
)

func main() { 
	printing.Prints()
	variables.VariablesTest()
	inputs.InputsTest()
	conditionals.ConditionalsTest()
	loops.LoopsTest()
	pointers.PointersTest()
	functions.FunctionsTest()
	datastructures.DataStructures()
	mapsObject.MapsTest()

	// Imported from Github
	analizador.PrintEstadistica("Arnulfo")

	datatypesdefinition.DataTypesDefinitionTest()
	structs.StructsTest()
	interfaces.InterfacesTest()
	iowriterreader.IOsTest()
}