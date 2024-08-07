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
func handleUploadBookPreviewImage(files []*multipart.FileHeader, uploadDir string, bookId uint, bookName string) ([]entities.BookPreviewImage, string) {
	var bookPreviewImages []entities.BookPreviewImage

	// validate Image file
	for _, file := range files {
		if !helpers.ValidateFileType(file, "image") {
			return nil, fmt.Sprintf("File name '%s' are not image type.", file.Filename)
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
			return nil, "Cannot create book preview image."
		}

		bookPreviewImages = append(bookPreviewImages, bookPreviewImageCreate)
	}

	return bookPreviewImages, "OK"
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
	var paginationDto requests.PaginationDto
	var books []entities.Book

	if err := c.ShouldBind(&paginationDto); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := entities.DB().Preload("BookPreviewImages").Preload("Reviews").Preload("Genres").Scopes(entities.Paginate(books, &paginationDto)).Find(&books).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if len(books) == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, "No data.", nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", books)
}

// @Summary Get a book by book id
// @Tag Book
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book/{id} [get]
func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book entities.Book

	if tx := entities.DB().Preload("BookPreviewImages").Preload("Reviews").Preload("Genres").Where("id = ?", id).First(&book); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusBadRequest, fmt.Sprintf("Book [id: %s] not found.", id), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", book)
}

// @Summary Get a book by book id
// @Tag Book
// @Param bookCreateDto query requests.BookCreateDto true "BookCreateDto"
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book/create [post]
func CreateBook(c *gin.Context) {
	// api case
	var bookDto requests.BookCreateDto
	var category entities.Category
	var genres []entities.Genre

	if err := c.ShouldBind(&bookDto); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// check category is exist?
	if tx := entities.DB().Where("id = ?", bookDto.CategoryID).First(&category); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, "Category [id = "+strconv.Itoa(int(bookDto.CategoryID))+"] not found.", nil)
		return
	}

	// check genres are exist?
	if tx := entities.DB().Where("id IN ?", bookDto.Genres).Find(&genres); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Genres [id IN %+v ] not found.", bookDto.Genres), nil)
		return
	}

	// Handle image uploads
	form, err := c.MultipartForm()
	if err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, "Error parsing form: "+err.Error(), nil)
		return
	}

	file_book_cover_image := form.File["book_cover_image"]
	files_book_preview_images := form.File["book_preview_images"]
	file_book_pdf := form.File["book_pdf"]
	file_book_epub := form.File["book_epub"]

	coverImageLink, status := handleUploadBookCoverImage(file_book_cover_image, "bookCoverImages", bookDto.Name)
	if status != "OK" {
		helpers.RespondWithJSON(c, http.StatusBadRequest, status, nil)
		return
	}

	pdfLink, status := handleUploadBookPdf(file_book_pdf, "pdfs", bookDto.Name)
	if status != "OK" {
		helpers.RespondWithJSON(c, http.StatusBadRequest, status, nil)
		return
	}

	epubLink, status := handleUploadBookEpub(file_book_epub, "epubs", bookDto.Name)
	if status != "OK" {
		helpers.RespondWithJSON(c, http.StatusBadRequest, status, nil)
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
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Upload + Create BookPreviewImage
	book_preview_images, status := handleUploadBookPreviewImage(files_book_preview_images, "previewImages", bookCreate.ID, bookCreate.Name)
	if status != "OK" {
		helpers.RespondWithJSON(c, http.StatusBadRequest, status, nil)
		return
	}

	bookCreate.BookPreviewImages = book_preview_images

	helpers.RespondWithJSON(c, http.StatusOK, "OK", bookCreate)
}

// @Summary Get a list of book preview images
// @Tag Book
// @Param bookUpdateDto query requests.BookUpdateDto true "BookUpdateDto"
// @Security bearerToken
// @Produce json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book [patch]
func UpdateBook(c *gin.Context) {
	var bookUpdateDto requests.BookUpdateDto
	var book entities.Book
	var genres []entities.Genre

	if err := c.ShouldBind(&bookUpdateDto); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// check book is exist?
	if tx := entities.DB().Where("id = ?", bookUpdateDto.ID).First(&book); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Book [id: %d] not found.", bookUpdateDto.ID), nil)
		return
	}

	// check genres is exist?
	if tx := entities.DB().Where("id IN ?", bookUpdateDto.Genres).First(&genres); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusBadRequest, fmt.Sprintf("Genres [id: %+v] not found.", bookUpdateDto.Genres), nil)
		return
	}

	// TODO: Update file

	book.Name = bookUpdateDto.Name
	book.Description = bookUpdateDto.Description
	book.Price = bookUpdateDto.Price
	book.CategoryID = bookUpdateDto.CategoryID
	book.Genres = genres

	if err := entities.DB().Save(&book).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", book)
}

// @Summary Get a book by book id
// @Tag Book
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book entities.Book

	if tx := entities.DB().Where("id = ?", id).First(&book); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Book [id: %s] not found.", id), nil)
		return
	}

	if err := entities.DB().Where("id = ?", id).Delete(&book).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", book)
}
