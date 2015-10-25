package gate

//import (
//	"fmt"
//)

//func main() {
//	fmt.Println("hello gate")
//}




func main() {
    var dog = new(dogs.Dog)
    stdin := bufio.NewReader(os.Stdin)
    sz, _ := binary.ReadVarint(stdin)
    bytes := make([]byte, sz)
    stdin.Read(bytes)
    err := buf.Unmarshal(dog, bytes)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Fprintf(os.Stderr, "Receiving %s of length %d\n", render(dog), sz)

}