package upload

import (
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

// File Upload the file
func File(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Error().Msgf("Error Retrieving the File", err)
		return
	}
	defer file.Close()
	log.Info().Msgf("Uploaded File: %+v\n", handler.Filename)
	log.Info().Msgf("File Size: %+v\n", handler.Size)
	log.Info().Msgf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("/datas", "upload-*.jpeg")
	if err != nil {
		log.Error().Msgf("Error ioCreate", err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error().Msgf("Error ioRead: ", err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	log.Info().Msgf("Successfully Uploaded File\n")
}
