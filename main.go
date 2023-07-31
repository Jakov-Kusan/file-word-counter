package main
import (
  "fmt"
  "io/ioutil"
  "sort"
)

func main(){
	var path string
	var words []string
	var word_start int = 0
	var show_number int = 0
	var counter int
	var reappearing_words = map[int]string{}
	var word_number = map[int]int{}
	var banned_words []int
	var used_words []int
	var banned bool
	var used bool
	var sorted_numbers = []int{}
	
	fmt.Print("Path: ")
	fmt.Scan(&path)

	var file,file_error = ioutil.ReadFile(path)

	if file_error != nil{
		fmt.Println("\n[!]ERROR[!]")
		fmt.Println(file_error)
	}else{

		fmt.Print("Only show if the number of duplicates is: ")
        	fmt.Scan(&show_number)
        	show_number--
        }

	for word_end,val := range file{
		if val == '\n'{
			words = append(words,string(file[word_start+1:word_end]))
			word_start = word_end
		}else if val == ' '{
			if word_start == 0{
				words = append(words,string(file[word_start:word_end]))
			}else{
				words = append(words,string(file[word_start+1:word_end]))
			}
			word_start = word_end
		}
	}
	

	for index,word := range words{
		counter = (index+1)/(index+1)-1
		for index2,comp_word := range words{
			banned = false
			for ban_number := range banned_words{
				if banned_words[ban_number] == index2{
					banned = true
				}
			}
                	if word == comp_word && banned == false{
				counter++

				if counter > show_number{
					banned_words = append(banned_words,index2)
				}
			}
        	}
		if counter > show_number{
			reappearing_words[len(reappearing_words)]=string(word)
			word_number[len(reappearing_words)-1]=counter
		}

	}
	for i := 0;i<len(reappearing_words);i++{
		sorted_numbers = append(sorted_numbers,word_number[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sorted_numbers)))

	for i := 0;i<len(sorted_numbers);i++{
		for j := 0;j<len(reappearing_words);j++{
			used = false

			for s := 0;s<len(used_words);s++{
				if used_words[s] == j{
					used = true
				}
			}

			if sorted_numbers[i] == word_number[j] && used == false{
				used_words = append(used_words,j)
				fmt.Println(reappearing_words[j],word_number[j])
			}
		}
	}


}


