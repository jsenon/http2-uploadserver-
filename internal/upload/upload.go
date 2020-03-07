package upload

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	mymetrics "github.com/jsenon/http2-uploadserver/internal/metrics"
	"github.com/opentracing/opentracing-go"
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
		mymetrics.UploadNOK.Inc()
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
		mymetrics.UploadNOK.Inc()
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Error().Msgf("Error ioRead: %v", err)
		mymetrics.UploadNOK.Inc()
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	mymetrics.UploadOK.Inc()

	// return that we have successfully uploaded our file!
	log.Info().Msgf("Successfully Uploaded File")
	w.Write([]byte(fmt.Sprintf("Successfully Uploaded File\n")))
}

// OStream implement upload for octectstream
func OStream(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("File Upload Octect Stream Hit")

	dir := viper.GetString("OUTPUTDIR")
	log.Debug().Msgf("Output directory: %v", dir)
	body := r.Body
	n, err := copystream(r.Context(), dir, body)
	if err != nil {
		return
	}
	mymetrics.UploadOK.Inc()
	w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))
	// return that we have successfully uploaded our file!
	log.Info().Msgf("Successfully Uploaded File")

}

func copystream(ctx context.Context, dir string, body io.Reader) (int64, error) {
	parent, _ := opentracing.StartSpanFromContext(ctx, "(*http2-uploaderserver).upload.copystream")
	log.Debug().Msgf("Have found a ctx, generate span %v", parent)
	defer parent.Finish()

	log.Info().Msgf("Starting Upload: %v", dir)

	targetfile := dir + "/result"
	file, err := os.Create(targetfile)
	if err != nil {
		log.Error().Msgf("Error Creating the File: %v", err)
		mymetrics.UploadNOK.Inc()
		return 0, err
	}
	defer file.Close()

	n, err := io.Copy(file, body)
	if err != nil {
		log.Error().Msgf("Error Copying the File: %v", err)
		mymetrics.UploadNOK.Inc()
		return 0, err
	}
	return n, nil

}
