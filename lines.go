package piscine

func CheckNumber(arg string) bool {
	if len(arg) == 0 {return false}

	for _, char := range arg {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func CountAlpha(s string) int {
	count := 0
	for _, char := range s {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			count++
		}
	}
	return count
}

func CountChar(str string, c rune) int {
    if len(str) == 0 {return 0}

	count := 0
	for _, char := range str {
		if char == c {
			count++
		}
	}
	return count
}

func PrintIf(str string) string {
	if len(str) == 0 || len(str) >=3 {
		return  "G\n"
	}

	return "Invalid Input\n"
}

func PrintIfNot(str string) string {
	if len(str) < 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}

func RectPerimeter(w, h int) int {
 return 2 * (w + h)
}

func RetainFirstHalf(str string) string {
	if len(str) <= 1 {
		return str
	}
	return str[:len(str)/2]
}

//////
package piscine
//cameltosnakecase
func IsLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func IsUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func IsAlpha(r rune) bool {
	return IsLower(r) || IsUpper(r)
}

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n; i++ {
		if !IsAlpha(runes[i]) {
			return s
		}
	}

	if IsUpper(runes[n-1]) {
		return s
	}

	for i := 0; i < n-1; i++ {
		if IsUpper(runes[i]) && IsUpper(runes[i+1]) {
			return s
		}
	}

	var result []rune
	for i := 0; i < n; i++ {
		result = append(result, runes[i])
		if i < n-1 && IsLower(runes[i]) && IsUpper(runes[i+1]) {
			result = append(result, '_')
		}
	}

	return string(result)
}
//
func CountRepeats(s string) int {
	if len(s) < 2 {
		return 0
	}

	count := 0
	inRepeat := false

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			if !inRepeat {
				count++
				inRepeat = true
			}
		} else {
			inRepeat = false
		}
	}

	return count
}

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	if n < 0 {
		n = -n
	}
	count := 0
	for n > 0 {
		n = n / base
		count++
	}
	return count
}

func FirstWord(s string) string {
	start := 0
	for start < len(s) && s[start] == ' ' {
		start++
	}
	end := start
	for end < len(s) && s[end] != ' ' {
		end++
	}
	return s[start:end] + "\n"
}

func FishAndChips(n int) string {
	if n < 0 { return "error: number is negative" }
	if n%2 == 0 && n%3 == 0 {return "fish and chips"}
	if n%2 == 0 {return "fish"}
	if n%3 == 0 {return "chips"}
	return "error: non divisible"
}


func Gcd(a, b uint) uint {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func HashCode(dec string) string {
	size := len(dec)
	result := make([]byte, size)

	for i, c := range dec {
		hashed := (int(c) + size) % 127
		if hashed < 32 {
			hashed += 33
		}
		result[i] = byte(hashed)
	}

	return string(result)
}

func LastWord(s string) string {
	end := len(s) - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}
	if end < 0 {
		return "\n"
	}
	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}
	return s[start+1:end+1] + "\n"
}

 //{
func LongestWord(s string) string {
	words := fields(s)
	if len(words) == 0 {
		return ""
	}

	longest := words[0]
	for _, word := range words[1:] {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

func fields(s string) []string {
	var words []string
	var current []byte
	inWord := false

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			if inWord {
				words = append(words, string(current))
				current = nil
				inWord = false
			}
		} else {
			current = append(current, c)
			inWord = true
		}
	}

	if inWord {
		words = append(words, string(current))
	}

	return words
}
// }

func RepeatAlpha(s string) string {
	var result string
	for _, r := range s {
		repeatCount := 1 

		if r >= 'a' && r <= 'z' {
			repeatCount = int(r - 'a' + 1)
		} else if r >= 'A' && r <= 'Z' {
			repeatCount = int(r - 'A' + 1)
		}

		for i := 0; i < repeatCount; i++ {
			result += string(r)
		}
	}
	return result
}

// { searchreplace
	package main

	import (
		"fmt"
		"os"
	)
	
	func main() {
		if len(os.Args) != 4 {
			return
		}
	
		str := os.Args[1]
		oldArg := os.Args[2]
		newArg := os.Args[3]

		if len(oldArg) == 0 || len(newArg) == 0 {
			fmt.Println(str)
			return
		}

		oldChar := oldArg[0]
		newChar := newArg[0]

		strBytes := []byte(str)
		found := false
	
		for i := 0; i < len(strBytes); i++ {
			if strBytes[i] == oldChar {
				strBytes[i] = newChar
				found = true
			}
		}
	
		fmt.Println(string(strBytes))
	}
	
// ||
	import (
		"fmt"
		"os"
	)
	
	func main() {
		args := os.Args[1:]
		if len(args) != 3 {
			return
		}
	
		str := args[0]
		oldChar := args[1]
		newChar := args[2]
	
		if len(oldChar) != 1 || len(newChar) != 1 {
			return
		}
	
		if !contains(str, oldChar) {
			fmt.Println(str)
			return
		}
	
		result := replaceAll(str, oldChar, newChar)
		fmt.Println(result)
	}
	
	func contains(s, substr string) bool {
		if len(substr) == 0 {
			return true
		}
		for i := 0; i <= len(s)-len(substr); i++ {
			if s[i:i+len(substr)] == substr {
				return true
			}
		}
		return false
	}
	
	func replaceAll(s, old, new string) string {
		if len(old) == 0 {
			return s
		}
		var result []byte
		for i := 0; i < len(s); {
			if i <= len(s)-len(old) && s[i:i+len(old)] == old {
				result = append(result, new...)
				i += len(old)
			} else {
				result = append(result, s[i])
				i++
			}
		}
		return string(result)
	}

// }

// {
func WordAnatomy(word string) (string, string, string) {
	prefixes := []string{"un", "re", "pre", "mis", "dis", "over", "under", "anti", "inter", "sub"}
	suffixes := []string{"ing", "ed", "er", "est", "ly", "ness", "ment", "tion", "able", "ful"}

	prefix := ""
	root := word
	suffix := ""

	for _, p := range prefixes {
		if len(word) >= len(p) && word[:len(p)] == p && len(p) > len(prefix) {
			prefix = p
		}
	}

	for _, s := range suffixes {
		if len(word) >= len(s) && word[len(word)-len(s):] == s && len(s) > len(suffix) {
			suffix = s
		}
	}

	if len(prefix) > 0 && len(suffix) > 0 {
		root = word[len(prefix) : len(word)-len(suffix)]
	} else if len(prefix) > 0 {
		root = word[len(prefix):]
	} else if len(suffix) > 0 {
		root = word[:len(word)-len(suffix)]
	}

	return prefix, root, suffix
}

func main() {
	tests := []string{
		"unhappy",
		"runner",
		"redoing",
		"kindness",
		"planet",
		"careful",
		"misunderstanding",
		"greatest",
		"underpaid",
	}

	for _, word := range tests {
		prefix, root, suffix := WordAnatomy(word)

		fmt.Printf(
			"%s -> Prefix:%q Root:%q Suffix:%q\n",
			word,
			prefix,
			root,
			suffix,
		)
	}
}
// }

//4
package main
// cleanstr
import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	words := strings.Fields(os.Args[1])
	if len(words) == 0 {
		fmt.Println()
		return
	}
	fmt.Println(strings.Join(words, " "))
}

// ||

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	str := os.Args[1]
	inWord := false
	firstWord := true

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' || str[i] == '\t' {
			inWord = false
		} else {
			if !inWord {
				if !firstWord {
					fmt.Print(" ")
				}
				firstWord = false
				inWord = true
			}
			fmt.Print(string(str[i]))
		}
	}

	fmt.Println()
}

// ||

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	str := os.Args[1]
	inWord := false
	firstWord := true

	for _, r := range str {
		if r == ' ' || r == '\t' {
			inWord = false
		} else {
			if !inWord {
				if !firstWord {
					z01.PrintRune(' ')
				}
				firstWord = false
				inWord = true
			}

			z01.PrintRune(r)
		}
	}

	z01.PrintRune('\n')
}

package main

// expandstr
import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	str := os.Args[1]
	inWord := false
	firstWord := true

	for _, r := range str {
		if r == ' ' || r == '\t' {
			inWord = false
		} else {
			if !inWord {
				if !firstWord {
					z01.PrintRune(' ')
					z01.PrintRune(' ')
					z01.PrintRune(' ')
				}
				firstWord = false
				inWord = true
			}

			z01.PrintRune(r)
		}
	}

	if !firstWord {
		z01.PrintRune('\n')
	}
}
 // ||
 package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	words := strings.Fields(os.Args[1])
	if len(words) == 0 {
		return
	}
	fmt.Println(strings.Join(words, "   "))
}
 //findeprevprime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	if from == to {
		return fmt.Sprintf("%02d\n", from)
	}

	result := ""
	if from < to {
		for i := from; i <= to; i++ {
			if i > from {
				result += ", "
			}
			result += fmt.Sprintf("%02d", i)
		}
	} else {
		for i := from; i >= to; i-- {
			if i < from {
				result += ", "
			}
			result += fmt.Sprintf("%02d", i)
		}
	}
	result += "\n"
	return result
}

// ||

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	var result string
	if from <= to {
		for i := from; i <= to; i++ {
			result += intToString(i)
			if i < to {
				result += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			result += intToString(i)
			if i > to {
				result += ", "
			}
		}
	}
	return result + "\n"
}

func intToString(n int) string {
	tens := n / 10
	ones := n % 10
	
	return string(rune(tens+'0')) + string(rune(ones+'0'))
}

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	inWord := false
	for _, c := range s {
		if c == ' ' {
			inWord = false
		} else if !inWord {
			inWord = true
			if c >= 'a' && c <= 'z' {
				return false
			}
		}
	}
	return true
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var isNegative bool
	if n < 0 {
		isNegative = true
		n = -n
	}

	var result []byte

	for n > 0 {
		digit := n % 10
		result = append(result, byte('0'+digit))
		n /= 10
	}

	if isNegative {
		result = append(result, '-')
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
//

package piscine

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	for i, b := range arr {
		h1 := b >> 4
		h2 := b & 0x0f

		z01.PrintRune(rune(getHexDigit(h1)))
		z01.PrintRune(rune(getHexDigit(h2)))

		if (i+1)%4 == 0 && i != len(arr)-1 {
			z01.PrintRune('\n')
		} else if i != len(arr)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')

	for _, b := range arr {
		if b >= 32 && b <= 126 {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func getHexDigit(n byte) byte {
	if n < 10 {
		return '0' + n
	}
	return 'a' + (n - 10)
}

// ||
func PrintMemory(arr [10]byte) {
    for i := 0; i < 10; i++ {
        fmt.Printf("%02x ", arr[i])
        if (i+1)%4 == 0 || i == 9 {
            fmt.Println()
        }
    }

    for _, b := range arr {
        if b >= 32 && b <= 126 {
            fmt.Printf("%c", b)
        } else {
            fmt.Print(".")
        }
    }
    fmt.Println()
}

package main

import "github.com/01-edu/z01"

func main() {
	first := true

	for i := 9; i >= 0; i-- {
		for j := 9; j >= 0; j-- {
			for k := 9; k >= 0; k-- {
				if i > j && j > k {
					if !first {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					first = false

					z01.PrintRune(rune(i + '0'))
					z01.PrintRune(rune(j + '0'))
					z01.PrintRune(rune(k + '0'))
				}
			}
		}
	}
	z01.PrintRune('\n')
}
// ||
package main

import "fmt"

func main() {
	for a := 9; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {
			for c := b - 1; c >= 0; c-- {
				fmt.Printf("%d%d%d", a, b, c)
				if a > 2 || b > 1 || c > 0 {
					fmt.Print(", ")
				}
			}
		}
	}
	fmt.Println()
}

func ThirdTimeIsACharm(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	if len(result) == 0 {
		return "\n"
	}
	return result + "\n"
}

func WeAreUnique(str1, str2 string) int {
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}

	seen1 := make(map[rune]bool)
	seen2 := make(map[rune]bool)
	seenBoth := make(map[rune]bool)

	for _, c := range str1 {
		seen1[c] = true
	}
	for _, c := range str2 {
		seen2[c] = true
	}

	count := 0
	for c := range seen1 {
		if !seen2[c] && !seenBoth[c] {
			count++
			seenBoth[c] = true
		}
	}
	for c := range seen2 {
		if !seen1[c] && !seenBoth[c] {
			count++
			seenBoth[c] = true
		}
	}

	return count
}

import (
	"strconv"
)

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	var result []byte
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			result = append(result, []byte(strconv.Itoa(count))...)
			result = append(result, s[i-1])
			count = 1
		}
	}
	result = append(result, []byte(strconv.Itoa(count))...)
	result = append(result, s[len(s)-1])

	return string(result)
}

// ||
func ZipString(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		count := 1

		for i+count < len(s) && s[i] == s[i+count] {
			count++
		}

		result += string(rune('0'+count))
		result += string(s[i])

		i += count - 1
	}

	return result
}

//5
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	result := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0, false
		}
		result = result*10 + int(ch-'0')
	}
	return result, true
}

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digits []rune
	for n > 0 {
		digits = append(digits, rune('0'+(n%10)))
		n /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		PrintNbr(0)
		z01.PrintRune('\n')
		return
	}

	num, ok := Atoi(os.Args[1])
	if !ok || num <= 0 {
		PrintNbr(0)
		z01.PrintRune('\n')
		return
	}

	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	PrintNbr(sum)
	z01.PrintRune('\n')
}


// |\
package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println(0)
		return
	}
	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}


func CanJump(nums []uint) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	pos := 0
	for pos < len(nums)-1 {
		steps := int(nums[pos])
		if steps == 0 {
			return false
		}
		pos += steps
		if pos >= len(nums)-1 {
			return true
		}
	}
	return pos == len(nums)-1
}


import "fmt"

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}

	if len(slice) == 0 {
		fmt.Println([]int{})
		return
	}

	var result [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	fmt.Println(result)
}

// ||
package main

import (
	"github.com/01-edu/z01"
)

// PrintNbr prints an integer rune by rune.
func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var digits []rune
	for n > 0 {
		digits = append(digits, rune('0'+(n%10)))
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

func Chunk(slice []int, size int) {
	if size <= 0 {
		z01.PrintRune('\n')
		return
	}

	if len(slice) == 0 {
		z01.PrintRune('[')
		z01.PrintRune(']')
		z01.PrintRune('\n')
		return
	}

	var chunks [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	z01.PrintRune('[')
	for i, chunk := range chunks {
		if i > 0 {
			z01.PrintRune(' ')
		}
		z01.PrintRune('[')
		for j, val := range chunk {
			if j > 0 {
				z01.PrintRune(' ')
			}
			PrintNbr(val)
		}
		z01.PrintRune(']')
	}
	z01.PrintRune(']')
	z01.PrintRune('\n')
}


func ConcatAlternate(slice1, slice2 []int) []int {
	var result []int

	if len(slice1) >= len(slice2) {
		for i := 0; i < len(slice2); i++ {
			result = append(result, slice1[i], slice2[i])
		}
		result = append(result, slice1[len(slice2):]...)
	} else {
		for i := 0; i < len(slice1); i++ {
			result = append(result, slice2[i], slice1[i])
		}
		result = append(result, slice2[len(slice1):]...)
	}

	return result
}

func ConcatSlice(slice1, slice2 []int) []int {
	result := make([]int, len(slice1)+len(slice2))
	copy(result, slice1)
	copy(result[len(slice1):], slice2)
	return result
}

//
package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		return
	}
	first := true
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			if !first {
				fmt.Print("*")
			}
			fmt.Print(i)
			first = false
			n /= i
		}
	}
	fmt.Println()
}

// {
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	res := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0, false
		}
		res = res*10 + int(ch-'0')
	}
	return res, true
}

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var digits []rune
	for n > 0 {
		digits = append(digits, rune('0'+(n%10)))
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, ok := Atoi(os.Args[1])
	if !ok || num <= 1 {
		return
	}

	divisor := 2
	first := true

	for num > 1 {
		if num%divisor == 0 {
			if !first {
				z01.PrintRune('*')
			}
			PrintNbr(divisor)
			first = false
			num /= divisor
		} else {
			divisor++
		}
	}

	z01.PrintRune('\n')
}
//}

package main

import (
	"fmt"
	"os"
)

func isHidden(s1, s2 string) bool {
	if len(s1) == 0 {
		return true
	}
	j := 0
	for i := 0; i < len(s2) && j < len(s1); i++ {
		if s2[i] == s1[j] {
			j++
		}
	}
	return j == len(s1)
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	if isHidden(os.Args[1], os.Args[2]) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
//inter

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	seen := make(map[rune]bool)
	for _, c := range s1 {
		if !seen[c] {
			for _, d := range s2 {
				if c == d {
					fmt.Printf("%c", c)
					seen[c] = true
					break
				}
			}
		}
	}
	fmt.Println()
}
// ||
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	inS2 := make(map[rune]bool)
	for _, char := range s2 {
		inS2[char] = true
	}

	seen := make(map[rune]bool)

	for _, char := range s1 {
		if inS2[char] && !seen[char] {
			z01.PrintRune(char)
			seen[char] = true
		}
	}

	z01.PrintRune('\n')
}
//reversestrcap

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	}
	return r
}

func isWordChar(r rune) bool {
	return r != ' '
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	for _, arg := range os.Args[1:] {
		runes := []rune(arg)
		for i, char := range runes {
			if isWordChar(char) {
				isLastLetter := (i+1 == len(runes)) || runes[i+1] == ' '

				if isLastLetter {
					z01.PrintRune(toUpper(char))
				} else {
					z01.PrintRune(toLower(char))
				}
			} else {
				z01.PrintRune(char)
			}
		}
		z01.PrintRune('\n')
	}
}
// ||

package main

import (
	"fmt"
	"os"
	"strings"
)

func reverseCap(s string) string {
	words := strings.Fields(s)
	var result []string
	for _, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			for i := 0; i < len(runes)-1; i++ {
				runes[i] = rune(strings.ToLower(string(runes[i]))[0])
			}
			runes[len(runes)-1] = rune(strings.ToUpper(string(runes[len(runes)-1]))[0])
			result = append(result, string(runes))
		}
	}
	return strings.Join(result, " ")
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(reverseCap(os.Args[i]))
	}
}
//saveandmiss
func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	result := ""
	pos := 0
	for _, c := range arg {
		if pos < num {
			result += string(c)
		}
		pos++
		if pos == 2*num {
			pos = 0
		}
	}
	return result
}

// union
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}
	seen := make(map[rune]bool)
	combined := os.Args[1] + os.Args[2]

	for _, char := range combined {
		if !seen[char] {
			z01.PrintRune(char)
			seen[char] = true
		}
	}
	z01.PrintRune('\n')
}
//wdmatch
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	i := 0 
	j := 0 

	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}

	if i == len(s1) {
		for _, char := range s1 {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
// |
package main

import (
	"fmt"
	"os"
)

func canWrite(s1, s2 string) bool {
	if len(s1) == 0 {
		return true
	}
	j := 0
	for i := 0; i < len(s2) && j < len(s1); i++ {
		if s2[i] == s1[j] {
			j++
		}
	}
	return j == len(s1)
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	if canWrite(os.Args[1], os.Args[2]) {
		fmt.Println(os.Args[1])
	}
}

// 6
package piscine

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	var filtered []rune
	for _, ch := range str {
		if ch != ' ' {
			filtered = append(filtered, ch)
		}
	}

	if len(filtered) < 5 {
		return "Invalid Input\n"
	}

	var result []rune
	count := 0

	for i := 0; i < len(filtered); i++ {
		if (i+1)%6 == 0 {
			continue
		}

		if count > 0 && count%5 == 0 {
			result = append(result, ' ')
		}

		result = append(result, filtered[i])
		count++
	}

	result = append(result, '\n')

	return string(result)
}
// ||
func FifthAndSkip(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	noSpaces := ""
	for _, c := range str {
		if c != ' ' {
			noSpaces += string(c)
		}
	}

	if len(noSpaces) < 5 {
		return "Invalid Input\n"
	}

	result := ""
	for i := 0; i < len(noSpaces); {
		end := i + 5
		if end > len(noSpaces) {
			end = len(noSpaces)
		}
		if i > 0 {
			result += " "
		}
		result += noSpaces[i:end]
		i += 5
		if i < len(noSpaces) {
			i++
		}
	}

	return result + "\n"
}

//
package main

import "fmt"

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	dotCount := 0
	for i, ch := range dec {
		if ch == '-' {
			if i != 0 {
				return dec + "\n"
			}
		} else if ch == '.' {
			dotCount++
			if dotCount > 1 {
				return dec + "\n"
			}
		} else if ch < '0' || ch > '9' {
			return dec + "\n"
		}
	}

	dotIndex := -1
	for i := 0; i < len(dec); i++ {
		if dec[i] == '.' {
			dotIndex = i;
			break
		}
	}

	if dotIndex == -1 {
		return dec + "\n"
	}

	fraction := dec[dotIndex+1:]
	
	if fraction == "" || fraction == "0" {
		return dec + "\n"
	}

	result := dec[:dotIndex] + fraction

	isNegative := false
	if len(result) > 0 && result[0] == '-' {
		isNegative = true
		result = result[1:]
	}

	start := 0
	for start < len(result)-1 && result[start] == '0' {
		start++
	}
	result = result[start:]

	if isNegative {
		result = "-" + result
	}

	return result + "\n"
}

// ||
package main

import (
	"fmt"
	"strings"
)
func NotDecimal(dec string) string {
	if len(dec) == 0 {
		return "\n"
	}
	hasDecimal := false
	hasNonZeroAfterDecimal := false
	dotIndex := -1

	for i, c := range dec {
		if c == '.' {
			if hasDecimal {
				return dec + "\n" 
			}
			hasDecimal = true
			dotIndex = i
		} else if c == '-' && i == 0 {
		} else if c < '0' || c > '9' {
			return dec + "\n" 
		}
	}

	if !hasDecimal {
		return dec + "\n" 
	}

	afterDot := dec[dotIndex+1:]
	for _, c := range afterDot {
		if c != '0' {
			hasNonZeroAfterDecimal = true
			break
		}
	}
	if !hasNonZeroAfterDecimal {
		return dec + "\n" 
	}

	result := strings.Replace(dec, ".", "", 1)
	return result + "\n"
}
//
func RevConcatAlternate(slice1, slice2 []int) []int {
	idx1 := len(slice1) - 1
	idx2 := len(slice2) - 1
	result := make([]int, 0, len(slice1)+len(slice2))

	for idx1 >= 0 || idx2 >= 0 {
		if idx1 > idx2 {
			result = append(result, slice1[idx1])
			idx1--
		} else if idx2 > idx1 {
			result = append(result, slice2[idx2])
			idx2--
		} else {
			for idx1 >= 0 {
				result = append(result, slice1[idx1], slice2[idx2])
				idx1--
				idx2--
			}
		}
	}

	return result
}

//

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return nil
	}

	start := nbrs[0]
	if start < 0 {
		start = len(a) + start
	}
	if start < 0 {
		start = 0
	}
	if start > len(a) {
		return nil
	}

	if len(nbrs) == 1 {
		return a[start:]
	}

	end := nbrs[1]
	if end < 0 {
		end = len(a) + end
	}
	if end < 0 {
		end = 0
	}
	if end > len(a) {
		end = len(a)
	}
	if end < start {
		return nil
	}

	return a[start:end]
}

//7

//findpairs
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findPairs(arr []int, target int) [][]int {
	var pairs [][]int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				pairs = append(pairs, []int{i, j})
			}
		}
	}
	return pairs
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}
	arrStr := os.Args[1]
	targetStr := os.Args[2]
	if !strings.HasPrefix(arrStr, "[") || !strings.HasSuffix(arrStr, "]") {
		fmt.Println("Invalid input.")
		return
	}
	arrStr = strings.Trim(arrStr, "[]")
	if arrStr == "" {
		fmt.Println("Invalid input.")
		return
	}
	parts := strings.Split(arrStr, ", ")
	var arr []int
	for _, p := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			fmt.Printf("Invalid number: %s\n", p)
			return
		}
		arr = append(arr, n)
	}
	target, err := strconv.Atoi(targetStr)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}
	pairs := findPairs(arr, target)
	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
	} else {
		fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
	}
}
// ||

package main

import (
	"os"
)

func main() {
	if len(os.Args) != 3 {
		printString("Invalid input.\n")
		return
	}

	arrStr := os.Args[1]
	targetStr := os.Args[2]

	arrStr = trimSpace(arrStr)
	n := len(arrStr)
	if n < 2 || arrStr[0] != '[' || arrStr[n-1] != ']' {
		printString("Invalid input.\n")
		return
	}

	inner := arrStr[1 : n-1]
	
	// Handle empty array case []
	var arr []int
	if !innerAllSpace(inner) {
		rawElements := splitByComma(inner)
		for _, elem := range rawElements {
			trimmed := trimSpace(elem)
			val, ok := customAtoi(trimmed)
			if !ok {
				printString("Invalid number: ")
				printString(trimmed)
				printString("\n")
				return
			}
			arr = append(arr, val)
		}
	}

	targetTrimmed := trimSpace(targetStr)
	target, ok := customAtoi(targetTrimmed)
	if !ok {
		printString("Invalid target sum.\n")
		return
	}

	// Find all pairs that sum up to the target
	var pairs [][2]int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				pairs = append(pairs, [2]int{i, j})
			}
		}
	}

	// Output results based on findings
	if len(pairs) == 0 {
		printString("No pairs found.\n")
	} else {
		printString("Pairs with sum ")
		printInt(target)
		printString(": ")
		printPairs(pairs)
		printString("\n")
	}
}

// --- Helper Functions (Replacing strings and strconv) ---

func printString(s string) {
	os.Stdout.WriteString(s)
}

func trimSpace(s string) string {
	start := 0
	for start < len(s) && (s[start] == ' ' || s[start] == '\t' || s[start] == '\n' || s[start] == '\r') {
		start++
	}
	end := len(s)
	for end > start && (s[end-1] == ' ' || s[end-1] == '\t' || s[end-1] == '\n' || s[end-1] == '\r') {
		end--
	}
	return s[start:end]
}

func innerAllSpace(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' && s[i] != '\r' {
			return false
		}
	}
	return true
}

func splitByComma(s string) []string {
	var result []string
	current := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ',' {
			result = append(result, current)
			current = ""
		} else {
			current += string(s[i])
		}
	}
	result = append(result, current)
	return result
}

func customAtoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	sign := 1
	start := 0
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	if start >= len(s) {
		return 0, false
	}

	val := 0
	for i := start; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			return 0, false
		}
		val = val*10 + int(ch-'0')
	}
	return val * sign, true
}

func printInt(n int) {
	if n == 0 {
		printString("0")
		return
	}
	if n < 0 {
		printString("-")
		n = -n
	}
	var digits []byte
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	os.Stdout.Write(digits)
}

func printPairs(pairs [][2]int) {
	printString("[")
	for i, p := range pairs {
		if i > 0 {
			printString(" ")
		}
		printString("[")
		printInt(p[0])
		printString(" ")
		printInt(p[1])
		printString("]")
	}
	printString("]")
}
//revwstr

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	words := strings.Fields(os.Args[1])
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Print(words[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
// ||

package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	s := os.Args[1]

	var words []string
	current := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			words = append(words, current)
			current = ""
		} else {
			current += string(s[i])
		}
	}
	words = append(words, current)
	for i := len(words) - 1; i >= 0; i-- {
		os.Stdout.WriteString(words[i])
		if i > 0 {
			os.Stdout.WriteString(" ")
		}
	}
	os.Stdout.WriteString("\n")
}
//rostring

package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("\n")
		return
	}

	s := os.Args[1]

	var words []string
	current := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' || s[i] == '\r' {
			if current != "" {
				words = append(words, current)
				current = ""
			}
		} else {
			current += string(s[i])
		}
	}
	if current != "" {
		words = append(words, current)
	}

	if len(words) == 0 {
		os.Stdout.WriteString("\n")
		return
	}

	rotated := make([]string, len(words))
	copy(rotated, words[1:])
	rotated[len(words)-1] = words[0]

	for i, w := range rotated {
		os.Stdout.WriteString(w)
		if i < len(rotated)-1 {
			os.Stdout.WriteString(" ")
		}
	}
	os.Stdout.WriteString("\n")
}
// ||
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	words := strings.Fields(os.Args[1])
	if len(words) == 0 {
		fmt.Println()
		return
	}
	result := strings.Join(words[1:], " ") + " " + words[0]
	fmt.Println(strings.TrimSpace(result))
}
//

import (
	"fmt"
	"strings"
)

func WordFlip(str string) string {
	if len(str) == 0 {
		return "Invalid Output\n"
	}

	words := strings.Fields(str)
	if len(words) == 0 {
		return "\n"
	}

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ") + "\n"
}
// ||
package piscine

func WordFlip(str string) string {
	start := 0
	for start < len(str) && (str[start] == ' ' || str[start] == '\t' || str[start] == '\n' || str[start] == '\r') {
		start++
	}
	end := len(str)
	for end > start && (str[end-1] == ' ' || str[end-1] == '\t' || str[end-1] == '\n' || str[end-1] == '\r') {
		end--
	}

	if start >= end {
		return "Invalid Output\n"
	}

	trimmed := str[start:end]

	var words []string
	current := ""
	for i := 0; i < len(trimmed); i++ {
		if trimmed[i] == ' ' || trimmed[i] == '\t' || trimmed[i] == '\n' || trimmed[i] == '\r' {
			if current != "" {
				words = append(words, current)
				current = ""
			}
		} else {
			current += string(trimmed[i])
		}
	}
	if current != "" {
		words = append(words, current)
	}

	if len(words) == 0 {
		return "Invalid Output\n"
	}

	result := ""
	for i := len(words) - 1; i >= 0; i-- {
		result += words[i]
		if i > 0 {
			result += " "
		}
	}
	result += "\n"

	return result
}
