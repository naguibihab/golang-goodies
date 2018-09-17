package main

import (
	"encoding/json"
	"log"
	"regexp"
)

const open = 91
const comma = 44
const close = 93

func parse(v string) (*node, string, error) {
	rem := v
	root := &node{}
	root.Name = ""
	
	r, _ := regexp.Compile("\\w*")
	root.Name = r.FindString(rem)
	
	// log.Printf("Found name %s in %s",root.Name,rem)
	
	if(len(root.Name) > 0) {
		rem = rem[len(root.Name):]
	}
	
	// log.Printf("Rem is now %s",rem)
	
	for len(rem) > 0 {
		currentChar := rem[0]
		
		if currentChar == open {
			
			// log.Println("Found an opening character, first born")
			child, newRem, _ := parse(rem[1:])
			rem = newRem
			root.Children = append(root.Children,child)
		
		} else if currentChar == close {
			
			// log.Println("Found a closing character, last born, return to parent")
			break
			
		} else if currentChar == comma {
			
			// log.Println("Found a separating character, next sibling, return to parent")
			break
			
		} else {
		
			// log.Println("Found a letter, because it's not preceded by an opening character then it must be a sibling to existing children")
			child, newRem, _ := parse(rem)
			rem = newRem
			root.Children = append(root.Children,child)
		}
		rem = rem[1:]
	}

	return root, rem, nil
}

type node struct {
	Name     string  `json:"name"`
	Children []*node `json:"children,omitempty"` 
}

var examples = []string{
	"[a,b,c]",
	"[a[aa[aaa],ab,ac],b,c[ca,cb,cc[cca]]]",
}

func main() {
	for i, example := range examples {
		result,_, err := parse(example)
		if err != nil {
			panic(err)
		}
		j, err := json.MarshalIndent(result, " ", " ")
		if err != nil {
			panic(err)
		}
		log.Printf("Example %d: %s - %s", i, example, string(j))
	}
}
