package filterMetadata

import (
	"encoding/json"
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

	// Consulta en la tabla file
	var id int
	var namefile, metadata string
	query := "SELECT id, namefile, metadata FROM file WHERE uuid = ? LIMIT 1" //

	err = dbFacade.QueryRowByField(query, fileId, &id, &namefile, &metadata)
	if err != nil {
		log.Printf("Error al obtener info del archivo con uuid: %s [ERROR] ->[%v]", fileId, err)
	} else {
		fmt.Printf("archivo encontrado: ID=%d, NombreArchivo=%s, metadata=%s\n", id, namefile, metadata)

		updatedMetadata := ComplementMetadata(namefile, metadata, fileId)
		SaveData(updatedMetadata)
		fmt.Printf("Metadata actualizada: %s\n", updatedMetadata)

	}

}

// ComplementMetadata complementa el JSON de metadata con fileName y fileId
func ComplementMetadata(namefile, metadata, fileId string) string {
	// Parsear metadata a un mapa
	metadataMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(metadata), &metadataMap); err != nil {
		log.Printf("Error al analizar el JSON de metadata: %v", err)
		return metadata // Retornar la metadata original en caso de error
	}

	// Agregar valores a metadata
	metadataMap["fileName"] = namefile
	metadataMap["fileId"] = fileId

	// Convertir de nuevo a JSON
	updatedMetadata, err := json.Marshal(metadataMap)
	if err != nil {
		log.Printf("Error al convertir metadata a JSON: %v", err)
		return metadata // Retornar la metadata original en caso de error
	}

	return string(updatedMetadata)
}

func SaveData(updatedMetadata string) {
	cfg, err := config.GetConnectionDatabaseNoSQL()

	if err != nil {
		log.Fatalf("Error al obtener las propiedades de conexión a base de datos: %v", err)
	}

	factoryMongo := &factoryDB.DatabaseFactory{}

	connector, err := factoryMongo.GetDatabaseConnectorNoSQL(cfg.Database)
	if err != nil {
		log.Fatalf("Error al obtener el conector: %v", err)
	}

	dbFacade, err := facadeDB.NewDatabaseFacadeNoSql(connector, cfg.ConnectionString)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	var metadataDoc interface{}

	err = json.Unmarshal([]byte(updatedMetadata), &metadataDoc)

	if err != nil {
		log.Printf("Error al convertir la metadata actualizada a documento: %v", err)
		return
	}

	if err := dbFacade.Insert("metadataFiles", metadataDoc); err != nil {
		log.Printf("Error al guardar el documento: %v", err)
	} else {
		log.Println("Documento guardado exitosamente en MongoDB")

	}

}
