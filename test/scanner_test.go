package hOwl

import (
	"fmt"
	"testing"

	tokenizer "github.com/passeriform/hOwl/internal/pkg/tokenizer"
)

func TestOnlyParen(t *testing.T) {
	input := "())]([}){{{{})]{((]}({}{}){}()}}[(})][{]"

	fmt.Printf("%+v\n", tokenizer.Tokenize(input))
	// if tokenizer.tokenize(input) != expect {
	// 	t.Errorf("excepted %#v but got %s", except, err.Error())
	// }
}

// Now this is fucked
func TestOnlyWS(t *testing.T) {
	input := "\t\t\t\t\n\t\n\n\n\t\n\n\t\t\t\t\n\n\n\n\t\t\n\n\n\t\n\t\n\t\t\n\n\t\n\t\n\t\n\t"

	fmt.Printf("%+v\n", tokenizer.Tokenize(input))
	// if tokenizer.tokenize(input) != expect {
	// 	t.Errorf("excepted %#v but got %s", except, err.Error())
	// }
}

func TestOnlyContext(t *testing.T) {
	input := "!??..???..!!?!!!!.?!?.?.????.!...?.??!?!"

	fmt.Printf("%+v\n", tokenizer.Tokenize(input))
	// if tokenizer.tokenize(input) != expect {
	// 	t.Errorf("excepted %#v but got %s", except, err.Error())
	// }
}

func TestOnlyComparison(t *testing.T) {
	input := "=> < >= <<=< <=<=>>=>=<==<> <= ><>== =="

	fmt.Printf("%+v\n", tokenizer.Tokenize(input))
	// if tokenizer.tokenize(input) != expect {
	// 	t.Errorf("excepted %#v but got %s", except, err.Error())
	// }
}

func TestOnlyArrows(t *testing.T) {
	input := "-->-==->>>>=>>==>-=>-=>----->---=>=>>=>="

	fmt.Printf("%+v\n", tokenizer.Tokenize(input))
	// if tokenizer.tokenize(input) != expect {
	// 	t.Errorf("excepted %#v but got %s", except, err.Error())
	// }
}

func TestOnlyIdent(t *testing.T) {
	input := `~jkrI OEvX\~5H&9g @Ex7dA35b'Fm"v&a~lC5bx$$XGDXh#emGo3p'||uv,G~9Yy 3|b~oNMYJ8gRI&`

	// println(input)
	fmt.Printf("%+v\n", tokenizer.Tokenize(input))
	// if tokenizer.tokenize(input) != expect {
	// 	t.Errorf("excepted %#v but got %s", except, err.Error())
	// }
}

func TestComplex(t *testing.T) {
	input := `j9/r8|YycLzD(4bT/!8!bYB6~S6q1j{o/\PWqn?v\tR\t8*:Q->:|sYus=gdjd,{
	==Z9f==>bo3x57=>xp0zY,A	m3/\mP->}u/>\nf:hGG\ - 2ldQZ.9>=zXXx?|dAO)}jz
	pP9TCW,1\ - orvwQBP6~4EKFfkU. b{u8J<>.<aaLi	->PfxD=>QN:8e=9\nL~\7|k	t\t1~PQ+)V1}h>=~HqFD==o
	HyC725b/e7S\t H,Bxvf.	<>WG1)AdWgWm}kKtsLx=*	>=6z=>NwS\n-)jOH{(kfMqN	Nr=u\nGX Ru9W?|xL5QR>=
	=>=>><i.K<JJ}-(c\tx>Lx9	:( - .<>*i82)C< - \t{zdAN7zi:<>*g*|<>*M>2N(Po<S*\ddG->=6ipn<>.IYW2?sbj
	\n.4ae=>ONQ>5=>DlJs=>EIe!shXnS3mmTq - Ad.gx*8=/nFuPUY,->+(r.\t>=9t.yH	6/!e3lmW}:*3\S?b5yuq
	15+I|Gz=s}sF9Pt<h|IR2tc?<}Hg\nlHUqS - :0sOs	~t(vbh7c>=AP-SLBwj*Bo	n:U(wwHcTVDU{?l~T+Ih
	:m9=>WXShecB+\tUN==sMCb	 }|cK47+*,>=qQ}f35o\td	oBb(aR=>1W - (c|<Z?AIi9	<>j+ub(\?G->f\nudt - ==eV\`

	// println(input)
	fmt.Printf("%+v\n", tokenizer.Tokenize(input))
	// if tokenizer.tokenize(input) != expect {
	// 	t.Errorf("excepted %#v but got %s", except, err.Error())
	// }
}
