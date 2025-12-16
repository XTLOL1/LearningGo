package main
import "fmt"

func Hello(name string, language string) string{
	languageMap := map[string]string{
		"English": "Hello, ",
		"Portuguese": "Ola, ",
	}
	if name == "" {
		name = "World"
	}
	if language == "" {
		language = "English"
	}
	val, exists := languageMap[language]
	if exists {
		return val + name
	} else {
		return "Language Unknown"
	}
}

func main(){
	fmt.Println(Hello("", ""))
}
