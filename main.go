package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"math/cmplx"
	"math/rand"
	"net/http"
	"runtime"
	"strings"
	"sync"
	exr "test3/exercise-reader"
	"test3/images"
	"time"
)

const Si = 3.14

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

type Ver struct {
	X, Y float64
}

func (v *Ver) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Ver, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Ver) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Ver) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	//packagesVariablesStatements()

	//flowControlStatements()

	//structures()

	//maps()

	//functionValues()

	//methods()

	//interfaces()

	//groutine.Lesson_1()

	//groutine.Lesson_2()

	//groutine.Lesson_3()

	//groutine.Lesson_4()

	//groutine.Lesson_5()

	//groutine.Lesson_6()

	//groutine.Lesson_7()

	//rangeBytes()

	//parallelDeclare()

	//makeNew()

	//emb := embeded.NewOne()
	//_ = &embeded.One{Two: "as", Three: 1}
	//fmt.Println(emb)
	//fmt.Printf("%T!\n", emb)
	//emb.Embed()
	//
	//now := embeded.Now()
	//// now.Embed() невозможно вызвать, так как тип интерфейса остается только в пакете и не может быть экспортирован
	//// за его пределы, но можно now.Embed() вызвать внутри пакеты где объявлен этот интерфейс
	//fmt.Printf("%T!%d\n", now, now)

	//http.HandleFunc("/", sayhello)           // Устанавливаем роутер
	//err := http.ListenAndServe(":8083", nil) // устанавливаем порт веб-сервера
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	//start()
	//mapType()

	byteString()
	// dfsfsd
}

func byteString() {
	bt := []byte("GetItems:")
	bt = append(bt, []byte("fg")...)
	str := string(bt)

	newBtStr := []byte(str)
	equals := bytes.Equal(bt, newBtStr)

	log.Printf("\n bytes : %v,\n string : %v,\n new bytes : %v,\n equals: %v", bt, str, newBtStr, equals)

}

func mapType() {
	type expectedPropertiesPerPayment map[int]struct {
		status  string
		reasons []string
	}

	type structure1 struct {
		status  string
		reasons []string
	}

	structure := structure1{status: "23", reasons: []string{"23", "55"}}
	//structure := structure1{}

	exp := make(expectedPropertiesPerPayment)
	exp[1] = structure

	fmt.Printf("%v", exp)
}

func start() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	wq := make(chan interface{})
	workerCount := 2

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doit(i, wq, done, &wg, i+2)
	}

	for i := 0; i < workerCount; i++ {
		wq <- i
	}

	//done <- struct{}{}
	//done <- struct{}{}
	close(done)
	wg.Wait()

	fmt.Println("all done!")
}

func doit(workerId int, wq <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup, sleep int) {
	fmt.Printf("[%v] is running\n", workerId)
	defer wg.Done()
	for {
		time.Sleep(time.Second * time.Duration(sleep))
		select {
		case m, l := <-wq:
			fmt.Printf("[%v] m => %v l=%v\n", workerId, m, l)
		case m, l := <-done:
			fmt.Printf("[%v] is exit => %v read=%v done\n", workerId, m, l)
			return
		}
	}
}

func wg() {
	fmt.Println(runtime.NumGoroutine()) //1

	var wg sync.WaitGroup
	var mtx sync.Mutex
	var b int32 = 2

	for i := 1; i <= 2; i++ {
		wg.Add(1)

		i := i
		go func() {
			defer wg.Done()

			mtx.Lock()
			b++
			mtx.Unlock()

			worker(i)
		}()
	}
	fmt.Println(runtime.NumGoroutine())      //3
	fmt.Printf("lock variable b is %v\n", b) //2

	wg.Wait()
	fmt.Printf("lock variable b is %v\n", b) // 4
}

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет!")
}

func makeNew() {
	var p *[]int = new([]int)          // allocates slice structure; *p == nil; rarely useful
	var v []int = make([]int, 10, 100) // the slice v now refers to a new array of 100 ints

	//// Unnecessarily complex:
	//var p *[]int = new([]int)
	//*p = make([]int, 100, 100)
	//
	//// Idiomatic:
	//v := make([]int, 100)

	var n float32
	var f float64
	var d uintptr
	n = -200000000000000000000000000000000000000
	f = -30000000000000000000000000000000000000000000000000000000000000000000
	d = 2211111111111111111

	//min := int(^uint(0) >> 1)
	min := ^uint(0)
	fmt.Println(p, v, n, f, d, min)

	up := 3 << 1
	down := 5 >> 2
	fmt.Println(up, down)

	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	x = append(x, y...)
	fmt.Println(x)

}

func parallelDeclare() {
	a := []int{1, 2, 3, 4, 5, 6, 7}

Loop:
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
		break Loop
	}

	fmt.Println(a)
}

func rangeBytes() {
	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func interfaces() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
	// lesson 15 приведение типов

	var a interface{} = "hello"

	s := a.(string)
	fmt.Println(s)

	s, ok := a.(string)
	fmt.Println(s, ok)

	f, ok := a.(float64)
	fmt.Println(f, ok)

	//f = a.(float64) // panic
	//fmt.Println(f)

	// lesson 16
	do(21)
	do("hello")
	do(true)

	// php __toString()
	b := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(b, z)

	//lesson 18
	task()

	//lesson 19
	err := run()
	if err != nil {
		fmt.Println(err)
	}

	//lesson 21
	reader()
	//lesson 22
	exr.Reader()
	//lesson 24
	images.Image()
}

func reader() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	buffer := bytes.Buffer{}
	for key, mask := range ip {
		if key == len(ip)-1 {
			buffer.WriteString(fmt.Sprintf("%v", mask))
		} else {
			buffer.WriteString(fmt.Sprintf("%v.", mask))
		}
	}

	return buffer.String()
}

func task() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func methods() {
	var newVer Ver = Ver{1, 2}
	fmt.Println(newVer)
	fmt.Println(newVer.Abs())
	newVer.Scale(10)
	fmt.Println(newVer)

	selectVer := &Ver{2, 4}
	selectVer.Scale(10)
	fmt.Println(selectVer)

	v := Ver{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Ver{3, 4}
	p.Scale(2)
	ScaleFunc(p, 10)

	fmt.Println(v, p)

	s := &Ver{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*s))

}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	//замыкает sum как статическую переменную для каждого контекста вызова
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func functionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func maps() {
	type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex

	m = make(map[string]Vertex)
	fmt.Println(m)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var v = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(v)

	var z = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(z)
	fmt.Printf("%T\n", z["Bell Labs"])

	g := make(map[string]int)
	g["Answer"] = 42
	g["Own"] = 77
	fmt.Println("The value:", g["Answer"])

	g["Answer"] = 48
	fmt.Println("The value:", g["Answer"])

	delete(g, "Answer")
	fmt.Println("The value:", g["Answer"])

	u, existKeyBool := g["Answer"]
	fmt.Println("The value:", u, "Present?", existKeyBool)

	u, existKeyBool = g["Own"]
	fmt.Println("The value:", u, "Present?", existKeyBool)
}

func packagesVariablesStatements() {
	fmt.Println("Hello, 世界")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 13))

	newF, newS := swap("one", "two")
	fmt.Println(newF, newS)

	one, two, tree := split(22)
	fmt.Println(one, two, tree)

	var c, python, java = true, false, "no!"
	fmt.Println(c, python, java)

	var q, w int = 1, 2
	k := 3
	ci, pythone, javac := true, false, "no!"

	fmt.Println(q, w, k, ci, pythone, javac)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var o int
	var f float64
	var b bool
	var s string
	var n uint64
	var m uintptr
	fmt.Printf("%v %v %v %q %v %v\n", o, f, b, s, n, m)

	var x, y int = 3, 4
	var l float64 = math.Sqrt(float64(x*x + y*y))
	var v uint = uint(l)
	fmt.Println(x, y, v)

	fmt.Printf("v is of type %T\n", v)

	fmt.Println(Si)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	//fmt.Printf("v is of type %T\n", Big) // constant 1267650600228229401496703205376 overflows int
	fmt.Printf("v is of type %v\n", Small)
}

func flowControlStatements() {
	var sum = 0
	// for
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	// while
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//while true
	for {
		fmt.Println(sum)
		break
	}

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		po(3, 3, 10),
		po(3, 3, 20),
	)

	fmt.Print("Go runs on ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

type Vertex struct {
	X int
	Y int
}

func structures() {
	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	fmt.Println(p)  // read i through the pointer
	fmt.Println(&p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j // point to j
	fmt.Println(*p)
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	var vertex Vertex = Vertex{1, 2}
	vertex.X = 22
	fmt.Println(vertex, vertex.X, vertex.Y)

	v := Vertex{1, 2}
	n := &v
	(*n).X = 1e9
	n.Y = 2e9
	fmt.Println(v)

	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		k  = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, v2, v3, k, *k)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// включительно со 2 ключа до 3, исключая 4 и далее
	var s []int = primes[2:4]
	s[1] = 22
	primes[2] = 33
	fmt.Println(primes)
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	x := names[0:2]
	b := names[1:3]
	fmt.Println(x, b)

	// меняет значение в срезе, это изменяет базовый массив, все срезы ссылающиеся на массив меняют это значение
	b[0] = "XXX"
	fmt.Println(x, b)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	h := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(h)

	s1 := []int{2, 3, 5, 7, 11, 13}

	//начиная с индекса 1, исключая индекс 4 и выше
	s2 := s1[1:4]
	fmt.Println(s2)

	// начиная с индекса 0, исключая индекс 2 и выше
	s3 := s1[:2]
	fmt.Println(s3)

	// начиная с 1 индекса и до конца
	s4 := s1[1:]
	fmt.Println(s4)

	e := []int{2, 3, 5, 7, 11, 13}
	printSlice(e)

	// Slice the slice to give it zero length.
	s5 := e[:0]
	printSlice(s5)

	// Extend its length.
	s6 := e[:4]
	printSlice(s6)

	// Drop its first two values.
	s7 := e[2:]
	printSlice(s7)

	var w []int
	fmt.Println(w, len(w), cap(w))
	if w == nil {
		fmt.Println("nil!")
	}

	sliceWithMake()

	sliceOfSlices()

	appendToSlice()

	rangeSlice()
}

func rangeSlice() {
	/*	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

		for i, v := range pow {
			fmt.Printf("2**%d = %d\n", i, v)
		}*/

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func appendToSlice() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func sliceWithMake() {
	a := make([]int, 0, 5)
	printSliceOne("a", a)

	b := make([]int, 5)
	b[1], b[2], b[3], b[4] = 1, 2, 3, 4
	printSliceOne("b", b)

	// включаеет с 0 по 1 индекс
	c := b[:2]
	printSliceOne("c", c)

	// включает с 0 до 1 индекс среза "с" потому что он уже из двух элементов
	d := c[0:]
	printSliceOne("d", d)

	// независимо от того, что "d" срез уже из двух элементов, берется базовый массив и с него делается срез, т.к.
	// в нем указаны необходимые ключи
	f := d[0:5]
	printSliceOne("d", f)
}

func sliceOfSlices() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceOne(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func po(x, n, lim float64) float64 {
	v := math.Pow(x, n)
	if v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y, s int) {
	x = sum * 4 / 9
	y = sum - x
	s = x + 7
	return
}

func sqrt(x float64) string {
	if int(x) < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
