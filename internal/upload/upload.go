package upload

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// File Upload the file
func File(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("File Upload Endpoint Hit")

	dir := viper.GetString("OUTPUTDIR")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Error().Msgf("Error Retrieving the File: %v", err)
		return
	}
	defer file.Close()
	log.Info().Msgf("Uploaded File: %v", handler.Filename)
	log.Info().Msgf("File Size: %v", handler.Size)
	log.Info().Msgf("MIME Header: %v", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile(dir, "upload-*.jpeg")
	if err != nil {
		log.Error().Msgf("Error ioCreate %v", err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error().Msgf("Error ioRead: %v", err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	log.Info().Msgf("Successfully Uploaded File")
	w.Write([]byte(fmt.Sprintf("Successfully Uploaded File\n")))
}

// OStream implement upload for octectstream
func OStream(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("File Upload Octect Stream Hit")

	dir := viper.GetString("OUTPUTDIR")
	log.Info().Msgf("Output directory: %v", dir)

	targetfile := dir + "/result"
	file, err := os.Create(targetfile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	n, err := io.Copy(file, r.Body)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))
	// return that we have successfully uploaded our file!
	log.Info().Msgf("Successfully Uploaded File")
}
