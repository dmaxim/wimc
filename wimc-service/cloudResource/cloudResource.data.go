package cloudResource

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/dmaxim/wimc/database"
)

func getCloudResource(cloudResourceId int) (*CloudResource, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `SELECT CloudResourceId, 
	CloudId,
	Location,
	Name, 
	Notes
	FROM CloudResource
	WHERE CloudResourceId = ?`, cloudResourceId)

	resource := &CloudResource{}
	err := row.Scan(
		&resource.CloudResourceId,
		&resource.CloudId,
		&resource.Location,
		&resource.Name,
		&resource.Notes,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return resource, nil
}

func removeResource(cloudResourceId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `DELETE FROM CloudResource WHERE CloudResourceId = ?`, cloudResourceId)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func getCloudResourceList() ([]CloudResource, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `SELECT CloudResourceId, 
	CloudId,
	Location,
	Name, 
	Notes
	FROM CloudResource`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	resources := make([]CloudResource, 0)
	for results.Next() {
		var resource CloudResource
		results.Scan(&resource.CloudResourceId,
			&resource.CloudId,
			&resource.Location,
			&resource.Name,
			&resource.Notes)
		resources = append(resources, resource)
	}
	return resources, nil
}

func insertCloudResource(cloudResource CloudResource) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `INSERT INTO CloudResource (
		CloudId,
		Location,
		Name,
		Notes)
		VALUES (?, ?, ?, ?)`,
		cloudResource.CloudId,
		cloudResource.Location,
		cloudResource.Name,
		cloudResource.Notes)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return int(insertedId), nil
}

func updateCloudResource(cloudResource CloudResource) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if cloudResource.CloudResourceId == 0 {
		return errors.New("CloudResource has an invalid id")
	}
	_, err := database.DbConn.ExecContext(ctx, `UPDATE CloudResource SET 
		CloudId=?,
		Location=?,
		Name=?,
		Notes=?
		WHERE CloudResourceId=?`,
		cloudResource.CloudId,
		cloudResource.Location,
		cloudResource.Name,
		cloudResource.Notes,
		cloudResource.CloudResourceId)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil

}
