package app

func Sum(a *int, b *int) int {
	return *a + *b
}
func Sub(a *int, b *int) int {
	return *a - *b
}
func Mult(a *int, b *int) int {
	return *a * *b
}
func Divi(a, b *int) (int, int) {
	return *a / *b, *a % *b

}
