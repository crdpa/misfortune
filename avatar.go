package main

import (
	"math/rand"
	"time"
)

func avatar(avatar string) string {
	var chosenFace string
	personas := make(map[string]string)
	personas["goat"] = `
       (\\           ,'
    .--,\\\__       /
     '-.    a'-.__ 
       |         ')
      / \ _.-'-,';
     /     |   { /
    '      ;    )
          ;'    '
    `
	personas["demon"] = `
            ,'
       (_ )  
    \\\", ) ,
      \/, \( 
      cXc_/_)
             
    `
	personas["skelly"] = `
      .---.  ,' 
     /     \  
    ( () () ) 
     \  M  /  
      |HHH|   
      '---'   
    `
	personas["bat"] = `
    =/\            ,"   /\=
    / \'._   (\_/)'  _.'/ \
   / .''._'--(o.o)--'_.''. \
  /.' _/ |''=/ " \=''| \_ \.\
 /  .' '\;-,'\___/',-;/' \. '\
/.-'       '\(-V-)/'       '-.\
`

	if chosenFace, ok := personas[avatar]; ok {
		return chosenFace
	} else {
		rand.Seed(time.Now().UnixNano())
		k := rand.Intn(len(personas))
		for _, face := range personas {
			if k == 0 {
				return face
			}
			k--
		}
	}
	return chosenFace
}
