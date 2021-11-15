package utils

import "fmt"

//def bar(i, all_i, info="", lines=40):
//    i += 1
//    if i > all_i:
//        i = all_i
//
//    pair = all_i / lines / 10
//    i_pairs = int(i / pair)
//    i_left = i_pairs % 10
//    i_comp = int((i_pairs - i_left) / 10)
//    i_empty = lines - i_comp - int(i_left != 0)
//
//    a = '+' * i_comp
//    if i_left != 0:
//        b = str(i_left)
//    else:
//        b = ''
//    c = '-' * i_empty
//
//    comp = (i / all_i) * 100
//    print(f"\r {comp:>7.3f}% [{a}{b}{c}] {info} | ", end='')

func bar(crt float64, all float64, head string) {
	space := 40

	crt++
	if crt > all {
		crt = all
	}

	rate := crt / all
	done := int(rate * float64(space))
	todo := space - done

	doneStr := '+' * done
	todoStr := '-' * todo

	line := fmt.Sprintf("\r %>.3f%/ [%v%v] %v", rate, doneStr, todoStr, head)
	//    print(f"\r {comp:>7.3f}% [{a}{b}{c}] {info} | ", end='')

	fmt.Println(line)
}

//def bar(i, all_i, info="", lines=40):
//    i += 1
//    if i > all_i:
//        i = all_i
//
//    pair = all_i / lines / 10
//    i_pairs = int(i / pair)
//    i_left = i_pairs % 10
//    i_comp = int((i_pairs - i_left) / 10)
//    i_empty = lines - i_comp - int(i_left != 0)
//
//    a = '+' * i_comp
//    if i_left != 0:
//        b = str(i_left)
//    else:
//        b = ''
//    c = '-' * i_empty
//
//    comp = (i / all_i) * 100
//    print(f"\r {comp:>7.3f}% [{a}{b}{c}] {info} | ", end='')

