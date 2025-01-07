package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {
	// Parse the multipart form:
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse form",
		})
	}

	// Get the file from the form:
	files := form.File["image"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No file uploaded",
		})
	}

	var filenames []string

	// Save the file:
	for _, file := range files {
		// Check if the file is an image:
		if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Only JPEG and PNG images are allowed",
			})
		}

		// Generate a unique filename:
		ext := filepath.Ext(file.Filename)
		baseFilename := fmt.Sprintf("%s-%s", "agtaimage", file.Filename[:len(file.Filename)-len(ext)])
		filename := baseFilename + ext
		filepath := fmt.Sprintf("./uploads/%s", filename)
		counter := 1

		// Ensure the filename is unique:
		for {
			if _, err := os.Stat(filepath); os.IsNotExist(err) {
				break
			}
			filename = fmt.Sprintf("%s-%d%s", baseFilename, counter, ext)
			filepath = fmt.Sprintf("./uploads/%s", filename)
			counter++
		}

		// Save the file to the specified path:
		if err := c.SaveFile(file, filepath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save file",
			})
		}

		filenames = append(filenames, filename)
	}

	return c.JSON(fiber.Map{
		"message":   "File uploaded successfully",
		"filenames": filenames,
	})
}

func DeleteImage(c *fiber.Ctx) error {
	// Get the filename from the request body:
	filename := c.FormValue("filename")
	if filename == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Filename is required",
		})
	}

	// Construct the file path:
	filepath := fmt.Sprintf("./uploads/%s", filename)

	// Delete the file from the filesystem:
	if err := os.Remove(filepath); err != nil {
		if os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "File not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete file",
		})
	}

	return c.JSON(fiber.Map{
		"message": "File deleted successfully",
	})
}
