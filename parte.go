// Parte un archvo con 25,000,000 de rergistros en formato csv,
//  a 10 archivos con 2,500,010 registros.
// Archivo de entrada: ratings.csv
// 10 archivos de salida:
// parte_1.csv ... parte_10.csv

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Ruta del archivo CSV de entrada
	inputFile := "ratings.csv"
	// Número de líneas por cada archivo dividido
	linesPerFile := 2500010

	// Abrir el archivo CSV de entrada
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %v\n", err)
		return
	}
	defer file.Close()

	// Crear un lector de CSV
	reader := csv.NewReader(file)

	// Leer todas las filas del archivo CSV
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error al leer el archivo: %v\n", err)
		return
	}

	// Dividir y escribir las filas en múltiples archivos
	totalRows := len(rows)
	fileCount := 1
	for i := 0; i < totalRows; i += linesPerFile {
		// Determinar el rango de filas para cada archivo
		end := i + linesPerFile
		if end > totalRows {
			end = totalRows
		}

		// Nombre del archivo de salida
		outputFile := "parte_" + strconv.Itoa(fileCount) + ".csv"

		// Crear el archivo CSV de salida
		outFile, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("Error al crear el archivo: %v\n", err)
			return
		}
		defer outFile.Close()

		// Crear un escritor de CSV
		writer := csv.NewWriter(outFile)

		// Escribir las filas en el nuevo archivo CSV
		err = writer.WriteAll(rows[i:end])
		if err != nil {
			fmt.Printf("Error al escribir en el archivo: %v\n", err)
			return
		}

		writer.Flush()
		
		fileCount++
	}

	fmt.Println("El archivo CSV se ha dividido correctamente.")
}
