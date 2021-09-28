package mongo

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path"
	"testing"
)

func loadEnv(t *testing.T) error {
	dirname, err := os.Getwd()
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	dir, err := os.Open(path.Join(dirname, "../../../"))
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	err = godotenv.Load(os.ExpandEnv(dir.Name() + "/.env.mongo.test"))
	if err != nil {
		log.Println("ERROR:", err.Error())
		t.Fail()
	}
	return err
}
