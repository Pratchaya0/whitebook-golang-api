package book

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Pratchaya0/whitebook-golang-api/dtos/requests"
	"github.com/Pratchaya0/whitebook-golang-api/dtos/responses"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/Pratchaya0/whitebook-golang-api/helpers"
	"github.com/gin-gonic/gin"
)

// HandleUploadBookCoverImage validates and saves uploaded Image files
func handleUploadBookCoverImage(files []*multipart.FileHeader, uploadDir string, bookName string) (string, string) {
	var fileLinks []string

	// Validate Image files
	for _, file := range files {
		if !helpers.ValidateFileType(file, "image") {
			return "", fmt.Sprintf("File name '%s' is not a image type.", file.Filename)
		}
	}

	// Save
	for _, file := range files {
		fmt.Println("temp", file.Filename)
		// Generate a unique filename and save the file
		dst := filepath.Join(uploadDir, fmt.Sprintf("%s_%s%s", strings.ToLower(bookName), time.Now().Format("20060102"), filepath.Ext(file.Filename)))
		// if err := saveFile(file, dst); err != nil {
		// 	return "", fmt.Sprintf("Unable to save file '%s': %s", file.Filename, err.Error())
		// }
		// Generate file URL or path (assuming `uploadDir` is a relative path)
		fileLinks = append(fileLinks, dst)
	}

	return strings.Join(fileLinks, ","), "OK"
}

// HandleUploadBookPreviewImage validates and saves uploaded Image files
func handleUploadBookPreviewImage(files []*multipart.FileHeader, uploadDir string, bookId uint, bookName string) string {
	// var bookPreviewImages []entities.BookPreviewImage

	// validate Image file
	for _, file := range files {
		if !helpers.ValidateFileType(file, "image") {
			return fmt.Sprintf("File name '%s' are not image type.", file.Filename)
		}
	}

	// Save
	for i, file := range files {
		fmt.Println("temp", file.Filename)
		// Generate a unique filename and save the file
		dst := filepath.Join(uploadDir, fmt.Sprintf("%s_%s_%s%s", strings.ToLower(bookName), strconv.Itoa(i+1), time.Now().Format("20060102"), filepath.Ext(file.Filename)))
		// if err := saveFile(file, dst); err != nil {
		// 	return "", fmt.Sprintf("Unable to save file '%s': %s", file.Filename, err.Error())
		// }

		bookPreviewImageCreate := entities.BookPreviewImage{
			PageNumber: uint(i) + 1,
			ImageLink:  dst,

			BookID: bookId,
		}

		if err := entities.DB().Create(&bookPreviewImageCreate).Error; err != nil {
			return "Cannot create book preview image."
		}

		// bookPreviewImages = append(bookPreviewImages, bookPreviewImageCreate)
	}

	return "OK"
}

// HandleUploadBookPdf validates and saves uploaded PDF files
func handleUploadBookPdf(files []*multipart.FileHeader, uploadDir string, bookName string) (string, string) {
	var fileLinks []string

	// Validate PDF files
	for _, file := range files {
		if !helpers.ValidateFileType(file, "pdf") {
			return "", fmt.Sprintf("File name '%s' is not a PDF type.", file.Filename)
		}
	}

	// Save
	for _, file := range files {
		fmt.Println("temp", file.Filename)
		// Generate a unique filename and save the file
		dst := filepath.Join(uploadDir, fmt.Sprintf("%s_%s.pdf", strings.ToLower(bookName), time.Now().Format("20060102")))
		// if err := saveFile(file, dst); err != nil {
		// 	return "", fmt.Sprintf("Unable to save file '%s': %s", file.Filename, err.Error())
		// }

		// Generate file URL or path (assuming `uploadDir` is a relative path)
		fileLinks = append(fileLinks, dst)
	}

	return strings.Join(fileLinks, ","), "OK"
}

// HandleUploadBookEpub validates and saves uploaded EPUB files
func handleUploadBookEpub(files []*multipart.FileHeader, uploadDir string, bookName string) (string, string) {
	var fileLinks []string

	// Validate EPUB files
	for _, file := range files {
		if !helpers.ValidateFileType(file, "epub") {
			return "", fmt.Sprintf("File name '%s' is not an EPUB type.", file.Filename)
		}
	}

	// Save
	for _, file := range files {

		fmt.Println("temp", file.Filename)
		// Generate a unique filename and save the file
		dst := filepath.Join(uploadDir, fmt.Sprintf("%s_%s.epub", strings.ToLower(bookName), time.Now().Format("20060102")))
		// if err := saveFile(file, dst); err != nil {
		// 	return "", fmt.Sprintf("Unable to save file '%s': %s", file.Filename, err.Error())
		// }

		// Generate file URL or path (assuming `uploadDir` is a relative path)
		fileLinks = append(fileLinks, dst)
	}

	return strings.Join(fileLinks, ","), "OK"
}

// @Summary Get a list of books in the the store
// @Tag Book
// @Security bearerToken
// @Produce json
// @Success 200 {object} responses.Response{} "ok"
// @Router /books [get]
func GetListBooks(c *gin.Context) {
	var books []entities.Book

	if err := entities.DB().Preload("BookPreviewImages").Preload("Reviews").Preload("Genres").Preload("Orders").Preload("Carts").Raw("SELECT * FROM books").Find(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   books,
	}

	c.JSON(http.StatusOK, webResponse)
}

// @Summary Get a book by book id
// @Tag Book
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book/{bookId} [get]
func GetBook(c *gin.Context) {
	bookId := c.Param("id")
	var book entities.Book

	if err := entities.DB().Preload("BookPreviewImages").Preload("Reviews").Preload("Genres").Preload("Orders").Preload("Carts").Where("id = ?", bookId).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   book,
	}

	c.JSON(http.StatusOK, webResponse)
}

// @Summary Get a book by book id
// @Tag Book
// @Param bookCreateDto query requests.BookCreateDto true "BookCreateDto"
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book/create [post]
func CreateBook(c *gin.Context) {

	fmt.Println("Create book started")
	// api case
	var bookDto requests.BookCreateDto
	var category entities.Category
	var genres []entities.Genre

	if err := c.ShouldBind(&bookDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Parsed book DTO:", bookDto)

	// check category is exist?
	if tx := entities.DB().Where("id = ?", bookDto.CategoryID).First(&category); tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category [id = " + strconv.Itoa(int(bookDto.CategoryID)) + "] not found."})
		return
	}

	// check genres are exist?
	if tx := entities.DB().Where("id IN ?", bookDto.Genres).Find(&genres); tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Genres [id IN %+v ] not found.", bookDto.Genres)})
		return
	}

	// Handle image uploads
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing form: " + err.Error()})
		return
	}

	file_book_cover_image := form.File["book_cover_image"]
	files_book_preview_images := form.File["book_preview_images"]
	file_book_pdf := form.File["book_pdf"]
	file_book_epub := form.File["book_epub"]

	coverImageLink, status := handleUploadBookCoverImage(file_book_cover_image, "bookCoverImages", bookDto.Name)
	if status != "OK" {
		c.JSON(http.StatusBadRequest, gin.H{"error": status})
	}

	pdfLink, status := handleUploadBookPdf(file_book_pdf, "pdfs", bookDto.Name)
	if status != "OK" {
		c.JSON(http.StatusBadRequest, gin.H{"error": status})
		return
	}

	epubLink, status := handleUploadBookEpub(file_book_epub, "epubs", bookDto.Name)
	if status != "OK" {
		c.JSON(http.StatusBadRequest, gin.H{"error": status})
		return
	}

	bookCreate := entities.Book{
		Name:        bookDto.Name,
		Description: bookDto.Description,
		Price:       bookDto.Price,
		CategoryID:  bookDto.CategoryID,

		CoverImage: coverImageLink,
		BookPdf:    pdfLink,
		BookEpub:   epubLink,

		Genres: genres,
	}

	if err := entities.DB().Create(&bookCreate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Upload + Create BookPreviewImage
	if status := handleUploadBookPreviewImage(files_book_preview_images, "previewImages", bookCreate.ID, bookCreate.Name); status != "OK" {
		c.JSON(http.StatusBadRequest, gin.H{"error": status})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bookCreate,
	}

	c.JSON(http.StatusOK, webResponse)
}

// @Summary Get a list of book preview images
// @Tag Book
// @Param bookUpdateDto query requests.BookUpdateDto true "BookUpdateDto"
// @Security bearerToken
// @Produce json
// @Success 200 {object} responses.Response{} "ok"
// @Router /bookPreviewImages/{id} [get]
func UpdateBook(c *gin.Context) {
	var bookUpdateDto requests.BookUpdateDto
	var book entities.Book
	var genres []entities.Genre

	if err := c.ShouldBind(&bookUpdateDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check book is exist?
	if tx := entities.DB().Where("id = ?", bookUpdateDto.ID).First(&book); tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Book [id: %d] not found.", bookUpdateDto.ID)})
		return
	}

	// check genres is exist?
	if tx := entities.DB().Where("id IN ?", bookUpdateDto.Genres).First(&genres); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Genres [id: %+v] not found.", bookUpdateDto.Genres)})
		return
	}

	book = entities.Book{
		Name:        bookUpdateDto.Name,
		Description: bookUpdateDto.Description,
		Price:       bookUpdateDto.Price,
		CategoryID:  bookUpdateDto.CategoryID,
		Genres:      genres,
	}

	if err := entities.DB().Save(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   book,
	}

	c.JSON(http.StatusOK, webResponse)
}
