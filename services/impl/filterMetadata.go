package filterMetadata

import (
	"fmt"
	"log"
	config "tuberias/config"
	facadeDB "tuberias/infraestructure/facade"
	factoryDB "tuberias/infraestructure/factory"
)

func FiletMetadata(fileId string) {

	cfg, error := config.GetConnectionDatabse()

	if error != nil {
		log.Fatalf("Error al obtener las propiedades de conexión a base de datos: %v", error)
	}

	factory := &factoryDB.DatabaseFactory{}

	connector, err := factory.GetDatabaseConnector(cfg.Database)

	if err != nil {
		log.Fatalf("Error al obtener el conector: %v", err)
	}

	dbFacade, err := facadeDB.NewDatabaseFacade(connector, cfg.ConnectionString)

	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer dbFacade.Close()

	// Define el correo a buscar

	// Consulta en la tabla usuarios
	var id int
	var namefile, metadata string
	query := "SELECT id, namefile, metadata FROM file WHERE uuid = ? LIMIT 1" // Cambia ? por $1 para PostgreSQL

	err = dbFacade.QueryRowByField(query, fileId, &id, &namefile, &metadata)
	if err != nil {
		log.Fatalf("Error al obtener el usuario: %v", err)
	}

	fmt.Printf("Usuario encontrado: ID=%d, Nombre=%s, Correo=%s\n", id, namefile, metadata)

}
