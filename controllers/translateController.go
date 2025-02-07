package controllers

import (
	"context"
	"fmt"
	//"sync"

	//"io"

	//gtranslate "github.com/gilang-as/google-translate"

	"log"
	"net/http"

	//	"strconv"
	"time"

	//	constant "github.com/DeniChan90/go-dictionary-api/constants"
	"github.com/DeniChan90/go-dictionary-api/database"

	//	helper "github.com/DeniChan90/go-dictionary-api/helpers"
	"github.com/DeniChan90/go-dictionary-api/models"
	"github.com/gin-gonic/gin"

	//	"github.com/go-playground/validator/v10"
	//	"golang.org/x/crypto/bcrypt"
	//"golang.org/x/text/language"
	// "go.mongodb.org/mongo-driver/bson"
	gt "github.com/bas24/googletranslatefree"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var translationsCollection *mongo.Collection = database.OpenCollection(database.Client, "translations")
var wordsCollection *mongo.Collection = database.OpenCollection(database.Client, "words")

func Translate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var request models.TranslateRequest

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		defer cancel()

		translated, err := gt.Translate(request.Text, request.From, request.To)

		if err != nil {
			panic(err.Error())
		} else {
			response := models.TranslateResponse{request.To, translated}
			log.Print(response)
			c.JSON(http.StatusOK, response)
		}
	}
}

func SaveTranslations() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var translations []models.Tranlsation
		var word models.Word

		word.Word_id = primitive.NewObjectID().Hex()
		word.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		word.User_id = c.Param("user_id")

		if err := c.BindJSON(&translations); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, insertWordErr := wordsCollection.InsertOne(ctx, word)

		if insertWordErr != nil {
			msg := fmt.Sprintf("User was not created.")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()

		translationsValue := make([]interface{}, len(translations))

		for i, _ := range translations {
			translations[i].Word_id = word.Word_id
			translations[i].User_id = word.User_id
			translationsValue[i] = translations[i]
		}
		_, insertTranslationErr := translationsCollection.InsertMany(ctx, translationsValue)

		if insertTranslationErr != nil {
			msg := fmt.Sprintf("User was not created.")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()

		c.JSON(http.StatusOK, translations)

	}
}

func DeleteTranslations() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		word_id := c.Param("word_id")
		user_id := c.Param("user_id")

		_, err := translationsCollection.DeleteMany(ctx, bson.D{{"user_id", user_id}, {"word_id", word_id}})

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete error"})
		}

		res, e := wordsCollection.DeleteOne(ctx, bson.D{{"_id", word_id}})

		if e != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete error"})
		}

		defer cancel()

		c.JSON(http.StatusOK, res)

	}
}

func GetUserTanslations() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		lang := c.Query("lang")
		user_id := c.Param("user_id")
		translations, err := translationsCollection.Find(ctx, bson.D{{"user_id", user_id}, {"lang", lang}})

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "BAD Translation ERROR OCCURED :( "})
		}
		defer cancel()

		var translationsValue []bson.M

		if err = translations.All(ctx, &translationsValue); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, translationsValue)

	}
}
func GetRelatedTanslations() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		word_id := c.Query("word_id")
		user_id := c.Param("user_id")
		translations, err := translationsCollection.Find(ctx, bson.D{{"user_id", user_id}, {"word_id", word_id}})

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "BAD Translation ERROR OCCURED :( "})
		}
		defer cancel()

		var translationsValue []bson.M

		if err = translations.All(ctx, &translationsValue); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, translationsValue)

	}
}
