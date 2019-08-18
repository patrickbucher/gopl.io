package pipeline

func Pipeline(message string, nodes uint) <-chan string {
	root := make(chan string)
	source := root
	var target chan string
	for i := 0; i < int(nodes); i++ {
		target = make(chan string)
		go func(source <-chan string, target chan<- string) {
			message := <-source
			target <- message
		}(source, target)
		source = target
	}
	root <- message
	return target
}
