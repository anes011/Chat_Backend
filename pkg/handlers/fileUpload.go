package handlers

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/anes011/chat/pkg/utils"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Response struct {
	Success bool
	URL     string
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) //10MB
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file from form data", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Read the file into memory and encode it in Base64
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}
	base64Encoded := base64.StdEncoding.EncodeToString(fileBytes)

	// Construct the data URI
	mimeType := http.DetectContentType(fileBytes)
	dataURI := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Encoded)

	//Cloudinary upload
	cldURL := os.Getenv("CLOUDINARY_URL")

	cld, err := cloudinary.NewFromURL(cldURL)

	if err != nil {
		log.Fatalf("Failed to authenticate account, %v\n", err)
	}

	var ctx = context.Background()
	uploadResult, err := cld.Upload.Upload(
		ctx,
		dataURI,
		uploader.UploadParams{
			UniqueFilename: api.Bool(true),
			Overwrite:      api.Bool(true)})

	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}

	utils.RespondWithJson(w, 201, &Response{
		Success: true,
		URL:     uploadResult.SecureURL,
	})
}
