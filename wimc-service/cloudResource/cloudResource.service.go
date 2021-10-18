package cloudResource

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dmaxim/wimc/cors"
)

const cloudResourcesBasePath = "azureresources"

func SetupRoutes(apiBasePath string) {
	handleCloudResources := http.HandlerFunc(azureResourcesHandler)
	handleCloudResource := http.HandlerFunc(azureResourceHandler)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, cloudResourcesBasePath), cors.Middleware(handleCloudResources))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, cloudResourcesBasePath), cors.Middleware(handleCloudResource))
}

func azureResourcesHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		resourceList, err := getCloudResourceList()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		resourcesJson, err := json.Marshal(resourceList)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			writer.Header().Set("Content-Type", "application/json")
			writer.Write(resourcesJson)
		}
	case http.MethodPost:
		var newResource CloudResource
		requestBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
		} else {
			err = json.Unmarshal(requestBytes, &newResource)
			if err != nil {
				writer.WriteHeader(http.StatusBadRequest)
			}
			if newResource.CloudResourceId != 0 {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			newId, err := insertCloudResource(newResource)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			writer.WriteHeader(http.StatusCreated)
			writer.Write([]byte(fmt.Sprintf(`{"cloudResourceId}:%d`, newId)))
			return
		}
	case http.MethodOptions:
		return
	}
}

func azureResourceHandler(writer http.ResponseWriter, request *http.Request) {
	urlPathSegments := strings.Split(request.URL.Path, "azureresources/")
	resourceId, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	switch request.Method {
	case http.MethodGet:
		resource, err := getCloudResource(resourceId)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		if resource == nil {
			writer.WriteHeader(http.StatusFound)
			return
		}
		resourceJson, err := json.Marshal(resource)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(resourceJson)
	case http.MethodPut:
		var updatedResource CloudResource
		err := json.NewDecoder(request.Body).Decode(&updatedResource)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatedResource.CloudResourceId != resourceId {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		err = updateCloudResource(updatedResource)
		if err != nil {
			log.Print(err)
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
		return

	case http.MethodDelete:
		removeResource(resourceId)
		return
	case http.MethodOptions:
		return
	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}
