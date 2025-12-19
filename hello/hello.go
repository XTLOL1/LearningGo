package main
import "fmt"

func Hello(name string, language string) string{
	if name == "" {
		name = "World"
	}
	
	return helloPrefix(language) + name
}

func helloPrefix(language string) (prefix string) {
	languageMap := map[string]string {
		"English": "Hello, ",
		"Portuguese": "Ola, ",
	}
	
	if language == "" {
		language = "English"
	}
	val, exists := languageMap[language]
	if exists {
		prefix = val
	} else {
		prefix = "Language Unknown"
	}
	return
}

func main(){
	fmt.Println(Hello("", ""))
}
