package main

// import "fmt"

// type bot interface {
// 	getGreeeting() string
// }

// type englishBot struct{}
// type spanishBot struct{}

// func main() {
// 	eb := englishBot{}
// 	sb := spanishBot{}

// 	printGreeting(eb)
// 	printGreeting(sb)
// }

// func printGreeting(b bot) {
// 	fmt.Println(b.getGreeeting())
// }

// // レシーバを関数内の処理で使用しないのであれば型指定だけで十分
// func (englishBot) getGreeeting() string {
// 	// 英語の挨拶を生成するためのロジック
// 	return "Hi There!"
// }

// func (spanishBot) getGreeeting() string {
// 	// スペイン語の挨拶を生成するためのロジック
// 	return "Hola!"
// }
