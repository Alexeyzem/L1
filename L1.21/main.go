//Реализовать паттерн «адаптер» на любом примере.

package main

type dataRecording interface {
	SendToPostgresql()
}

type PostgresData struct {
}

func (d *PostgresData) SendToPostgresql() {
	//высылыет данные в постгрес.
}

type MariaDbData struct {
}

func (d *MariaDbData) convertToPostgres() {
	//конвертирует запрос с синтаксисом для данной бд в запрос с синтаксисом для постгреса
}

type MariaDbDataAdapter struct {
	mariaDbData *MariaDbData
}

func (adapter *MariaDbDataAdapter) SendToPostgresql() {
	adapter.mariaDbData.convertToPostgres() // адаптер сконвертировал запрос из мариидб в запрос для постгреса
}

func main() {

}
