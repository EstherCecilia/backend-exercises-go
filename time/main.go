package main

import (
	"fmt"
	"time"
)

func main() {
	// Obtendo a data e hora atuais
	currentTime := time.Now()
	fmt.Println("Data e Hora atuais:", currentTime)

	// Formatando a data e hora
	// "2006-01-02 15:04:05" é o formato padrão que o Go usa para representar a data e hora
	formattedTime := currentTime.Format("02/01/2006 15:04:05")
	fmt.Println("Data formatada:", formattedTime)

	// Adicionando 2 horas
	newTime := currentTime.Add(2 * time.Hour)
	fmt.Println("Hora após adicionar 2 horas:", newTime)

	// Subtraindo 30 minutos
	subtractedTime := currentTime.Add(-30 * time.Minute)
	fmt.Println("Hora após subtrair 30 minutos:", subtractedTime)

	// Calculando a diferença entre duas datas
	startTime := time.Date(2024, time.November, 1, 10, 0, 0, 0, time.UTC)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Println("Duração entre o início e agora:", duration)

	// Convertendo uma string para uma data
	dateString := "2024-11-01 10:00:00"
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("Erro ao parsear a data:", err)
		return
	}
	fmt.Println("Data parseada:", parsedTime)

	// Verificando se uma data é anterior ou posterior
	if currentTime.Before(parsedTime) {
		fmt.Println("A data atual é antes da data parseada")
	} else {
		fmt.Println("A data atual é depois da data parseada")
	}

	// Exibindo apenas a data ou hora
	fmt.Println("Somente a data:", currentTime.Format("02/01/2006"))
	fmt.Println("Somente a hora:", currentTime.Format("15:04:05"))
}
