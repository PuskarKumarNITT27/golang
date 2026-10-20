package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	MongoDB    string
	ServerPort string
}

func Load() (Config, error) {
	// loading env variables form .env file using godotenv package
	// godotenv.Load() ->  loads the env variables in the os file and 
	// then user access using os.G
	// etenv()

	if err := godotenv.Load(); err != nil{
		return Config{}, fmt.Errorf("Error in loading env file ",err)
	}

	mongoURI, err := extractEnv("MONGO_URI")

	if err != nil {
		return Config{}, err
	}

	mongoDB , err := extractEnv("MONGO_DB_NAME")

	if err != nil {
		return Config{}, err
	}

	port, err := extractEnv("PORT")
	if err != nil {
		return Config{},err 
	}

	return Config{
		MongoURI: mongoURI,
		MongoDB: mongoDB,
		ServerPort: port,
	},nil
}

func extractEnv(key string )(string, error){
	val := os.Getenv(key)

	if val == ""{
		return "", fmt.Errorf("Missing env variable")
	}
	return val, nil
}