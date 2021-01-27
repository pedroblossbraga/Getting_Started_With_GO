# Getting Started with Go
![https://miro.medium.com/max/700/1*Ifpd_HtDiK9u6h68SZgNuA.png](gopher.png)
(https://miro.medium.com/max/700/1*Ifpd_HtDiK9u6h68SZgNuA.png)


## 1. Hello Go!
- Hello World in Go

## 2. Truncate 
- Rounding a number, removing decimal cases
#### Topics: Basic Variables

## 3. Findian:  Find "I" "A" "N"
- Given a string, prints "Found!" if has:
    - "I" in the begginning
    - "A" in the middle
    - "N" in the end
#### Topics: Strings formatting, functions

## 4. Slice.go 
- get input values from user
- start slice with size 3 and grow as needed
- sort slice at each iteration
#### Content: user prompt, slices, sorting, loops

## 5. Makejson: map -> JSON
- user prompt "name" and "adress" values
- create a map with "name" and "address" keys
- Convert map to JSON (Marshalling)
#### Topics: user prompt, maps, JSON, Marshalling

## 6. read.go : read a .txt file
- user prompt filename (ex: names.txt) -> txt file with name and lastname in each line
- read file 
- create struct "Person" with attributes "fname" and "lname"
- store structs in a slice (for each line)

#### Topics: .txt file reading, structs, slices, functions, loops

-----------------------------------------------------
# From Python to Go
![](py_to_go.png)
(https://www.google.com/url?sa=i&url=https%3A%2F%2Ftowardsdatascience.com%2Fmoving-to-go-from-python-9ebbd9a8aec4&psig=AOvVaw2N1l1u4Tj-SEL-Q1Ztj7ST&ust=1611859456819000&source=images&cd=vfe&ved=0CAIQjRxqFwoTCMie9sjivO4CFQAAAAAdAAAAABAD)
#### Maps vs Dictionaries (Python)
- Go Maps are similar to Python dictionaries: a bijection object wiyh a association {key : value}

#### Slices/arrays vs lists/arrays
- In Go, slices are variable length sequences, and arrays have fixed length
- In Python, lists and arrays have variable length, but arrays support vector operations, and lists support string-like operations (concatenation, multiplication)

- Go slice

                package main
                import "fmt"
                func main() {
                    s1 := []int{1,2,3}
                    s2 := []int{4,5,6}
                    fmt.Println(append(s1,s2...))
                }
                /// [1 2 3 4 5 6]


- Python list

                >>>[1,2,3]+[4,5,6] # = [1,2,3,4,5,6]
                >>>[0]*5 # = [0,0,0,0,0]
- Python array
                
                >>> import numpy as np
                >>> np.array([1,2,3])+np.array([4,5,6]) # = np.array([1+4, 2+5, 3+6])
                >>> np.array([1,2,3])*2    # = np.array([1*2,2*2,3*2])

-----------------------------------------------------
## Coursera link
- https://www.coursera.org/learn/golang-getting-started

### Author
- author: Pedro Blöß Braga

### License
- MIT

### Contact: Any questions, optimization? Feel free to send me.
- https://www.linkedin.com/in/pedro-bl%C3%B6%C3%9F-braga-3263a1136/
