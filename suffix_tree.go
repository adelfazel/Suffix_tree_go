package main

import (
	"fmt"
)

type recMap struct {
	thisLayer map[string]recMap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func (m recMap) findBestMatch(matchKey *string) (int, string) {
	for key := range m.thisLayer {
		idx := 0
		if key[idx] == (*matchKey)[idx] {
			mi := min(len(key), len(*matchKey))
			for idx < mi {
				if key[idx] == (*matchKey)[idx] {
					idx++
				} else {
					break
				}
			}
			return idx, key
		}
	}
	return 0, ""
}

func stringSplitAt(str string, at int) (string, string) {
	return str[:at], str[at:]
}
func (m recMap) isBranch(key string) bool {
	return len(m.thisLayer[key].thisLayer) == 0
}

func (m recMap) updateKey(newKey *string) {
	// cases to consider
	// 1) if the new key has zero match with exsiting keys; just add it
	// 2) if the new key is contained in an existing key, nothing to do
	// 3) if existing key is contained in the new key;
	//    a) this could be branch. In this case, recursively push ramining part
	//    b) this is not a branch; update key to match the new key
	// 4) neither of above. Create a branch;
	matchIdx, existingkey := m.findBestMatch(newKey)
	if existingkey == "" {
		m.thisLayer[*newKey] = createNewTree() // add the key if there is no match
	} else {
		endOfKeyBranch := m.thisLayer[existingkey]
		bestMatch, existingKeyRemaining := stringSplitAt(existingkey, matchIdx)
		newKeyRemaining := (*newKey)[matchIdx:]
		if newKeyRemaining == "" {
			return
		}
		if existingKeyRemaining == "" {
			endOfKeyBranch.updateKey(&newKeyRemaining)
			/*
				if m.isBranch(existingkey) {

					endOfKeyBranch.updateKey(&newKeyRemaining)
				} else {
					delete(m.thisLayer, existingkey)
					m.thisLayer[*newKey] = createNewTree()
					}
			*/

		}

		delete(m.thisLayer, existingkey)
		m.thisLayer[bestMatch] = createNewTree()
		m.thisLayer[bestMatch].thisLayer[existingKeyRemaining] = endOfKeyBranch
		m.thisLayer[bestMatch].thisLayer[newKeyRemaining] = createNewTree()
	}
}

func createNewTree() recMap {
	newTree := recMap{}
	newTree.thisLayer = make(map[string]recMap)
	return newTree
}

func make_suffix_tree(inputString *string) recMap {
	res := createNewTree()
	for outeridx := 0; outeridx < len(*inputString); outeridx++ {
		currentStr := (*inputString)[outeridx:]
		//fmt.Println("outeridx:", outeridx, "currentStr:", currentStr)
		res.updateKey(&currentStr)
		fmt.Println(currentStr)
		fmt.Println("res:", res)
	}
	return res
}

func printRes(res recMap) {
	for k, v := range res.thisLayer {
		fmt.Println(k)
		if k != "$" {
			printRes(v)
		}
	}
}

func main() {
	var inputString string
	fmt.Scanf("%s", &inputString)
	suffixTree := make_suffix_tree(&inputString)
	printRes(suffixTree)
}
